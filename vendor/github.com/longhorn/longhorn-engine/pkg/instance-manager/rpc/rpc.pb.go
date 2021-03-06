// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProcessSpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Binary               string   `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	Args                 []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	PortCount            int32    `protobuf:"varint,4,opt,name=port_count,json=portCount,proto3" json:"port_count,omitempty"`
	PortArgs             []string `protobuf:"bytes,5,rep,name=port_args,json=portArgs,proto3" json:"port_args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessSpec) Reset()         { *m = ProcessSpec{} }
func (m *ProcessSpec) String() string { return proto.CompactTextString(m) }
func (*ProcessSpec) ProtoMessage()    {}
func (*ProcessSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *ProcessSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessSpec.Unmarshal(m, b)
}
func (m *ProcessSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessSpec.Marshal(b, m, deterministic)
}
func (m *ProcessSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessSpec.Merge(m, src)
}
func (m *ProcessSpec) XXX_Size() int {
	return xxx_messageInfo_ProcessSpec.Size(m)
}
func (m *ProcessSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessSpec proto.InternalMessageInfo

func (m *ProcessSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessSpec) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *ProcessSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ProcessSpec) GetPortCount() int32 {
	if m != nil {
		return m.PortCount
	}
	return 0
}

func (m *ProcessSpec) GetPortArgs() []string {
	if m != nil {
		return m.PortArgs
	}
	return nil
}

type ProcessStatus struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	PortStart            int32    `protobuf:"varint,3,opt,name=port_start,json=portStart,proto3" json:"port_start,omitempty"`
	PortEnd              int32    `protobuf:"varint,4,opt,name=port_end,json=portEnd,proto3" json:"port_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessStatus) Reset()         { *m = ProcessStatus{} }
