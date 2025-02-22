// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: example/library/book.svc.proto

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
	BookService_Add_FullMethodName   = "/example.library.BookService/Add"
	BookService_Get_FullMethodName   = "/example.library.BookService/Get"
	BookService_Patch_FullMethodName = "/example.library.BookService/Patch"
	BookService_Erase_FullMethodName = "/example.library.BookService/Erase"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	Add(ctx context.Context, in *BookAddRequest, opts ...grpc.CallOption) (*Book, error)
	Get(ctx context.Context, in *BookGetRequest, opts ...grpc.CallOption) (*Book, error)
	Patch(ctx context.Context, in *BookPatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Erase(ctx context.Context, in *BookGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Add(ctx context.Context, in *BookAddRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Get(ctx context.Context, in *BookGetRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Patch(ctx context.Context, in *BookPatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BookService_Patch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Erase(ctx context.Context, in *BookGetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BookService_Erase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility.
type BookServiceServer interface {
	Add(context.Context, *BookAddRequest) (*Book, error)
	Get(context.Context, *BookGetRequest) (*Book, error)
	Patch(context.Context, *BookPatchRequest) (*emptypb.Empty, error)
	Erase(context.Context, *BookGetRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServiceServer struct{}

func (UnimplementedBookServiceServer) Add(context.Context, *BookAddRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedBookServiceServer) Get(context.Context, *BookGetRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookServiceServer) Patch(context.Context, *BookPatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedBookServiceServer) Erase(context.Context, *BookGetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Erase not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}
func (UnimplementedBookServiceServer) testEmbeddedByValue()                     {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Add(ctx, req.(*BookAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*BookGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Patch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Patch(ctx, req.(*BookPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Erase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Erase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Erase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Erase(ctx, req.(*BookGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.library.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _BookService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _BookService_Patch_Handler,
		},
		{
			MethodName: "Erase",
			Handler:    _BookService_Erase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/library/book.svc.proto",
}
