// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package echov1

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

// OpenAPIServiceClient is the client API for OpenAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenAPIServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type openAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenAPIServiceClient(cc grpc.ClientConnInterface) OpenAPIServiceClient {
	return &openAPIServiceClient{cc}
}

func (c *openAPIServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/echo.v1.OpenAPIService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenAPIServiceServer is the server API for OpenAPIService service.
// All implementations should embed UnimplementedOpenAPIServiceServer
// for forward compatibility
type OpenAPIServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

// UnimplementedOpenAPIServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOpenAPIServiceServer struct {
}

func (UnimplementedOpenAPIServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

// UnsafeOpenAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenAPIServiceServer will
// result in compilation errors.
type UnsafeOpenAPIServiceServer interface {
	mustEmbedUnimplementedOpenAPIServiceServer()
}

func RegisterOpenAPIServiceServer(s grpc.ServiceRegistrar, srv OpenAPIServiceServer) {
	s.RegisterService(&OpenAPIService_ServiceDesc, srv)
}

func _OpenAPIService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenAPIServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.v1.OpenAPIService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenAPIServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenAPIService_ServiceDesc is the grpc.ServiceDesc for OpenAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.v1.OpenAPIService",
	HandlerType: (*OpenAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _OpenAPIService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo/v1/openapi.proto",
}
