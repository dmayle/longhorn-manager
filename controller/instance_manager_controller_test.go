package controller

import (
	"fmt"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/controller"

	"github.com/longhorn/longhorn-manager/datastore"
	"github.com/longhorn/longhorn-manager/types"

	longhorn "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	lhfake "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned/fake"
	lhinformerfactory "github.com/longhorn/longhorn-manager/k8s/pkg/client/informers/externalversions"

	. "gopkg.in/check.v1"
)

type InstanceManagerTestCase struct {
	controllerID string
	nodeDown     bool
	nodeID       string

	currentPodPhase v1.PodPhase
	currentOwnerID  string
	currentState    types.InstanceManagerState

	expectedOwnerID  string
	expectedPodCount int
	expectedStatus   types.InstanceManagerStatus
	expectedType     types.InstanceManagerType
}

func newTolerationSetting() *longhorn.Setting {
	return &longhorn.Setting{
		ObjectMeta: metav1.ObjectMeta{
			Name: string(types.SettingNameTaintToleration),
		},
		Setting: types.Setting{
			Value: "",
		},
	}
}

func newEngineImage(image string, state types.EngineImageState) *longhorn.EngineImage {
	return &longhorn.EngineImage{
		ObjectMeta: metav1.ObjectMeta{
			Name:       types.GetEngineImageChecksumName(image),
			Namespace:  TestNamespace,
			UID:        uuid.NewUUID(),
			Finalizers: []string{longhornFinalizerKey},
		},
		Spec: types.EngineImageSpec{
			Image: image,
		},
		Status: types.EngineImageStatus{
			OwnerID: TestNode1,
			State:   state,
			EngineVersionDetails: types.EngineVersionDetails{
				Version:   "latest",
				GitCommit: "latest",

				CLIAPIVersion:           3,
				CLIAPIMinVersion:        3,
				ControllerAPIVersion:    3,
				ControllerAPIMinVersion: 3,
				DataFormatVersion:       1,
				DataFormatMinVersion:    1,
			},
		},
	}
}

func newInstanceManager(
	name string,
	imType types.InstanceManagerType,
	currentState types.InstanceManagerState,
	currentOwnerID, nodeID, ip string,
	instances map[string]types.InstanceProcess,
	isDeleting bool) *longhorn.InstanceManager {

	im := &longhorn.InstanceManager{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: TestNamespace,
			UID:       uuid.NewUUID(),
			Labels:    types.GetInstanceManagerLabels(TestNode1, getTestEngineImageName(), imType),
		},
		Spec: types.InstanceManagerSpec{
			EngineImage: TestEngineImage,
			NodeID:      nodeID,
			Type:        imType,
		},
		Status: types.InstanceManagerStatus{
			OwnerID:      currentOwnerID,
			CurrentState: currentState,
			IP:           ip,
			Instances:    instances,
		},
	}

	if isDeleting {
		now := metav1.NewTime(time.Now())
		im.DeletionTimestamp = &now
	}
	return im
}

func newPod(phase v1.PodPhase, name, namespace, nodeID string) *v1.Pod {
	if phase == "" {
		return nil
	}
	ip := ""
	if phase == v1.PodRunning {
		ip = TestIP1
	}
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			ServiceAccountName: TestServiceAccount,
			NodeName:           nodeID,
		},
		Status: v1.PodStatus{
			Phase: phase,
			PodIP: ip,
		},
	}
}

func newTestInstanceManagerController(lhInformerFactory lhinformerfactory.SharedInformerFactory,
	kubeInformerFactory informers.SharedInformerFactory, lhClient *lhfake.Clientset, kubeClient *fake.Clientset,
	controllerID string) *InstanceManagerController {

	volumeInformer := lhInformerFactory.Longhorn().V1beta1().Volumes()
	engineInformer := lhInformerFactory.Longhorn().V1beta1().Engines()
	replicaInformer := lhInformerFactory.Longhorn().V1beta1().Replicas()
	engineImageInformer := lhInformerFactory.Longhorn().V1beta1().EngineImages()
	nodeInformer := lhInformerFactory.Longhorn().V1beta1().Nodes()
	settingInformer := lhInformerFactory.Longhorn().V1beta1().Settings()
	imInformer := lhInformerFactory.Longhorn().V1beta1().InstanceManagers()

	podInformer := kubeInformerFactory.Core().V1().Pods()
	cronJobInformer := kubeInformerFactory.Batch().V1beta1().CronJobs()
	daemonSetInformer := kubeInformerFactory.Apps().V1().DaemonSets()
	deploymentInformer := kubeInformerFactory.Apps().V1().Deployments()
	persistentVolumeInformer := kubeInformerFactory.Core().V1().PersistentVolumes()
	persistentVolumeClaimInformer := kubeInformerFactory.Core().V1().PersistentVolumeClaims()
	kubeNodeInformer := kubeInformerFactory.Core().V1().Nodes()

	ds := datastore.NewDataStore(
		volumeInformer, engineInformer, replicaInformer,
		engineImageInformer, nodeInformer, settingInformer, imInformer,
		lhClient,
		podInformer, cronJobInformer, daemonSetInformer, deploymentInformer,
		persistentVolumeInformer, persistentVolumeClaimInformer, kubeNodeInformer,
		kubeClient, TestNamespace)

	imc := NewInstanceManagerController(ds, scheme.Scheme, imInformer, podInformer, kubeClient, TestNamespace,
		controllerID, TestServiceAccount)
	fakeRecorder := record.NewFakeRecorder(100)
	imc.eventRecorder = fakeRecorder
	imc.imStoreSynced = alwaysReady
	imc.pStoreSynced = alwaysReady

	return imc
}

