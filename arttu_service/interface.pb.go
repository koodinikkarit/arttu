// Code generated by protoc-gen-go.
// source: interface.proto
// DO NOT EDIT!

/*
Package ArttuService is a generated protocol buffer package.

It is generated from these files:
	interface.proto

It has these top-level messages:
	Interface
	AddInterfaceRequest
	AddInterfaceResponse
*/
package ArttuService

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

type Interface struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mac   []byte `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
	Index uint32 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	Mtu   uint32 `protobuf:"varint,4,opt,name=mtu" json:"mtu,omitempty"`
	Flags uint32 `protobuf:"varint,5,opt,name=flags" json:"flags,omitempty"`
}

func (m *Interface) Reset()                    { *m = Interface{} }
func (m *Interface) String() string            { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()               {}
func (*Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Interface) GetMac() []byte {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *Interface) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Interface) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Interface) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type AddInterfaceRequest struct {
	ArttuId   uint32     `protobuf:"varint,1,opt,name=arttuId" json:"arttuId,omitempty"`
	Interface *Interface `protobuf:"bytes,2,opt,name=interface" json:"interface,omitempty"`
}

func (m *AddInterfaceRequest) Reset()                    { *m = AddInterfaceRequest{} }
func (m *AddInterfaceRequest) String() string            { return proto.CompactTextString(m) }
func (*AddInterfaceRequest) ProtoMessage()               {}
func (*AddInterfaceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddInterfaceRequest) GetArttuId() uint32 {
	if m != nil {
		return m.ArttuId
	}
	return 0
}

func (m *AddInterfaceRequest) GetInterface() *Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type AddInterfaceResponse struct {
}

func (m *AddInterfaceResponse) Reset()                    { *m = AddInterfaceResponse{} }
func (m *AddInterfaceResponse) String() string            { return proto.CompactTextString(m) }
func (*AddInterfaceResponse) ProtoMessage()               {}
func (*AddInterfaceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Interface)(nil), "ArttuService.Interface")
	proto.RegisterType((*AddInterfaceRequest)(nil), "ArttuService.AddInterfaceRequest")
	proto.RegisterType((*AddInterfaceResponse)(nil), "ArttuService.AddInterfaceResponse")
}

func init() { proto.RegisterFile("interface.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4e, 0x84, 0x30,
	0x18, 0x84, 0x53, 0x01, 0x0d, 0xbf, 0x10, 0x4d, 0x25, 0xda, 0x23, 0xe1, 0xc4, 0x89, 0x83, 0xc6,
	0x07, 0xe0, 0xc8, 0xb5, 0x3e, 0x41, 0xa5, 0x3f, 0x86, 0x44, 0x5a, 0xb6, 0x2d, 0x9b, 0x7d, 0xfc,
	0x4d, 0x4b, 0x80, 0xdd, 0xdb, 0x4c, 0xf3, 0xa5, 0x33, 0xf3, 0xc3, 0xcb, 0xa8, 0x1c, 0x9a, 0x41,
	0xf4, 0xd8, 0xcc, 0x46, 0x3b, 0x4d, 0xb3, 0xd6, 0x38, 0xb7, 0xfc, 0xa0, 0x39, 0x8f, 0x3d, 0x56,
	0x1a, 0xd2, 0x6e, 0x03, 0x28, 0x85, 0x58, 0x89, 0x09, 0x19, 0x29, 0x49, 0x9d, 0xf2, 0xa0, 0xe9,
	0x2b, 0x44, 0x93, 0xe8, 0xd9, 0x43, 0x49, 0xea, 0x8c, 0x7b, 0x49, 0x0b, 0x48, 0x46, 0x25, 0xf1,
	0xc2, 0xa2, 0x92, 0xd4, 0x39, 0x5f, 0x4d, 0xe0, 0xdc, 0xc2, 0xe2, 0xf0, 0xe6, 0xa5, 0xe7, 0x86,
	0x7f, 0xf1, 0x67, 0x59, 0xb2, 0x72, 0xc1, 0x54, 0x03, 0xbc, 0xb5, 0x52, 0xee, 0x99, 0x1c, 0x4f,
	0x0b, 0x5a, 0x47, 0x19, 0x3c, 0x09, 0xdf, 0xab, 0x93, 0x21, 0x3d, 0xe7, 0x9b, 0xa5, 0xdf, 0x90,
	0xee, 0x13, 0x42, 0x8d, 0xe7, 0xcf, 0x8f, 0xe6, 0x76, 0x43, 0x73, 0x7c, 0x76, 0x90, 0xd5, 0x3b,
	0x14, 0xf7, 0x39, 0x76, 0xd6, 0xca, 0xe2, 0xef, 0x63, 0xb8, 0xc2, 0xd7, 0x35, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x31, 0xcf, 0x52, 0x18, 0x01, 0x00, 0x00,
}
