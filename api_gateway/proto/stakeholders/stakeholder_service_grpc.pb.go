// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/stakeholders/stakeholder_service.proto

package stakeholders

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

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	CreatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllPeople(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultiPersonResponse, error)
	GetPersonById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PersonResponse, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.PersonService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetAllPeople(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultiPersonResponse, error) {
	out := new(MultiPersonResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.PersonService/GetAllPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetPersonById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PersonResponse, error) {
	out := new(PersonResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.PersonService/GetPersonById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations should embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	CreatePerson(context.Context, *Person) (*EmptyResponse, error)
	GetAllPeople(context.Context, *Empty) (*MultiPersonResponse, error)
	GetPersonById(context.Context, *Id) (*PersonResponse, error)
}

// UnimplementedPersonServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) CreatePerson(context.Context, *Person) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonServiceServer) GetAllPeople(context.Context, *Empty) (*MultiPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPeople not implemented")
}
func (UnimplementedPersonServiceServer) GetPersonById(context.Context, *Id) (*PersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonById not implemented")
}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.PersonService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetAllPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetAllPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.PersonService/GetAllPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetAllPeople(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetPersonById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPersonById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.PersonService/GetPersonById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPersonById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stakeholder.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
		{
			MethodName: "GetAllPeople",
			Handler:    _PersonService_GetAllPeople_Handler,
		},
		{
			MethodName: "GetPersonById",
			Handler:    _PersonService_GetPersonById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/stakeholders/stakeholder_service.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserResponse, error)
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultiUserResponse, error)
	GetUserByUsername(ctx context.Context, in *Username, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultiUserResponse, error) {
	out := new(MultiUserResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *Username, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.UserService/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*EmptyResponse, error)
	GetUserById(context.Context, *Id) (*UserResponse, error)
	GetAllUsers(context.Context, *Empty) (*MultiUserResponse, error)
	GetUserByUsername(context.Context, *Username) (*UserResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserById(context.Context, *Id) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *Empty) (*MultiUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *Username) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.UserService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.UserService/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stakeholder.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/stakeholders/stakeholder_service.proto",
}

// AccessTokenServiceClient is the client API for AccessTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessTokenServiceClient interface {
	RegisterUser(ctx context.Context, in *Register, opts ...grpc.CallOption) (*EmptyResponse, error)
	LoginUser(ctx context.Context, in *Login, opts ...grpc.CallOption) (*AccessToken, error)
}

type accessTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessTokenServiceClient(cc grpc.ClientConnInterface) AccessTokenServiceClient {
	return &accessTokenServiceClient{cc}
}

func (c *accessTokenServiceClient) RegisterUser(ctx context.Context, in *Register, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/stakeholder.AccessTokenService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessTokenServiceClient) LoginUser(ctx context.Context, in *Login, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/stakeholder.AccessTokenService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessTokenServiceServer is the server API for AccessTokenService service.
// All implementations should embed UnimplementedAccessTokenServiceServer
// for forward compatibility
type AccessTokenServiceServer interface {
	RegisterUser(context.Context, *Register) (*EmptyResponse, error)
	LoginUser(context.Context, *Login) (*AccessToken, error)
}

// UnimplementedAccessTokenServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccessTokenServiceServer struct {
}

func (UnimplementedAccessTokenServiceServer) RegisterUser(context.Context, *Register) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAccessTokenServiceServer) LoginUser(context.Context, *Login) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}

// UnsafeAccessTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessTokenServiceServer will
// result in compilation errors.
type UnsafeAccessTokenServiceServer interface {
	mustEmbedUnimplementedAccessTokenServiceServer()
}

func RegisterAccessTokenServiceServer(s grpc.ServiceRegistrar, srv AccessTokenServiceServer) {
	s.RegisterService(&AccessTokenService_ServiceDesc, srv)
}

func _AccessTokenService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.AccessTokenService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).RegisterUser(ctx, req.(*Register))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessTokenService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakeholder.AccessTokenService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).LoginUser(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessTokenService_ServiceDesc is the grpc.ServiceDesc for AccessTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stakeholder.AccessTokenService",
	HandlerType: (*AccessTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _AccessTokenService_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AccessTokenService_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/stakeholders/stakeholder_service.proto",
}
