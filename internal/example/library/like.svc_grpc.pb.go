// Code generated by "proto-orm-gen-ent". DO NOT EDIT

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: example/library/like.svc.proto

package library

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LikeService_Add_FullMethodName   = "/example.library.LikeService/Add"
	LikeService_Get_FullMethodName   = "/example.library.LikeService/Get"
	LikeService_Patch_FullMethodName = "/example.library.LikeService/Patch"
	LikeService_Erase_FullMethodName = "/example.library.LikeService/Erase"
)

// LikeServiceClient is the client API for LikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeServiceClient interface {
	Add(ctx context.Context, in *LikeAddRequest, opts ...grpc.CallOption) (*Like, error)
	Get(ctx context.Context, in *LikeGetRequest, opts ...grpc.CallOption) (*Like, error)
	Patch(ctx context.Context, in *LikePatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Erase(ctx context.Context, in *LikeGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type likeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeServiceClient(cc grpc.ClientConnInterface) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) Add(ctx context.Context, in *LikeAddRequest, opts ...grpc.CallOption) (*Like, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Like)
	err := c.cc.Invoke(ctx, LikeService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Get(ctx context.Context, in *LikeGetRequest, opts ...grpc.CallOption) (*Like, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Like)
	err := c.cc.Invoke(ctx, LikeService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Patch(ctx context.Context, in *LikePatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LikeService_Patch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Erase(ctx context.Context, in *LikeGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LikeService_Erase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServiceServer is the server API for LikeService service.
// All implementations must embed UnimplementedLikeServiceServer
// for forward compatibility.
type LikeServiceServer interface {
	Add(context.Context, *LikeAddRequest) (*Like, error)
	Get(context.Context, *LikeGetRequest) (*Like, error)
	Patch(context.Context, *LikePatchRequest) (*emptypb.Empty, error)
	Erase(context.Context, *LikeGetRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLikeServiceServer()
}

// UnimplementedLikeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLikeServiceServer struct{}

func (UnimplementedLikeServiceServer) Add(context.Context, *LikeAddRequest) (*Like, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedLikeServiceServer) Get(context.Context, *LikeGetRequest) (*Like, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLikeServiceServer) Patch(context.Context, *LikePatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedLikeServiceServer) Erase(context.Context, *LikeGetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Erase not implemented")
}
func (UnimplementedLikeServiceServer) mustEmbedUnimplementedLikeServiceServer() {}
func (UnimplementedLikeServiceServer) testEmbeddedByValue()                     {}

// UnsafeLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServiceServer will
// result in compilation errors.
type UnsafeLikeServiceServer interface {
	mustEmbedUnimplementedLikeServiceServer()
}

func RegisterLikeServiceServer(s grpc.ServiceRegistrar, srv LikeServiceServer) {
	// If the following call pancis, it indicates UnimplementedLikeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LikeService_ServiceDesc, srv)
}

func _LikeService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LikeService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Add(ctx, req.(*LikeAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LikeService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Get(ctx, req.(*LikeGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LikeService_Patch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Patch(ctx, req.(*LikePatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Erase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Erase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LikeService_Erase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Erase(ctx, req.(*LikeGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LikeService_ServiceDesc is the grpc.ServiceDesc for LikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.library.LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _LikeService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LikeService_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _LikeService_Patch_Handler,
		},
		{
			MethodName: "Erase",
			Handler:    _LikeService_Erase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/library/like.svc.proto",
}