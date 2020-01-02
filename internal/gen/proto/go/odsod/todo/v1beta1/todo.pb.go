// Code generated by protoc-gen-go. DO NOT EDIT.
// source: odsod/todo/v1beta1/todo.proto

package todov1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A todo resource.
type Todo struct {
	// Resource name of the todo.
	// For example: "todos/1234"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Indicates if the todo is deleted.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// The creation timestamp of the todo.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the todo.
	//
	// Note: update_time is updated when create/update/delete operation is
	// performed.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the todo.
	DeleteTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The text context of the todo.
	// For example: "Do chores."
	Text string `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	// Flag for marking the todo as done.
	Done bool `protobuf:"varint,7,opt,name=done,proto3" json:"done,omitempty"`
	// Resource name of the owner of the todo.
	// For example: "users/1234"
	UserId               string   `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3390c545f097bf6, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Todo) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Todo) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Todo) GetDeleteTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Todo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*Todo)(nil), "odsod.todo.v1beta1.Todo")
}

func init() { proto.RegisterFile("odsod/todo/v1beta1/todo.proto", fileDescriptor_c3390c545f097bf6) }

var fileDescriptor_c3390c545f097bf6 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0x84, 0x40,
	0x10, 0x85, 0xc3, 0x8a, 0x70, 0xb7, 0x24, 0x16, 0x5b, 0xe8, 0xe6, 0x12, 0x23, 0xb1, 0xa2, 0x5a,
	0x72, 0x5a, 0xda, 0x5d, 0x6c, 0xac, 0xb8, 0x10, 0x62, 0x61, 0x2e, 0xb9, 0x00, 0xb3, 0x5e, 0x48,
	0x0e, 0x87, 0xc0, 0x62, 0xfc, 0x3d, 0x56, 0xc6, 0x9f, 0xe2, 0xaf, 0x32, 0x3b, 0x0b, 0x36, 0x16,
	0xd7, 0xed, 0xbc, 0xf7, 0x3e, 0x1e, 0x33, 0xfc, 0x1a, 0x61, 0x40, 0x48, 0x0d, 0x02, 0xa6, 0xef,
	0xeb, 0x4a, 0x9b, 0x72, 0x4d, 0x83, 0xea, 0x7a, 0x34, 0x28, 0x04, 0xd9, 0x8a, 0x94, 0xc9, 0x5e,
	0xdd, 0x1c, 0x10, 0x0f, 0x47, 0x9d, 0x52, 0xa2, 0x1a, 0x5f, 0x53, 0xd3, 0xb4, 0x7a, 0x30, 0x65,
	0xdb, 0x39, 0xe8, 0xf6, 0x8b, 0x71, 0xbf, 0x40, 0x40, 0x71, 0xc1, 0x59, 0x03, 0xd2, 0x8b, 0xbd,
	0x64, 0x99, 0xb3, 0x06, 0x84, 0xe4, 0x21, 0xe8, 0xa3, 0x36, 0x1a, 0x24, 0x8b, 0xbd, 0x64, 0x91,
	0xcf, 0xa3, 0x78, 0xe0, 0x51, 0xdd, 0xeb, 0xd2, 0xe8, 0xbd, 0xfd, 0x98, 0x3c, 0x8b, 0xbd, 0x24,
	0xba, 0x5b, 0x29, 0xd7, 0xa4, 0xe6, 0x26, 0x55, 0xcc, 0x4d, 0x39, 0x77, 0x71, 0x2b, 0x58, 0x78,
	0xec, 0xe0, 0x0f, 0xf6, 0x4f, 0xc3, 0x2e, 0x3e, 0xc3, 0xee, 0x27, 0x1c, 0x7c, 0x7e, 0x1a, 0x76,
	0x71, 0x82, 0x05, 0xf7, 0x8d, 0xfe, 0x30, 0x32, 0xa0, 0x15, 0xe9, 0x6d, 0x35, 0xc0, 0x37, 0x2d,
	0x43, 0xda, 0x90, 0xde, 0xe2, 0x8a, 0x87, 0xe3, 0xa0, 0xfb, 0x7d, 0x03, 0x72, 0x41, 0xd1, 0xc0,
	0x8e, 0x4f, 0xb0, 0xa9, 0xf9, 0x65, 0x8d, 0xad, 0xfa, 0x7f, 0xe5, 0xcd, 0xd2, 0x5e, 0x70, 0x6b,
	0xeb, 0xb7, 0xde, 0x4b, 0x64, 0xad, 0xc9, 0xf9, 0x64, 0x7e, 0x91, 0x3d, 0x66, 0xdf, 0x4c, 0x64,
	0x44, 0xd9, 0x9c, 0x7a, 0x76, 0xde, 0xcf, 0x24, 0xee, 0xac, 0xb8, 0x9b, 0xc4, 0x2a, 0xa0, 0x2d,
	0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x11, 0x23, 0x5c, 0xec, 0x01, 0x00, 0x00,
}