func (s *TestSuite) TestSyncInstanceManager(c *C) {
	var err error

	testCases := map[string]InstanceManagerTestCase{
		"instance manager change ownership": {
			TestNode1, false, TestNode1,
			v1.PodRunning, TestNode2, types.InstanceManagerStateUnknown,
			TestNode1, 1,
			types.InstanceManagerStatus{
				OwnerID:      TestNode1,
				CurrentState: types.InstanceManagerStateRunning,
				IP:           TestIP1,
			},
			types.InstanceManagerTypeEngine,
		},
		"instance manager error then restart immediately": {
			TestNode1, false, TestNode1,
			v1.PodFailed, TestNode1, types.InstanceManagerStateRunning,
			TestNode1, 1,
			types.InstanceManagerStatus{
				OwnerID:      TestNode1,
				CurrentState: types.InstanceManagerStateStarting,
			},
			types.InstanceManagerTypeEngine,
		},
		"instance manager node down": {
			TestNode2, true, TestNode1,
			v1.PodRunning, TestNode1, types.InstanceManagerStateRunning,
			TestNode2, 1,
			types.InstanceManagerStatus{
				OwnerID:      TestNode2,
				CurrentState: types.InstanceManagerStateUnknown,
			},
			types.InstanceManagerTypeEngine,
		},
		"instance manager restarting after error": {
			TestNode1, false, TestNode1,
			v1.PodRunning, TestNode1, types.InstanceManagerStateError,
			TestNode1, 1, types.InstanceManagerStatus{
				OwnerID:      TestNode1,
				CurrentState: types.InstanceManagerStateStarting,
			},
			types.InstanceManagerTypeEngine,
		},
		"instance manager running": {
			TestNode1, false, TestNode1,
			v1.PodRunning, TestNode1, types.InstanceManagerStateStarting,
			TestNode1, 1, types.InstanceManagerStatus{
				OwnerID:      TestNode1,
				CurrentState: types.InstanceManagerStateRunning,
				IP:           TestIP1,
			},
			types.InstanceManagerTypeEngine,
		},
		"instance manager starting engine": {
			TestNode1, false, TestNode1,
			"", TestNode1, types.InstanceManagerStateStopped,
			TestNode1, 1,
			types.InstanceManagerStatus{
				OwnerID:      TestNode1,
				CurrentState: types.InstanceManagerStateStarting,
			},
			types.InstanceManagerTypeEngine,
		},
		"instance manager starting replica": {
			TestNode1, false, TestNode1,
			"", TestNode1, types.InstanceManagerStateStopped,
			TestNode1, 1,
			types.InstanceManagerStatus{
				OwnerID:      TestNode1,
				CurrentState: types.InstanceManagerStateStarting,
			},
			types.InstanceManagerTypeReplica,
		},
	}

	for name, tc := range testCases {
		fmt.Printf("testing %v\n", name)

		kubeClient := fake.NewSimpleClientset()
		kubeInformerFactory := informers.NewSharedInformerFactory(kubeClient, controller.NoResyncPeriodFunc())
		pIndexer := kubeInformerFactory.Core().V1().Pods().Informer().GetIndexer()
		kubeNodeIndexer := kubeInformerFactory.Core().V1().Nodes().Informer().GetIndexer()

		lhClient := lhfake.NewSimpleClientset()
		lhInformerFactory := lhinformerfactory.NewSharedInformerFactory(lhClient, controller.NoResyncPeriodFunc())
		eiIndexer := lhInformerFactory.Longhorn().V1beta1().EngineImages().Informer().GetIndexer()
		imIndexer := lhInformerFactory.Longhorn().V1beta1().InstanceManagers().Informer().GetIndexer()
		sIndexer := lhInformerFactory.Longhorn().V1beta1().Settings().Informer().GetIndexer()
		lhNodeIndexer := lhInformerFactory.Longhorn().V1beta1().Nodes().Informer().GetIndexer()

		imc := newTestInstanceManagerController(lhInformerFactory, kubeInformerFactory, lhClient, kubeClient,
			tc.controllerID)

		// Controller logic depends on the existence of Instance Manager's Engine Image and Toleration Setting.
		setting := newTolerationSetting()
		setting, err = lhClient.LonghornV1beta1().Settings(TestNamespace).Create(setting)
		c.Assert(err, IsNil)
		err = sIndexer.Add(setting)
		c.Assert(err, IsNil)
		ei := newEngineImage(TestEngineImage, types.EngineImageStateReady)
		err = eiIndexer.Add(ei)
		c.Assert(err, IsNil)
		_, err = lhClient.LonghornV1beta1().EngineImages(ei.Namespace).Create(ei)
		c.Assert(err, IsNil)

		// Create Nodes for test. Conditionally add the first Node.
		if !tc.nodeDown {
			kubeNode1 := newKubernetesNode(TestNode1, v1.ConditionTrue, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionTrue)
			err = kubeNodeIndexer.Add(kubeNode1)
			c.Assert(err, IsNil)
			_, err = kubeClient.CoreV1().Nodes().Create(kubeNode1)
			c.Assert(err, IsNil)

			lhNode1 := newNode(TestNode1, TestNamespace, true, types.ConditionStatusTrue, "")
			err = lhNodeIndexer.Add(lhNode1)
			c.Assert(err, IsNil)
			_, err = lhClient.LonghornV1beta1().Nodes(lhNode1.Namespace).Create(lhNode1)
			c.Assert(err, IsNil)
		}

		kubeNode2 := newKubernetesNode(TestNode2, v1.ConditionTrue, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionFalse, v1.ConditionTrue)
		err = kubeNodeIndexer.Add(kubeNode2)
		c.Assert(err, IsNil)
		_, err = kubeClient.CoreV1().Nodes().Create(kubeNode2)

		lhNode2 := newNode(TestNode2, TestNamespace, true, types.ConditionStatusTrue, "")
		err = lhNodeIndexer.Add(lhNode2)
		c.Assert(err, IsNil)
		_, err = lhClient.LonghornV1beta1().Nodes(lhNode2.Namespace).Create(lhNode2)
		c.Assert(err, IsNil)

		im := newInstanceManager(TestInstanceManagerName1, tc.expectedType, tc.currentState, tc.currentOwnerID, tc.nodeID, "", nil, false)
		err = imIndexer.Add(im)
		c.Assert(err, IsNil)
		_, err = lhClient.LonghornV1beta1().InstanceManagers(im.Namespace).Create(im)
		c.Assert(err, IsNil)

		if tc.currentPodPhase != "" {
			pod := newPod(tc.currentPodPhase, im.Name, im.Namespace, im.Spec.NodeID)
			err = pIndexer.Add(pod)
			c.Assert(err, IsNil)
			_, err = kubeClient.CoreV1().Pods(im.Namespace).Create(pod)
			c.Assert(err, IsNil)
		}

		err = imc.syncInstanceManager(getKey(im, c))
		c.Assert(err, IsNil)
		podList, err := kubeClient.CoreV1().Pods(im.Namespace).List(metav1.ListOptions{})
		c.Assert(err, IsNil)
		c.Assert(podList.Items, HasLen, tc.expectedPodCount)

		// Check the Pod that was created by the Instance Manager.
		if tc.currentPodPhase == "" {
			pod, err := kubeClient.CoreV1().Pods(im.Namespace).Get(im.Name, metav1.GetOptions{})
			c.Assert(err, IsNil)
			switch im.Spec.Type {
			case types.InstanceManagerTypeEngine:
				c.Assert(pod.Spec.Containers[0].Name, Equals, "engine-manager")
			case types.InstanceManagerTypeReplica:
				c.Assert(pod.Spec.Containers[0].Name, Equals, "replica-manager")
			}
		}

		if tc.expectedStatus.CurrentState == types.InstanceManagerStateRunning {
			_, exist := imc.instanceManagerMonitorMap[im.Name]
			c.Assert(exist, Equals, true)
		} else {
			_, exist := imc.instanceManagerMonitorMap[im.Name]
			c.Assert(exist, Equals, false)
		}

		updatedIM, err := lhClient.LonghornV1beta1().InstanceManagers(im.Namespace).Get(im.Name, metav1.GetOptions{})
		c.Assert(err, IsNil)
		c.Assert(updatedIM.Status, DeepEquals, tc.expectedStatus)
	}
}
