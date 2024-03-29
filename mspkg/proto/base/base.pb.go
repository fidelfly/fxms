// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package base

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

type StringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}
func (*StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (m *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(m, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringValue)(nil), "mspkg.base.StringValue")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x2d, 0x2e, 0xc8, 0x4e, 0xd7, 0x03, 0x89,
	0x28, 0x29, 0x73, 0x71, 0x07, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x87, 0x25, 0xe6, 0x94, 0xa6, 0x0a,
	0x89, 0x70, 0xb1, 0x96, 0x81, 0x18, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x93,
	0x76, 0x94, 0x66, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x5a, 0x66,
	0x4a, 0x6a, 0x4e, 0x5a, 0x4e, 0xa5, 0x7e, 0x5a, 0x45, 0x6e, 0xb1, 0x3e, 0xd8, 0x2c, 0x7d, 0xb0,
	0xc1, 0xfa, 0x20, 0x13, 0x93, 0xd8, 0xc0, 0x6c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0xd3, 0xc0, 0xfe, 0x72, 0x00, 0x00, 0x00,
}