func (m *ProcessStatus) String() string { return proto.CompactTextString(m) }
func (*ProcessStatus) ProtoMessage()    {}
func (*ProcessStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *ProcessStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessStatus.Unmarshal(m, b)
}
func (m *ProcessStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessStatus.Marshal(b, m, deterministic)
}
func (m *ProcessStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessStatus.Merge(m, src)
}
func (m *ProcessStatus) XXX_Size() int {
	return xxx_messageInfo_ProcessStatus.Size(m)
}
func (m *ProcessStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessStatus proto.InternalMessageInfo

func (m *ProcessStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ProcessStatus) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func (m *ProcessStatus) GetPortStart() int32 {
	if m != nil {
		return m.PortStart
	}
	return 0
}

func (m *ProcessStatus) GetPortEnd() int32 {
	if m != nil {
		return m.PortEnd
	}
	return 0
}

type ProcessCreateRequest struct {
	Spec                 *ProcessSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProcessCreateRequest) Reset()         { *m = ProcessCreateRequest{} }
func (m *ProcessCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessCreateRequest) ProtoMessage()    {}
func (*ProcessCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *ProcessCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessCreateRequest.Unmarshal(m, b)
}
func (m *ProcessCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessCreateRequest.Marshal(b, m, deterministic)
}
func (m *ProcessCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessCreateRequest.Merge(m, src)
}
func (m *ProcessCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessCreateRequest.Size(m)
}
func (m *ProcessCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessCreateRequest proto.InternalMessageInfo

func (m *ProcessCreateRequest) GetSpec() *ProcessSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ProcessDeleteRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessDeleteRequest) Reset()         { *m = ProcessDeleteRequest{} }
func (m *ProcessDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessDeleteRequest) ProtoMessage()    {}
func (*ProcessDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *ProcessDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessDeleteRequest.Unmarshal(m, b)
}
func (m *ProcessDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessDeleteRequest.Marshal(b, m, deterministic)
}
func (m *ProcessDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessDeleteRequest.Merge(m, src)
}
func (m *ProcessDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessDeleteRequest.Size(m)
}
func (m *ProcessDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessDeleteRequest proto.InternalMessageInfo

func (m *ProcessDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProcessGetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessGetRequest) Reset()         { *m = ProcessGetRequest{} }
func (m *ProcessGetRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessGetRequest) ProtoMessage()    {}
func (*ProcessGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *ProcessGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessGetRequest.Unmarshal(m, b)
}
func (m *ProcessGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessGetRequest.Marshal(b, m, deterministic)
}
func (m *ProcessGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessGetRequest.Merge(m, src)
}
func (m *ProcessGetRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessGetRequest.Size(m)
}
func (m *ProcessGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessGetRequest proto.InternalMessageInfo

func (m *ProcessGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProcessResponse struct {
	Spec                 *ProcessSpec   `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Status               *ProcessStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Deleted              bool           `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProcessResponse) Reset()         { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()    {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessResponse.Unmarshal(m, b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessResponse.Size(m)
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

func (m *ProcessResponse) GetSpec() *ProcessSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ProcessResponse) GetStatus() *ProcessStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ProcessResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type ProcessListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessListRequest) Reset()         { *m = ProcessListRequest{} }
func (m *ProcessListRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessListRequest) ProtoMessage()    {}
func (*ProcessListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *ProcessListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessListRequest.Unmarshal(m, b)
}
func (m *ProcessListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessListRequest.Marshal(b, m, deterministic)
}
func (m *ProcessListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessListRequest.Merge(m, src)
}
func (m *ProcessListRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessListRequest.Size(m)
}
func (m *ProcessListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessListRequest proto.InternalMessageInfo

type ProcessListResponse struct {
	Processes            map[string]*ProcessResponse `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProcessListResponse) Reset()         { *m = ProcessListResponse{} }
func (m *ProcessListResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessListResponse) ProtoMessage()    {}
func (*ProcessListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *ProcessListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessListResponse.Unmarshal(m, b)
}
func (m *ProcessListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessListResponse.Marshal(b, m, deterministic)
}
func (m *ProcessListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessListResponse.Merge(m, src)
}
func (m *ProcessListResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessListResponse.Size(m)
}
func (m *ProcessListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessListResponse proto.InternalMessageInfo

func (m *ProcessListResponse) GetProcesses() map[string]*ProcessResponse {
	if m != nil {
		return m.Processes
	}
	return nil
}

type LogRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{8}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LogResponse struct {
	Line                 string   `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

type ProcessReplaceRequest struct {
	Spec                 *ProcessSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	TerminateSignal      string       `protobuf:"bytes,2,opt,name=terminate_signal,json=terminateSignal,proto3" json:"terminate_signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProcessReplaceRequest) Reset()         { *m = ProcessReplaceRequest{} }
func (m *ProcessReplaceRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessReplaceRequest) ProtoMessage()    {}
func (*ProcessReplaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{10}
}

func (m *ProcessReplaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessReplaceRequest.Unmarshal(m, b)
}
func (m *ProcessReplaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessReplaceRequest.Marshal(b, m, deterministic)
}
func (m *ProcessReplaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessReplaceRequest.Merge(m, src)
}
func (m *ProcessReplaceRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessReplaceRequest.Size(m)
}
func (m *ProcessReplaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessReplaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessReplaceRequest proto.InternalMessageInfo

func (m *ProcessReplaceRequest) GetSpec() *ProcessSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ProcessReplaceRequest) GetTerminateSignal() string {
	if m != nil {
		return m.TerminateSignal
	}
	return ""
}

func init() {
	proto.RegisterType((*ProcessSpec)(nil), "ProcessSpec")
	proto.RegisterType((*ProcessStatus)(nil), "ProcessStatus")
	proto.RegisterType((*ProcessCreateRequest)(nil), "ProcessCreateRequest")
	proto.RegisterType((*ProcessDeleteRequest)(nil), "ProcessDeleteRequest")
	proto.RegisterType((*ProcessGetRequest)(nil), "ProcessGetRequest")
	proto.RegisterType((*ProcessResponse)(nil), "ProcessResponse")
	proto.RegisterType((*ProcessListRequest)(nil), "ProcessListRequest")
	proto.RegisterType((*ProcessListResponse)(nil), "ProcessListResponse")
	proto.RegisterMapType((map[string]*ProcessResponse)(nil), "ProcessListResponse.ProcessesEntry")
	proto.RegisterType((*LogRequest)(nil), "LogRequest")
	proto.RegisterType((*LogResponse)(nil), "LogResponse")
	proto.RegisterType((*ProcessReplaceRequest)(nil), "ProcessReplaceRequest")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9b, 0xa4, 0x6d, 0xc6, 0xa5, 0x2d, 0xd3, 0xb4, 0x32, 0xa9, 0x90, 0xc2, 0x22, 0x95,
	0x02, 0xd2, 0x16, 0x05, 0x0e, 0x55, 0xc5, 0xa5, 0x2a, 0x15, 0x97, 0x16, 0x21, 0xe7, 0xc0, 0x31,
	0xda, 0x3a, 0x83, 0x89, 0x70, 0x6c, 0xb3, 0xbb, 0xae, 0x14, 0x89, 0x1f, 0xe0, 0x4f, 0x10, 0x5f,
	0x89, 0xbc, 0x5e, 0xdb, 0x09, 0x31, 0x85, 0xdb, 0xce, 0xdb, 0x79, 0x99, 0xe7, 0x37, 0x2f, 0x0b,
	0x5d, 0x99, 0x06, 0x3c, 0x95, 0x89, 0x4e, 0xfa, 0x47, 0x61, 0x92, 0x84, 0x11, 0x9d, 0x9a, 0xea,
	0x36, 0xfb, 0x7c, 0x4a, 0xb3, 0x54, 0xcf, 0x8b, 0x4b, 0xf6, 0xc3, 0x01, 0xf7, 0xa3, 0x4c, 0x02,
	0x52, 0x6a, 0x94, 0x52, 0x80, 0x08, 0xed, 0x58, 0xcc, 0xc8, 0x73, 0x06, 0xce, 0x49, 0xd7, 0x37,
	0x67, 0x3c, 0x84, 0x8d, 0xdb, 0x69, 0x2c, 0xe4, 0xdc, 0x5b, 0x37, 0xa8, 0xad, 0xf2, 0x5e, 0x21,
	0x43, 0xe5, 0xb5, 0x06, 0xad, 0xbc, 0x37, 0x3f, 0xe3, 0x63, 0x80, 0x34, 0x91, 0x7a, 0x1c, 0x24,
	0x59, 0xac, 0xbd, 0xf6, 0xc0, 0x39, 0xe9, 0xf8, 0xdd, 0x1c, 0xb9, 0xcc, 0x01, 0x3c, 0x02, 0x53,
	0x8c, 0x0d, 0xaf, 0x63, 0x78, 0x5b, 0x39, 0x70, 0x21, 0x43, 0xc5, 0xbe, 0xc3, 0x83, 0x52, 0x8a,
	0x16, 0x3a, 0x53, 0xd8, 0x83, 0x8e, 0xd2, 0x42, 0x97, 0x6a, 0x8a, 0x22, 0xff, 0x0d, 0x92, 0x32,
	0x91, 0xe3, 0x99, 0x0a, 0xad, 0xa2, 0x2d, 0x03, 0xdc, 0xa8, 0xb0, 0x9a, 0xaf, 0xb4, 0x90, 0xda,
	0x6b, 0xd5, 0xf3, 0x47, 0x39, 0x80, 0x8f, 0xc0, 0x8c, 0x1b, 0x53, 0x3c, 0xb1, 0xe2, 0x36, 0xf3,
	0xfa, 0x2a, 0x9e, 0xb0, 0x33, 0xe8, 0xd9, 0xe9, 0x97, 0x92, 0x84, 0x26, 0x9f, 0xbe, 0x65, 0xa4,
	0x34, 0x0e, 0xa0, 0xad, 0x52, 0x0a, 0x8c, 0x06, 0x77, 0xb8, 0xcd, 0x17, 0xdc, 0xf2, 0xcd, 0x0d,
	0x7b, 0x51, 0x31, 0xdf, 0x51, 0x44, 0x35, 0xb3, 0xc1, 0x4b, 0xf6, 0x0c, 0x1e, 0xda, 0xde, 0xf7,
	0xa4, 0xef, 0x6b, 0xcc, 0x60, 0xd7, 0x36, 0xfa, 0xa4, 0xd2, 0x24, 0x56, 0xf4, 0x6f, 0x25, 0x78,
	0x0c, 0x1b, 0xca, 0x58, 0x67, 0x7c, 0x71, 0x87, 0x3b, 0x7c, 0xc9, 0x50, 0xdf, 0xde, 0xa2, 0x07,
	0x9b, 0x13, 0x23, 0x75, 0x62, 0x2c, 0xda, 0xf2, 0xcb, 0x92, 0xf5, 0x00, 0x2d, 0xe5, 0x7a, 0xaa,
	0x4a, 0x81, 0xec, 0xa7, 0x03, 0xfb, 0x4b, 0xb0, 0x55, 0x74, 0x01, 0xdd, 0xb4, 0x80, 0x49, 0x79,
	0xce, 0xa0, 0x75, 0xe2, 0x0e, 0x9f, 0xf2, 0x86, 0xc6, 0x12, 0x23, 0x75, 0x15, 0x6b, 0x39, 0xf7,
	0x6b, 0x56, 0xff, 0x03, 0xec, 0x2c, 0x5f, 0xe2, 0x1e, 0xb4, 0xbe, 0xd2, 0xdc, 0x9a, 0x91, 0x1f,
	0xf1, 0x18, 0x3a, 0x77, 0x22, 0xca, 0xc8, 0x7e, 0xd5, 0x1e, 0xff, 0xc3, 0x19, 0xbf, 0xb8, 0x3e,
	0x5f, 0x3f, 0x73, 0xd8, 0x00, 0xe0, 0x3a, 0x09, 0xef, 0x73, 0xf6, 0x09, 0xb8, 0xa6, 0xc3, 0x7e,
	0x03, 0x42, 0x3b, 0x9a, 0xc6, 0x64, 0x93, 0x64, 0xce, 0x6c, 0x02, 0x07, 0xd5, 0x88, 0x34, 0x12,
	0xc1, 0xff, 0x87, 0x01, 0x9f, 0xc3, 0x9e, 0x26, 0x39, 0x9b, 0xc6, 0x42, 0xd3, 0x58, 0x4d, 0xc3,
	0x58, 0x44, 0xf6, 0xa7, 0x77, 0x2b, 0x7c, 0x64, 0xe0, 0xe1, 0xaf, 0x56, 0x35, 0xe6, 0x46, 0xc4,
	0x22, 0x24, 0x39, 0x22, 0x79, 0x37, 0x0d, 0x08, 0xcf, 0xab, 0x7f, 0x42, 0x91, 0x45, 0x3c, 0xe0,
	0x4d, 0xd9, 0xec, 0xaf, 0x38, 0xc1, 0xd6, 0x16, 0xb8, 0x45, 0x1a, 0x6b, 0xee, 0x52, 0x3a, 0x1b,
	0xb9, 0x6f, 0x00, 0xea, 0x74, 0x22, 0xf2, 0x95, 0xa8, 0xfe, 0x65, 0xa2, 0xbb, 0xb0, 0x73, 0xdc,
	0xe7, 0xab, 0x09, 0xea, 0xf7, 0x9a, 0x62, 0xc1, 0xd6, 0xf0, 0x65, 0x35, 0xf1, 0x3a, 0x09, 0xd1,
	0xe5, 0xf5, 0xee, 0xfa, 0xdb, 0x7c, 0x61, 0x4d, 0x6c, 0xed, 0x95, 0x83, 0x6f, 0x61, 0xdb, 0x36,
	0x7f, 0x12, 0x3a, 0xf8, 0x82, 0x87, 0xbc, 0x78, 0xda, 0x78, 0xf9, 0xb4, 0xf1, 0xab, 0xfc, 0x69,
	0x6b, 0x12, 0x69, 0xd8, 0x3b, 0xcb, 0x4b, 0xc5, 0x43, 0xde, 0xb8, 0xe5, 0x26, 0xfe, 0xed, 0x86,
	0x99, 0xf1, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xb1, 0x63, 0xc9, 0x59, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessManagerServiceClient is the client API for ProcessManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessManagerServiceClient interface {
	ProcessCreate(ctx context.Context, in *ProcessCreateRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	ProcessDelete(ctx context.Context, in *ProcessDeleteRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	ProcessGet(ctx context.Context, in *ProcessGetRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	ProcessList(ctx context.Context, in *ProcessListRequest, opts ...grpc.CallOption) (*ProcessListResponse, error)
	ProcessLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (ProcessManagerService_ProcessLogClient, error)
	ProcessWatch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ProcessManagerService_ProcessWatchClient, error)
	ProcessReplace(ctx context.Context, in *ProcessReplaceRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
}

type processManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewProcessManagerServiceClient(cc *grpc.ClientConn) ProcessManagerServiceClient {
	return &processManagerServiceClient{cc}
}

func (c *processManagerServiceClient) ProcessCreate(ctx context.Context, in *ProcessCreateRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/ProcessManagerService/ProcessCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerServiceClient) ProcessDelete(ctx context.Context, in *ProcessDeleteRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/ProcessManagerService/ProcessDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerServiceClient) ProcessGet(ctx context.Context, in *ProcessGetRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/ProcessManagerService/ProcessGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerServiceClient) ProcessList(ctx context.Context, in *ProcessListRequest, opts ...grpc.CallOption) (*ProcessListResponse, error) {
	out := new(ProcessListResponse)
	err := c.cc.Invoke(ctx, "/ProcessManagerService/ProcessList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processManagerServiceClient) ProcessLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (ProcessManagerService_ProcessLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessManagerService_serviceDesc.Streams[0], "/ProcessManagerService/ProcessLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &processManagerServiceProcessLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessManagerService_ProcessLogClient interface {
	Recv() (*LogResponse, error)
	grpc.ClientStream
}

type processManagerServiceProcessLogClient struct {
	grpc.ClientStream
}

func (x *processManagerServiceProcessLogClient) Recv() (*LogResponse, error) {
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processManagerServiceClient) ProcessWatch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ProcessManagerService_ProcessWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessManagerService_serviceDesc.Streams[1], "/ProcessManagerService/ProcessWatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &processManagerServiceProcessWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessManagerService_ProcessWatchClient interface {
	Recv() (*ProcessResponse, error)
	grpc.ClientStream
}

type processManagerServiceProcessWatchClient struct {
	grpc.ClientStream
}

func (x *processManagerServiceProcessWatchClient) Recv() (*ProcessResponse, error) {
	m := new(ProcessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processManagerServiceClient) ProcessReplace(ctx context.Context, in *ProcessReplaceRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/ProcessManagerService/ProcessReplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessManagerServiceServer is the server API for ProcessManagerService service.
type ProcessManagerServiceServer interface {
	ProcessCreate(context.Context, *ProcessCreateRequest) (*ProcessResponse, error)
	ProcessDelete(context.Context, *ProcessDeleteRequest) (*ProcessResponse, error)
	ProcessGet(context.Context, *ProcessGetRequest) (*ProcessResponse, error)
	ProcessList(context.Context, *ProcessListRequest) (*ProcessListResponse, error)
	ProcessLog(*LogRequest, ProcessManagerService_ProcessLogServer) error
	ProcessWatch(*empty.Empty, ProcessManagerService_ProcessWatchServer) error
	ProcessReplace(context.Context, *ProcessReplaceRequest) (*ProcessResponse, error)
}

// UnimplementedProcessManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProcessManagerServiceServer struct {
}

func (*UnimplementedProcessManagerServiceServer) ProcessCreate(ctx context.Context, req *ProcessCreateRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessCreate not implemented")
}
func (*UnimplementedProcessManagerServiceServer) ProcessDelete(ctx context.Context, req *ProcessDeleteRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessDelete not implemented")
}
func (*UnimplementedProcessManagerServiceServer) ProcessGet(ctx context.Context, req *ProcessGetRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessGet not implemented")
}
func (*UnimplementedProcessManagerServiceServer) ProcessList(ctx context.Context, req *ProcessListRequest) (*ProcessListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessList not implemented")
}
func (*UnimplementedProcessManagerServiceServer) ProcessLog(req *LogRequest, srv ProcessManagerService_ProcessLogServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessLog not implemented")
}
func (*UnimplementedProcessManagerServiceServer) ProcessWatch(req *empty.Empty, srv ProcessManagerService_ProcessWatchServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessWatch not implemented")
}
func (*UnimplementedProcessManagerServiceServer) ProcessReplace(ctx context.Context, req *ProcessReplaceRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessReplace not implemented")
}

func RegisterProcessManagerServiceServer(s *grpc.Server, srv ProcessManagerServiceServer) {
	s.RegisterService(&_ProcessManagerService_serviceDesc, srv)
}

func _ProcessManagerService_ProcessCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServiceServer).ProcessCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessManagerService/ProcessCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServiceServer).ProcessCreate(ctx, req.(*ProcessCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManagerService_ProcessDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServiceServer).ProcessDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessManagerService/ProcessDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServiceServer).ProcessDelete(ctx, req.(*ProcessDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManagerService_ProcessGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServiceServer).ProcessGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessManagerService/ProcessGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServiceServer).ProcessGet(ctx, req.(*ProcessGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManagerService_ProcessList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServiceServer).ProcessList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessManagerService/ProcessList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServiceServer).ProcessList(ctx, req.(*ProcessListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessManagerService_ProcessLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessManagerServiceServer).ProcessLog(m, &processManagerServiceProcessLogServer{stream})
}

type ProcessManagerService_ProcessLogServer interface {
	Send(*LogResponse) error
	grpc.ServerStream
}

type processManagerServiceProcessLogServer struct {
	grpc.ServerStream
}

func (x *processManagerServiceProcessLogServer) Send(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessManagerService_ProcessWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessManagerServiceServer).ProcessWatch(m, &processManagerServiceProcessWatchServer{stream})
}

type ProcessManagerService_ProcessWatchServer interface {
	Send(*ProcessResponse) error
	grpc.ServerStream
}

type processManagerServiceProcessWatchServer struct {
	grpc.ServerStream
}

func (x *processManagerServiceProcessWatchServer) Send(m *ProcessResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessManagerService_ProcessReplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessReplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessManagerServiceServer).ProcessReplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessManagerService/ProcessReplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessManagerServiceServer).ProcessReplace(ctx, req.(*ProcessReplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProcessManagerService",
	HandlerType: (*ProcessManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessCreate",
			Handler:    _ProcessManagerService_ProcessCreate_Handler,
		},
		{
			MethodName: "ProcessDelete",
			Handler:    _ProcessManagerService_ProcessDelete_Handler,
		},
		{
			MethodName: "ProcessGet",
			Handler:    _ProcessManagerService_ProcessGet_Handler,
		},
		{
			MethodName: "ProcessList",
			Handler:    _ProcessManagerService_ProcessList_Handler,
		},
		{
			MethodName: "ProcessReplace",
			Handler:    _ProcessManagerService_ProcessReplace_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessLog",
			Handler:       _ProcessManagerService_ProcessLog_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProcessWatch",
			Handler:       _ProcessManagerService_ProcessWatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}
