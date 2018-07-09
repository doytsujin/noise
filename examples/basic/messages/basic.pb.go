// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/basic.proto

package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BasicMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasicMessage) Reset()         { *m = BasicMessage{} }
func (m *BasicMessage) String() string { return proto.CompactTextString(m) }
func (*BasicMessage) ProtoMessage()    {}
func (*BasicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_basic_3a5902b328cb205e, []int{0}
}
func (m *BasicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicMessage.Unmarshal(m, b)
}
func (m *BasicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicMessage.Marshal(b, m, deterministic)
}
func (dst *BasicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicMessage.Merge(dst, src)
}
func (m *BasicMessage) XXX_Size() int {
	return xxx_messageInfo_BasicMessage.Size(m)
}
func (m *BasicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BasicMessage proto.InternalMessageInfo

func (m *BasicMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*BasicMessage)(nil), "messages.BasicMessage")
}

func init() { proto.RegisterFile("messages/basic.proto", fileDescriptor_basic_3a5902b328cb205e) }

var fileDescriptor_basic_3a5902b328cb205e = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0xce, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x80, 0x89, 0x2a, 0x69, 0x70, 0xf1, 0x38, 0x81, 0x24, 0x7c, 0x21, 0x02, 0x42, 0x12,
	0x5c, 0xec, 0x50, 0x39, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x37, 0x89, 0x0d, 0xac,
	0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x8a, 0xd3, 0x62, 0x52, 0x00, 0x00, 0x00,
}
