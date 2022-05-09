// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SampleServiceClient is the client API for SampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleServiceClient interface {
	CreateElement(ctx context.Context, in *CreateElementRequest, opts ...grpc.CallOption) (*CreateElementResponse, error)
	DeleteElement(ctx context.Context, in *DeleteElementRequest, opts ...grpc.CallOption) (*DeleteElementResponse, error)
}

type sampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleServiceClient(cc grpc.ClientConnInterface) SampleServiceClient {
	return &sampleServiceClient{cc}
}

func (c *sampleServiceClient) CreateElement(ctx context.Context, in *CreateElementRequest, opts ...grpc.CallOption) (*CreateElementResponse, error) {
	out := new(CreateElementResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.SampleService/CreateElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleServiceClient) DeleteElement(ctx context.Context, in *DeleteElementRequest, opts ...grpc.CallOption) (*DeleteElementResponse, error) {
	out := new(DeleteElementResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.SampleService/DeleteElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServiceServer is the server API for SampleService service.
// All implementations must embed UnimplementedSampleServiceServer
// for forward compatibility
type SampleServiceServer interface {
	CreateElement(context.Context, *CreateElementRequest) (*CreateElementResponse, error)
	DeleteElement(context.Context, *DeleteElementRequest) (*DeleteElementResponse, error)
	mustEmbedUnimplementedSampleServiceServer()
}

// UnimplementedSampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServiceServer struct {
}

func (UnimplementedSampleServiceServer) CreateElement(context.Context, *CreateElementRequest) (*CreateElementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateElement not implemented")
}
func (UnimplementedSampleServiceServer) DeleteElement(context.Context, *DeleteElementRequest) (*DeleteElementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteElement not implemented")
}
func (UnimplementedSampleServiceServer) mustEmbedUnimplementedSampleServiceServer() {}

// UnsafeSampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServiceServer will
// result in compilation errors.
type UnsafeSampleServiceServer interface {
	mustEmbedUnimplementedSampleServiceServer()
}

func RegisterSampleServiceServer(s grpc.ServiceRegistrar, srv SampleServiceServer) {
	s.RegisterService(&SampleService_ServiceDesc, srv)
}

func _SampleService_CreateElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateElementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).CreateElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.SampleService/CreateElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).CreateElement(ctx, req.(*CreateElementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleService_DeleteElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteElementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).DeleteElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.SampleService/DeleteElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).DeleteElement(ctx, req.(*DeleteElementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleService_ServiceDesc is the grpc.ServiceDesc for SampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "durudex.v1.SampleService",
	HandlerType: (*SampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateElement",
			Handler:    _SampleService_CreateElement_Handler,
		},
		{
			MethodName: "DeleteElement",
			Handler:    _SampleService_DeleteElement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/sample.proto",
}