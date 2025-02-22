// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: example/library/locker.svc.proto

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
	LockerService_Add_FullMethodName   = "/example.library.LockerService/Add"
	LockerService_Get_FullMethodName   = "/example.library.LockerService/Get"
	LockerService_Patch_FullMethodName = "/example.library.LockerService/Patch"
	LockerService_Erase_FullMethodName = "/example.library.LockerService/Erase"
)

// LockerServiceClient is the client API for LockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LockerServiceClient interface {
	Add(ctx context.Context, in *LockerAddRequest, opts ...grpc.CallOption) (*Locker, error)
	Get(ctx context.Context, in *LockerGetRequest, opts ...grpc.CallOption) (*Locker, error)
	Patch(ctx context.Context, in *LockerPatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Erase(ctx context.Context, in *LockerGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type lockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLockerServiceClient(cc grpc.ClientConnInterface) LockerServiceClient {
	return &lockerServiceClient{cc}
}

func (c *lockerServiceClient) Add(ctx context.Context, in *LockerAddRequest, opts ...grpc.CallOption) (*Locker, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Locker)
	err := c.cc.Invoke(ctx, LockerService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockerServiceClient) Get(ctx context.Context, in *LockerGetRequest, opts ...grpc.CallOption) (*Locker, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Locker)
	err := c.cc.Invoke(ctx, LockerService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockerServiceClient) Patch(ctx context.Context, in *LockerPatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LockerService_Patch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockerServiceClient) Erase(ctx context.Context, in *LockerGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LockerService_Erase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockerServiceServer is the server API for LockerService service.
// All implementations must embed UnimplementedLockerServiceServer
// for forward compatibility.
type LockerServiceServer interface {
	Add(context.Context, *LockerAddRequest) (*Locker, error)
	Get(context.Context, *LockerGetRequest) (*Locker, error)
	Patch(context.Context, *LockerPatchRequest) (*emptypb.Empty, error)
	Erase(context.Context, *LockerGetRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLockerServiceServer()
}

// UnimplementedLockerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLockerServiceServer struct{}

func (UnimplementedLockerServiceServer) Add(context.Context, *LockerAddRequest) (*Locker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedLockerServiceServer) Get(context.Context, *LockerGetRequest) (*Locker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLockerServiceServer) Patch(context.Context, *LockerPatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedLockerServiceServer) Erase(context.Context, *LockerGetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Erase not implemented")
}
func (UnimplementedLockerServiceServer) mustEmbedUnimplementedLockerServiceServer() {}
func (UnimplementedLockerServiceServer) testEmbeddedByValue()                       {}

// UnsafeLockerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LockerServiceServer will
// result in compilation errors.
type UnsafeLockerServiceServer interface {
	mustEmbedUnimplementedLockerServiceServer()
}

func RegisterLockerServiceServer(s grpc.ServiceRegistrar, srv LockerServiceServer) {
	// If the following call pancis, it indicates UnimplementedLockerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LockerService_ServiceDesc, srv)
}

func _LockerService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockerAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockerServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LockerService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockerServiceServer).Add(ctx, req.(*LockerAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LockerService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockerServiceServer).Get(ctx, req.(*LockerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockerService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockerPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockerServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LockerService_Patch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockerServiceServer).Patch(ctx, req.(*LockerPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockerService_Erase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockerServiceServer).Erase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LockerService_Erase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockerServiceServer).Erase(ctx, req.(*LockerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LockerService_ServiceDesc is the grpc.ServiceDesc for LockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.library.LockerService",
	HandlerType: (*LockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _LockerService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LockerService_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _LockerService_Patch_Handler,
		},
		{
			MethodName: "Erase",
			Handler:    _LockerService_Erase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/library/locker.svc.proto",
}
