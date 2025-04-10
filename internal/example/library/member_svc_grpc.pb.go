// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: example/library/member_svc.proto

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
	MemberService_Add_FullMethodName   = "/example.library.MemberService/Add"
	MemberService_Get_FullMethodName   = "/example.library.MemberService/Get"
	MemberService_Patch_FullMethodName = "/example.library.MemberService/Patch"
	MemberService_Erase_FullMethodName = "/example.library.MemberService/Erase"
)

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberServiceClient interface {
	Add(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*Member, error)
	Get(ctx context.Context, in *MemberGetRequest, opts ...grpc.CallOption) (*Member, error)
	Patch(ctx context.Context, in *MemberPatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Erase(ctx context.Context, in *MemberGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) Add(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*Member, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Member)
	err := c.cc.Invoke(ctx, MemberService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Get(ctx context.Context, in *MemberGetRequest, opts ...grpc.CallOption) (*Member, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Member)
	err := c.cc.Invoke(ctx, MemberService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Patch(ctx context.Context, in *MemberPatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MemberService_Patch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Erase(ctx context.Context, in *MemberGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MemberService_Erase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations must embed UnimplementedMemberServiceServer
// for forward compatibility.
type MemberServiceServer interface {
	Add(context.Context, *MemberAddRequest) (*Member, error)
	Get(context.Context, *MemberGetRequest) (*Member, error)
	Patch(context.Context, *MemberPatchRequest) (*emptypb.Empty, error)
	Erase(context.Context, *MemberGetRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMemberServiceServer()
}

// UnimplementedMemberServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMemberServiceServer struct{}

func (UnimplementedMemberServiceServer) Add(context.Context, *MemberAddRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMemberServiceServer) Get(context.Context, *MemberGetRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMemberServiceServer) Patch(context.Context, *MemberPatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedMemberServiceServer) Erase(context.Context, *MemberGetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Erase not implemented")
}
func (UnimplementedMemberServiceServer) mustEmbedUnimplementedMemberServiceServer() {}
func (UnimplementedMemberServiceServer) testEmbeddedByValue()                       {}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	// If the following call pancis, it indicates UnimplementedMemberServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Add(ctx, req.(*MemberAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Get(ctx, req.(*MemberGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_Patch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Patch(ctx, req.(*MemberPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Erase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Erase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_Erase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Erase(ctx, req.(*MemberGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.library.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MemberService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MemberService_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _MemberService_Patch_Handler,
		},
		{
			MethodName: "Erase",
			Handler:    _MemberService_Erase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/library/member_svc.proto",
}
