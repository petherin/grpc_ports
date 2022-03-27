// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortsClient is the client API for Ports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortsClient interface {
	Save(ctx context.Context, in *SavePortRequest, opts ...grpc.CallOption) (*SavePortResponse, error)
	GetPorts(ctx context.Context, in *GetPortsRequest, opts ...grpc.CallOption) (*PortList, error)
}

type portsClient struct {
	cc grpc.ClientConnInterface
}

func NewPortsClient(cc grpc.ClientConnInterface) PortsClient {
	return &portsClient{cc}
}

func (c *portsClient) Save(ctx context.Context, in *SavePortRequest, opts ...grpc.CallOption) (*SavePortResponse, error) {
	out := new(SavePortResponse)
	err := c.cc.Invoke(ctx, "/proto.Ports/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portsClient) GetPorts(ctx context.Context, in *GetPortsRequest, opts ...grpc.CallOption) (*PortList, error) {
	out := new(PortList)
	err := c.cc.Invoke(ctx, "/proto.Ports/GetPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortsServer is the server API for Ports service.
// All implementations must embed UnimplementedPortsServer
// for forward compatibility
type PortsServer interface {
	Save(context.Context, *SavePortRequest) (*SavePortResponse, error)
	GetPorts(context.Context, *GetPortsRequest) (*PortList, error)
	mustEmbedUnimplementedPortsServer()
}

// UnimplementedPortsServer must be embedded to have forward compatible implementations.
type UnimplementedPortsServer struct {
}

func (UnimplementedPortsServer) Save(context.Context, *SavePortRequest) (*SavePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedPortsServer) GetPorts(context.Context, *GetPortsRequest) (*PortList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPorts not implemented")
}
func (UnimplementedPortsServer) mustEmbedUnimplementedPortsServer() {}

// UnsafePortsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortsServer will
// result in compilation errors.
type UnsafePortsServer interface {
	mustEmbedUnimplementedPortsServer()
}

func RegisterPortsServer(s grpc.ServiceRegistrar, srv PortsServer) {
	s.RegisterService(&Ports_ServiceDesc, srv)
}

func _Ports_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ports/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsServer).Save(ctx, req.(*SavePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ports_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ports/GetPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsServer).GetPorts(ctx, req.(*GetPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ports_ServiceDesc is the grpc.ServiceDesc for Ports service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ports_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ports",
	HandlerType: (*PortsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Ports_Save_Handler,
		},
		{
			MethodName: "GetPorts",
			Handler:    _Ports_GetPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portsvc/proto/ports.proto",
}
