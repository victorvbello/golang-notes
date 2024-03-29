// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: greet.proto

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

const (
	GreetSSLService_Greet_FullMethodName = "/greet.GreetSSLService/Greet"
)

// GreetSSLServiceClient is the client API for GreetSSLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetSSLServiceClient interface {
	Greet(ctx context.Context, in *GreetSSLRequest, opts ...grpc.CallOption) (*GreetSSLResponse, error)
}

type greetSSLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetSSLServiceClient(cc grpc.ClientConnInterface) GreetSSLServiceClient {
	return &greetSSLServiceClient{cc}
}

func (c *greetSSLServiceClient) Greet(ctx context.Context, in *GreetSSLRequest, opts ...grpc.CallOption) (*GreetSSLResponse, error) {
	out := new(GreetSSLResponse)
	err := c.cc.Invoke(ctx, GreetSSLService_Greet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetSSLServiceServer is the server API for GreetSSLService service.
// All implementations must embed UnimplementedGreetSSLServiceServer
// for forward compatibility
type GreetSSLServiceServer interface {
	Greet(context.Context, *GreetSSLRequest) (*GreetSSLResponse, error)
	mustEmbedUnimplementedGreetSSLServiceServer()
}

// UnimplementedGreetSSLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetSSLServiceServer struct {
}

func (UnimplementedGreetSSLServiceServer) Greet(context.Context, *GreetSSLRequest) (*GreetSSLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetSSLServiceServer) mustEmbedUnimplementedGreetSSLServiceServer() {}

// UnsafeGreetSSLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetSSLServiceServer will
// result in compilation errors.
type UnsafeGreetSSLServiceServer interface {
	mustEmbedUnimplementedGreetSSLServiceServer()
}

func RegisterGreetSSLServiceServer(s grpc.ServiceRegistrar, srv GreetSSLServiceServer) {
	s.RegisterService(&GreetSSLService_ServiceDesc, srv)
}

func _GreetSSLService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetSSLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetSSLServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetSSLService_Greet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetSSLServiceServer).Greet(ctx, req.(*GreetSSLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetSSLService_ServiceDesc is the grpc.ServiceDesc for GreetSSLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetSSLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetSSLService",
	HandlerType: (*GreetSSLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetSSLService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}
