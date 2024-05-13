// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc2
// source: blog_service.proto

package blog

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
	BlogService_GetById_FullMethodName = "/blog.BlogService/GetById"
	BlogService_GetAll_FullMethodName  = "/blog.BlogService/GetAll"
	BlogService_Create_FullMethodName  = "/blog.BlogService/Create"
	BlogService_Update_FullMethodName  = "/blog.BlogService/Update"
	BlogService_Delete_FullMethodName  = "/blog.BlogService/Delete"
)

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*BlogResponse, error)
	GetAll(ctx context.Context, in *NumArray, opts ...grpc.CallOption) (*MultiBlogResponse, error)
	Create(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*EmptyResponse, error)
	Update(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogResponse, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, BlogService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetAll(ctx context.Context, in *NumArray, opts ...grpc.CallOption) (*MultiBlogResponse, error) {
	out := new(MultiBlogResponse)
	err := c.cc.Invoke(ctx, BlogService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Create(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BlogService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Update(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, BlogService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BlogService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	GetById(context.Context, *Id) (*BlogResponse, error)
	GetAll(context.Context, *NumArray) (*MultiBlogResponse, error)
	Create(context.Context, *Blog) (*EmptyResponse, error)
	Update(context.Context, *Blog) (*BlogResponse, error)
	Delete(context.Context, *Id) (*EmptyResponse, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) GetById(context.Context, *Id) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBlogServiceServer) GetAll(context.Context, *NumArray) (*MultiBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBlogServiceServer) Create(context.Context, *Blog) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBlogServiceServer) Update(context.Context, *Blog) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBlogServiceServer) Delete(context.Context, *Id) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetAll(ctx, req.(*NumArray))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Create(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Update(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _BlogService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BlogService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BlogService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BlogService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BlogService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog_service.proto",
}
