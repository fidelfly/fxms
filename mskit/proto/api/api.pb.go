// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type IdResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdResponse) Reset()         { *m = IdResponse{} }
func (m *IdResponse) String() string { return proto.CompactTextString(m) }
func (*IdResponse) ProtoMessage()    {}
func (*IdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *IdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdResponse.Unmarshal(m, b)
}
func (m *IdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdResponse.Marshal(b, m, deterministic)
}
func (m *IdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdResponse.Merge(m, src)
}
func (m *IdResponse) XXX_Size() int {
	return xxx_messageInfo_IdResponse.Size(m)
}
func (m *IdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdResponse proto.InternalMessageInfo

func (m *IdResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cols                 []string `protobuf:"bytes,2,rep,name=cols,proto3" json:"cols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateInfo) Reset()         { *m = UpdateInfo{} }
func (m *UpdateInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateInfo) ProtoMessage()    {}
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *UpdateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInfo.Unmarshal(m, b)
}
func (m *UpdateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInfo.Marshal(b, m, deterministic)
}
func (m *UpdateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfo.Merge(m, src)
}
func (m *UpdateInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateInfo.Size(m)
}
func (m *UpdateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfo proto.InternalMessageInfo

func (m *UpdateInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateInfo) GetCols() []string {
	if m != nil {
		return m.Cols
	}
	return nil
}

type Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*IdResponse)(nil), "mskit.api.IdResponse")
	proto.RegisterType((*UpdateInfo)(nil), "mskit.api.UpdateInfo")
	proto.RegisterType((*Request)(nil), "mskit.api.Request")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x46, 0x69, 0xfa, 0xf3, 0x4b, 0xee, 0xe0, 0x90, 0xa9, 0x83, 0x43, 0xe9, 0x54, 0x04, 0x1b,
	0xc1, 0x37, 0x70, 0xeb, 0x1a, 0x70, 0x71, 0x4b, 0x9b, 0x1b, 0xbd, 0xd8, 0x34, 0xd1, 0xa4, 0xa0,
	0x6f, 0x2f, 0x04, 0x04, 0x71, 0x3b, 0x7c, 0xdf, 0x19, 0x0e, 0x70, 0x1d, 0xa8, 0x0b, 0x0f, 0x9f,
	0xbc, 0xe0, 0x2e, 0xde, 0x28, 0x75, 0x3a, 0x50, 0xb3, 0x01, 0xe8, 0x8d, 0xc2, 0x18, 0xfc, 0x1c,
	0x51, 0xac, 0x81, 0x91, 0xa9, 0x8a, 0xba, 0x68, 0x4b, 0xc5, 0xc8, 0x34, 0x7b, 0x80, 0x53, 0x30,
	0x3a, 0x61, 0x3f, 0x5b, 0xff, 0xfb, 0x0a, 0x01, 0x7f, 0xa3, 0x9f, 0x62, 0xc5, 0xea, 0xb2, 0xe5,
	0x2a, 0x73, 0xb3, 0x83, 0x95, 0xc2, 0xfb, 0x82, 0x31, 0x7d, 0xe9, 0xfc, 0xa3, 0xcf, 0xda, 0x61,
	0xc5, 0xf2, 0x92, 0xf9, 0xb8, 0x3d, 0xb7, 0x17, 0x4a, 0xd7, 0x65, 0xe8, 0x46, 0xef, 0xa4, 0x25,
	0x83, 0x93, 0x9d, 0x5e, 0xd2, 0x3e, 0x5d, 0x94, 0x39, 0x52, 0xe6, 0x62, 0xa9, 0x03, 0x0d, 0xff,
	0x19, 0x0f, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xe0, 0x25, 0x5c, 0xc9, 0x00, 0x00, 0x00,
}
