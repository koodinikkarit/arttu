// Code generated by protoc-gen-go.
// source: computer.proto
// DO NOT EDIT!

/*
Package WompattiService is a generated protocol buffer package.

It is generated from these files:
	computer.proto

It has these top-level messages:
	Computer
	CreateComputerRequest
	CreateComputerResponse
	EditComputerRequest
	EditComputerResponse
	FetchComputersRequest
	FetchComputerByIdRequest
	FetchComputerByIdResponse
	RemoveComputerRequest
	RemoveComputerResponse
	PassiveComputerRequest
	PassiveComputerResponse
	WakeupRequest
	WakeupResponse
*/
package WompattiService

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

type Computer struct {
	Id      uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ArttuId uint32 `protobuf:"varint,2,opt,name=arttuId" json:"arttuId,omitempty"`
}

func (m *Computer) Reset()                    { *m = Computer{} }
func (m *Computer) String() string            { return proto.CompactTextString(m) }
func (*Computer) ProtoMessage()               {}
func (*Computer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Computer) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Computer) GetArttuId() uint32 {
	if m != nil {
		return m.ArttuId
	}
	return 0
}

type CreateComputerRequest struct {
}

func (m *CreateComputerRequest) Reset()                    { *m = CreateComputerRequest{} }
func (m *CreateComputerRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateComputerRequest) ProtoMessage()               {}
func (*CreateComputerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateComputerResponse struct {
	Computer *Computer `protobuf:"bytes,1,opt,name=computer" json:"computer,omitempty"`
}

func (m *CreateComputerResponse) Reset()                    { *m = CreateComputerResponse{} }
func (m *CreateComputerResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateComputerResponse) ProtoMessage()               {}
func (*CreateComputerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateComputerResponse) GetComputer() *Computer {
	if m != nil {
		return m.Computer
	}
	return nil
}

type EditComputerRequest struct {
	Id      uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ArttuId uint32 `protobuf:"varint,2,opt,name=arttuId" json:"arttuId,omitempty"`
}

func (m *EditComputerRequest) Reset()                    { *m = EditComputerRequest{} }
func (m *EditComputerRequest) String() string            { return proto.CompactTextString(m) }
func (*EditComputerRequest) ProtoMessage()               {}
func (*EditComputerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EditComputerRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EditComputerRequest) GetArttuId() uint32 {
	if m != nil {
		return m.ArttuId
	}
	return 0
}

type EditComputerResponse struct {
	Computer *Computer `protobuf:"bytes,1,opt,name=computer" json:"computer,omitempty"`
}

func (m *EditComputerResponse) Reset()                    { *m = EditComputerResponse{} }
func (m *EditComputerResponse) String() string            { return proto.CompactTextString(m) }
func (*EditComputerResponse) ProtoMessage()               {}
func (*EditComputerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EditComputerResponse) GetComputer() *Computer {
	if m != nil {
		return m.Computer
	}
	return nil
}

type FetchComputersRequest struct {
	Offset uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *FetchComputersRequest) Reset()                    { *m = FetchComputersRequest{} }
func (m *FetchComputersRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchComputersRequest) ProtoMessage()               {}
func (*FetchComputersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchComputersRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FetchComputersRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type FetchComputerByIdRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *FetchComputerByIdRequest) Reset()                    { *m = FetchComputerByIdRequest{} }
func (m *FetchComputerByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchComputerByIdRequest) ProtoMessage()               {}
func (*FetchComputerByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FetchComputerByIdRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FetchComputerByIdResponse struct {
	Computer *Computer `protobuf:"bytes,1,opt,name=computer" json:"computer,omitempty"`
}

func (m *FetchComputerByIdResponse) Reset()                    { *m = FetchComputerByIdResponse{} }
func (m *FetchComputerByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchComputerByIdResponse) ProtoMessage()               {}
func (*FetchComputerByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FetchComputerByIdResponse) GetComputer() *Computer {
	if m != nil {
		return m.Computer
	}
	return nil
}

type RemoveComputerRequest struct {
	ComputerId uint32 `protobuf:"varint,1,opt,name=computerId" json:"computerId,omitempty"`
}

func (m *RemoveComputerRequest) Reset()                    { *m = RemoveComputerRequest{} }
func (m *RemoveComputerRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveComputerRequest) ProtoMessage()               {}
func (*RemoveComputerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RemoveComputerRequest) GetComputerId() uint32 {
	if m != nil {
		return m.ComputerId
	}
	return 0
}

type RemoveComputerResponse struct {
}

func (m *RemoveComputerResponse) Reset()                    { *m = RemoveComputerResponse{} }
func (m *RemoveComputerResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveComputerResponse) ProtoMessage()               {}
func (*RemoveComputerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PassiveComputerRequest struct {
}

func (m *PassiveComputerRequest) Reset()                    { *m = PassiveComputerRequest{} }
func (m *PassiveComputerRequest) String() string            { return proto.CompactTextString(m) }
func (*PassiveComputerRequest) ProtoMessage()               {}
func (*PassiveComputerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type PassiveComputerResponse struct {
}

func (m *PassiveComputerResponse) Reset()                    { *m = PassiveComputerResponse{} }
func (m *PassiveComputerResponse) String() string            { return proto.CompactTextString(m) }
func (*PassiveComputerResponse) ProtoMessage()               {}
func (*PassiveComputerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type WakeupRequest struct {
	ComputerId uint32 `protobuf:"varint,1,opt,name=computerId" json:"computerId,omitempty"`
}

func (m *WakeupRequest) Reset()                    { *m = WakeupRequest{} }
func (m *WakeupRequest) String() string            { return proto.CompactTextString(m) }
func (*WakeupRequest) ProtoMessage()               {}
func (*WakeupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *WakeupRequest) GetComputerId() uint32 {
	if m != nil {
		return m.ComputerId
	}
	return 0
}

type WakeupResponse struct {
}

func (m *WakeupResponse) Reset()                    { *m = WakeupResponse{} }
func (m *WakeupResponse) String() string            { return proto.CompactTextString(m) }
func (*WakeupResponse) ProtoMessage()               {}
func (*WakeupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Computer)(nil), "WompattiService.Computer")
	proto.RegisterType((*CreateComputerRequest)(nil), "WompattiService.CreateComputerRequest")
	proto.RegisterType((*CreateComputerResponse)(nil), "WompattiService.CreateComputerResponse")
	proto.RegisterType((*EditComputerRequest)(nil), "WompattiService.EditComputerRequest")
	proto.RegisterType((*EditComputerResponse)(nil), "WompattiService.EditComputerResponse")
	proto.RegisterType((*FetchComputersRequest)(nil), "WompattiService.FetchComputersRequest")
	proto.RegisterType((*FetchComputerByIdRequest)(nil), "WompattiService.FetchComputerByIdRequest")
	proto.RegisterType((*FetchComputerByIdResponse)(nil), "WompattiService.FetchComputerByIdResponse")
	proto.RegisterType((*RemoveComputerRequest)(nil), "WompattiService.RemoveComputerRequest")
	proto.RegisterType((*RemoveComputerResponse)(nil), "WompattiService.RemoveComputerResponse")
	proto.RegisterType((*PassiveComputerRequest)(nil), "WompattiService.PassiveComputerRequest")
	proto.RegisterType((*PassiveComputerResponse)(nil), "WompattiService.PassiveComputerResponse")
	proto.RegisterType((*WakeupRequest)(nil), "WompattiService.WakeupRequest")
	proto.RegisterType((*WakeupResponse)(nil), "WompattiService.WakeupResponse")
}

func init() { proto.RegisterFile("computer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x69, 0xe1, 0xed, 0x5b, 0x1e, 0x69, 0x94, 0xb5, 0x49, 0x93, 0x8b, 0xc8, 0x9e, 0xc4,
	0x43, 0x04, 0x3f, 0xf0, 0x28, 0x58, 0x2a, 0xe4, 0x20, 0x4a, 0x3c, 0xf4, 0x1c, 0x93, 0x29, 0x2e,
	0x1a, 0x37, 0x66, 0x27, 0x05, 0xff, 0x7b, 0x31, 0xc9, 0x16, 0x9b, 0x2a, 0x28, 0x3d, 0xce, 0x3e,
	0x1f, 0xf3, 0x1b, 0x12, 0x38, 0xa9, 0xce, 0x8b, 0x8a, 0xa9, 0x0c, 0x8b, 0x52, 0xb3, 0x16, 0xbb,
	0x73, 0x9d, 0x17, 0x09, 0xb3, 0x7a, 0xa0, 0x72, 0xa9, 0x52, 0x92, 0xe7, 0x18, 0x4e, 0x5b, 0x8b,
	0x70, 0xd0, 0x57, 0x99, 0xdf, 0x3b, 0xec, 0x1d, 0x8d, 0xe2, 0xbe, 0xca, 0x84, 0x8f, 0xff, 0x49,
	0xc9, 0x5c, 0x45, 0x99, 0xdf, 0xaf, 0x1f, 0xed, 0x28, 0x27, 0x70, 0xa7, 0x25, 0x25, 0x4c, 0x36,
	0x1b, 0xd3, 0x5b, 0x45, 0x86, 0xe5, 0x1d, 0xbc, 0xae, 0x60, 0x0a, 0xfd, 0x6a, 0x48, 0x5c, 0x60,
	0x68, 0x59, 0xea, 0x15, 0x3b, 0xa7, 0x41, 0xd8, 0x81, 0x09, 0x57, 0xa1, 0x95, 0x55, 0x5e, 0x61,
	0x7f, 0x96, 0x29, 0xee, 0xec, 0xf9, 0x03, 0xea, 0x2d, 0xc6, 0xeb, 0x05, 0xdb, 0xf1, 0xcc, 0xe0,
	0xde, 0x10, 0xa7, 0x4f, 0x56, 0x32, 0x96, 0xc8, 0xc3, 0x40, 0x2f, 0x16, 0x86, 0xb8, 0xa5, 0x6a,
	0x27, 0x31, 0xc6, 0xbf, 0x17, 0x95, 0x2b, 0x6e, 0xb9, 0x9a, 0x41, 0x1e, 0xc3, 0x5f, 0xab, 0xb9,
	0x7e, 0x8f, 0xb2, 0x1f, 0x6e, 0x93, 0x31, 0x82, 0x6f, 0xbc, 0xdb, 0x9d, 0x71, 0x09, 0x37, 0xa6,
	0x5c, 0x2f, 0xbb, 0x1f, 0x50, 0x1c, 0x00, 0xd6, 0x14, 0x59, 0x88, 0x2f, 0x2f, 0xd2, 0x87, 0xd7,
	0x0d, 0x36, 0x24, 0x9f, 0xca, 0x7d, 0x62, 0x8c, 0xda, 0xe8, 0x94, 0x01, 0x26, 0x1b, 0x4a, 0x1b,
	0x3a, 0xc1, 0x68, 0x9e, 0x3c, 0x53, 0x55, 0xfc, 0x76, 0xff, 0x1e, 0x1c, 0x1b, 0x68, 0x2a, 0x1e,
	0x07, 0xf5, 0x9f, 0x7d, 0xf6, 0x11, 0x00, 0x00, 0xff, 0xff, 0x97, 0x67, 0xfb, 0xce, 0xeb, 0x02,
	0x00, 0x00,
}