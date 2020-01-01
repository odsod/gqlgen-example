// Code generated by protoc-gen-go. DO NOT EDIT.
// source: odsod/todo/v1beta1/todo_service.proto

package todopb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// List todos request.
type ListTodosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosRequest) Reset()         { *m = ListTodosRequest{} }
func (m *ListTodosRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodosRequest) ProtoMessage()    {}
func (*ListTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fff1e5a549d09c3, []int{0}
}

func (m *ListTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosRequest.Unmarshal(m, b)
}
func (m *ListTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosRequest.Marshal(b, m, deterministic)
}
func (m *ListTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosRequest.Merge(m, src)
}
func (m *ListTodosRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodosRequest.Size(m)
}
func (m *ListTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosRequest proto.InternalMessageInfo

// List todos response.
type ListTodosResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosResponse) Reset()         { *m = ListTodosResponse{} }
func (m *ListTodosResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodosResponse) ProtoMessage()    {}
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fff1e5a549d09c3, []int{1}
}

func (m *ListTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosResponse.Unmarshal(m, b)
}
func (m *ListTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosResponse.Marshal(b, m, deterministic)
}
func (m *ListTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosResponse.Merge(m, src)
}
func (m *ListTodosResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodosResponse.Size(m)
}
func (m *ListTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosResponse proto.InternalMessageInfo

// Get todo request.
type GetTodoRequest struct {
	// The ID of the todo.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fff1e5a549d09c3, []int{2}
}

func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(m, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Get todo response.
type GetTodoResponse struct {
	// The returned todo.
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fff1e5a549d09c3, []int{3}
}

func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(m, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*ListTodosRequest)(nil), "odsod.todo.v1beta1.ListTodosRequest")
	proto.RegisterType((*ListTodosResponse)(nil), "odsod.todo.v1beta1.ListTodosResponse")
	proto.RegisterType((*GetTodoRequest)(nil), "odsod.todo.v1beta1.GetTodoRequest")
	proto.RegisterType((*GetTodoResponse)(nil), "odsod.todo.v1beta1.GetTodoResponse")
}

func init() {
	proto.RegisterFile("odsod/todo/v1beta1/todo_service.proto", fileDescriptor_2fff1e5a549d09c3)
}

var fileDescriptor_2fff1e5a549d09c3 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x4f, 0x29, 0xce,
	0x4f, 0xd1, 0x2f, 0xc9, 0x4f, 0xc9, 0xd7, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0x04, 0x73,
	0xe2, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84,
	0xc0, 0xca, 0xf4, 0x40, 0x32, 0x7a, 0x50, 0x65, 0x52, 0xb2, 0x38, 0xb4, 0x42, 0xb4, 0x28, 0x09,
	0x71, 0x09, 0xf8, 0x64, 0x16, 0x97, 0x84, 0xe4, 0xa7, 0xe4, 0x17, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x28, 0x09, 0x73, 0x09, 0x22, 0x89, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x29,
	0x70, 0xf1, 0xb9, 0xa7, 0x82, 0xc5, 0xa0, 0xca, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0xec, 0xb9, 0xf8, 0xe1, 0x2a, 0x20, 0x9a,
	0x84, 0x74, 0xb8, 0x58, 0x40, 0x76, 0x81, 0x15, 0x71, 0x1b, 0x49, 0xe8, 0x61, 0xba, 0x4f, 0x0f,
	0xac, 0x1e, 0xac, 0xca, 0x68, 0x37, 0x23, 0x17, 0x37, 0x88, 0x1b, 0x0c, 0xf1, 0x94, 0x50, 0x04,
	0x17, 0x27, 0xdc, 0x1d, 0x42, 0x2a, 0xd8, 0x34, 0xa3, 0x3b, 0x5d, 0x4a, 0x95, 0x80, 0x2a, 0xa8,
	0xbb, 0x82, 0xb8, 0xd8, 0xa1, 0x4e, 0x15, 0x52, 0xc2, 0xa6, 0x03, 0xd5, 0xa7, 0x52, 0xca, 0x78,
	0xd5, 0x40, 0xcc, 0x74, 0x8a, 0xe3, 0x12, 0x4b, 0xce, 0xcf, 0xc5, 0xa2, 0xd2, 0x89, 0x13, 0xa4,
	0x2e, 0x00, 0x14, 0xdc, 0x01, 0x8c, 0x51, 0x6c, 0x20, 0xa9, 0x82, 0xa4, 0x45, 0x4c, 0x2c, 0x21,
	0xfe, 0x2e, 0xfe, 0xab, 0x98, 0x84, 0xfc, 0xc1, 0x1a, 0x40, 0x4a, 0xf4, 0xc2, 0x20, 0x1a, 0x4e,
	0x41, 0x05, 0x63, 0x40, 0x82, 0x31, 0x50, 0xc1, 0x24, 0x36, 0x70, 0x84, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x13, 0xa7, 0x1e, 0xcb, 0x0c, 0x02, 0x00, 0x00,
}