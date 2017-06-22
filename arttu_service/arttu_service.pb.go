// Code generated by protoc-gen-go.
// source: arttu_service.proto
// DO NOT EDIT!

/*
Package ArttuService is a generated protocol buffer package.

It is generated from these files:
	arttu_service.proto
	computer.proto
	interface.proto

It has these top-level messages:
	RegisterRequest
	RegisterResponse
	Computer
	InsertComputerBasicInformationRequest
	InsertComputerBasicInformationResponse
	ShutdownComputerRequest
	ShutdownComputerResponse
	Interface
	AddInterfaceRequest
	AddInterfaceResponse
*/
package ArttuService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterRequest struct {
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterResponse struct {
	ArttuId uint32 `protobuf:"varint,1,opt,name=arttuId" json:"arttuId,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterResponse) GetArttuId() uint32 {
	if m != nil {
		return m.ArttuId
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "ArttuService.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "ArttuService.RegisterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Arttu service

type ArttuClient interface {
	InsertComputerBasicInformation(ctx context.Context, in *InsertComputerBasicInformationRequest, opts ...grpc.CallOption) (*InsertComputerBasicInformationResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	AddInterface(ctx context.Context, opts ...grpc.CallOption) (Arttu_AddInterfaceClient, error)
	ShutdownComputer(ctx context.Context, in *ShutdownComputerRequest, opts ...grpc.CallOption) (*ShutdownComputerResponse, error)
}

type arttuClient struct {
	cc *grpc.ClientConn
}

func NewArttuClient(cc *grpc.ClientConn) ArttuClient {
	return &arttuClient{cc}
}

func (c *arttuClient) InsertComputerBasicInformation(ctx context.Context, in *InsertComputerBasicInformationRequest, opts ...grpc.CallOption) (*InsertComputerBasicInformationResponse, error) {
	out := new(InsertComputerBasicInformationResponse)
	err := grpc.Invoke(ctx, "/ArttuService.Arttu/insertComputerBasicInformation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arttuClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/ArttuService.Arttu/register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arttuClient) AddInterface(ctx context.Context, opts ...grpc.CallOption) (Arttu_AddInterfaceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Arttu_serviceDesc.Streams[0], c.cc, "/ArttuService.Arttu/addInterface", opts...)
	if err != nil {
		return nil, err
	}
	x := &arttuAddInterfaceClient{stream}
	return x, nil
}

type Arttu_AddInterfaceClient interface {
	Send(*AddInterfaceRequest) error
	CloseAndRecv() (*AddInterfaceResponse, error)
	grpc.ClientStream
}

type arttuAddInterfaceClient struct {
	grpc.ClientStream
}

func (x *arttuAddInterfaceClient) Send(m *AddInterfaceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *arttuAddInterfaceClient) CloseAndRecv() (*AddInterfaceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddInterfaceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arttuClient) ShutdownComputer(ctx context.Context, in *ShutdownComputerRequest, opts ...grpc.CallOption) (*ShutdownComputerResponse, error) {
	out := new(ShutdownComputerResponse)
	err := grpc.Invoke(ctx, "/ArttuService.Arttu/shutdownComputer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Arttu service

type ArttuServer interface {
	InsertComputerBasicInformation(context.Context, *InsertComputerBasicInformationRequest) (*InsertComputerBasicInformationResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	AddInterface(Arttu_AddInterfaceServer) error
	ShutdownComputer(context.Context, *ShutdownComputerRequest) (*ShutdownComputerResponse, error)
}

func RegisterArttuServer(s *grpc.Server, srv ArttuServer) {
	s.RegisterService(&_Arttu_serviceDesc, srv)
}

func _Arttu_InsertComputerBasicInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertComputerBasicInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArttuServer).InsertComputerBasicInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArttuService.Arttu/InsertComputerBasicInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArttuServer).InsertComputerBasicInformation(ctx, req.(*InsertComputerBasicInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arttu_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArttuServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArttuService.Arttu/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArttuServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arttu_AddInterface_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ArttuServer).AddInterface(&arttuAddInterfaceServer{stream})
}

type Arttu_AddInterfaceServer interface {
	SendAndClose(*AddInterfaceResponse) error
	Recv() (*AddInterfaceRequest, error)
	grpc.ServerStream
}

type arttuAddInterfaceServer struct {
	grpc.ServerStream
}

func (x *arttuAddInterfaceServer) SendAndClose(m *AddInterfaceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *arttuAddInterfaceServer) Recv() (*AddInterfaceRequest, error) {
	m := new(AddInterfaceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Arttu_ShutdownComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownComputerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArttuServer).ShutdownComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArttuService.Arttu/ShutdownComputer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArttuServer).ShutdownComputer(ctx, req.(*ShutdownComputerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Arttu_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ArttuService.Arttu",
	HandlerType: (*ArttuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "insertComputerBasicInformation",
			Handler:    _Arttu_InsertComputerBasicInformation_Handler,
		},
		{
			MethodName: "register",
			Handler:    _Arttu_Register_Handler,
		},
		{
			MethodName: "shutdownComputer",
			Handler:    _Arttu_ShutdownComputer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "addInterface",
			Handler:       _Arttu_AddInterface_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "arttu_service.proto",
}

func init() { proto.RegisterFile("arttu_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x36, 0x88, 0x3f, 0x0c, 0xd5, 0xd6, 0xf1, 0x52, 0x02, 0x16, 0x0d, 0x28, 0x3d, 0x48, 0x0e,
	0xd6, 0x17, 0xa8, 0x9e, 0x82, 0xb7, 0xf4, 0xe4, 0x49, 0xd6, 0x64, 0xaa, 0x7b, 0xe8, 0x6e, 0x9c,
	0x9d, 0xd5, 0xa7, 0xf0, 0x2d, 0x7c, 0x50, 0x31, 0xc9, 0x6a, 0x53, 0x30, 0xe2, 0x31, 0xf3, 0xfd,
	0x67, 0xe1, 0x58, 0xb1, 0x88, 0x7f, 0x70, 0xc4, 0xaf, 0xba, 0xa0, 0xb4, 0x62, 0x2b, 0x16, 0x07,
	0xf3, 0xaf, 0xe3, 0xa2, 0xb9, 0xc5, 0x87, 0x85, 0x5d, 0x55, 0x5e, 0x88, 0x1b, 0x34, 0x1e, 0x6a,
	0x23, 0xc4, 0x4b, 0x15, 0xe8, 0xc9, 0x11, 0x0c, 0x73, 0x7a, 0xd2, 0x4e, 0x88, 0x73, 0x7a, 0xf1,
	0xe4, 0x24, 0xb9, 0x84, 0xd1, 0xcf, 0xc9, 0x55, 0xd6, 0x38, 0xc2, 0x31, 0xec, 0xd5, 0x61, 0x59,
	0x39, 0x8e, 0x4e, 0xa3, 0xe9, 0x41, 0x1e, 0x3e, 0xaf, 0x3e, 0xb6, 0x61, 0xa7, 0x8e, 0xc4, 0xf7,
	0x08, 0x26, 0xda, 0x38, 0x62, 0xb9, 0x6d, 0x43, 0x6f, 0x94, 0xd3, 0x45, 0x66, 0x96, 0x96, 0x57,
	0x4a, 0xb4, 0x35, 0x38, 0x4b, 0xd7, 0xdb, 0xa5, 0x59, 0x2f, 0xbb, 0xed, 0x13, 0x5f, 0xff, 0x4f,
	0xd4, 0x34, 0x4e, 0xb6, 0xf0, 0x0e, 0xf6, 0xb9, 0xdd, 0x81, 0x27, 0x5d, 0x8f, 0x8d, 0xc9, 0xf1,
	0xe4, 0x37, 0xf8, 0xdb, 0xec, 0x1e, 0x06, 0xaa, 0x2c, 0xb3, 0xf0, 0xf7, 0xf0, 0xac, 0xab, 0x98,
	0xaf, 0x61, 0xc1, 0x34, 0xe9, 0xa3, 0x04, 0xe3, 0x69, 0x84, 0x05, 0x8c, 0xdc, 0xb3, 0x97, 0xd2,
	0xbe, 0x99, 0xb0, 0x0a, 0xcf, 0xbb, 0xda, 0xc5, 0x06, 0x1e, 0x22, 0x2e, 0xfe, 0xa2, 0x85, 0x98,
	0xc7, 0xdd, 0xfa, 0xb9, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x2a, 0x50, 0x6e, 0x34,
	0x02, 0x00, 0x00,
}