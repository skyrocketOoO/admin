// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: main.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Main_Login_FullMethodName           = "/proto.Main/Login"
	Main_Logout_FullMethodName          = "/proto.Main/Logout"
	Main_CreateAccount_FullMethodName   = "/proto.Main/CreateAccount"
	Main_ListAccount_FullMethodName     = "/proto.Main/ListAccount"
	Main_UpdateAccount_FullMethodName   = "/proto.Main/UpdateAccount"
	Main_DeleteAccount_FullMethodName   = "/proto.Main/DeleteAccount"
	Main_DeactiveAccount_FullMethodName = "/proto.Main/DeactiveAccount"
	Main_ActiveAccount_FullMethodName   = "/proto.Main/ActiveAccount"
	Main_CreateRole_FullMethodName      = "/proto.Main/CreateRole"
	Main_ListRole_FullMethodName        = "/proto.Main/ListRole"
	Main_GetRoleAuth_FullMethodName     = "/proto.Main/GetRoleAuth"
	Main_UpdateRole_FullMethodName      = "/proto.Main/UpdateRole"
	Main_DeleteRole_FullMethodName      = "/proto.Main/DeleteRole"
	Main_BindRole_FullMethodName        = "/proto.Main/BindRole"
	Main_UnBindRole_FullMethodName      = "/proto.Main/UnBindRole"
)

// MainClient is the client API for Main service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Empty, error)
	CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*Empty, error)
	ListAccount(ctx context.Context, in *ListAccountReq, opts ...grpc.CallOption) (*ListAccountResp, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*Empty, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*Empty, error)
	DeactiveAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*Empty, error)
	ActiveAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*Empty, error)
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*Empty, error)
	ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error)
	GetRoleAuth(ctx context.Context, in *GetRoleAuthReq, opts ...grpc.CallOption) (*GetRoleAuthResp, error)
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*Empty, error)
	DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*Empty, error)
	BindRole(ctx context.Context, in *BindRoleReq, opts ...grpc.CallOption) (*Empty, error)
	UnBindRole(ctx context.Context, in *UnBindRoleReq, opts ...grpc.CallOption) (*Empty, error)
}

type mainClient struct {
	cc grpc.ClientConnInterface
}

func NewMainClient(cc grpc.ClientConnInterface) MainClient {
	return &mainClient{cc}
}

func (c *mainClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Main_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) ListAccount(ctx context.Context, in *ListAccountReq, opts ...grpc.CallOption) (*ListAccountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAccountResp)
	err := c.cc.Invoke(ctx, Main_ListAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_UpdateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) DeleteAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_DeleteAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) DeactiveAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_DeactiveAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) ActiveAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_ActiveAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRoleResp)
	err := c.cc.Invoke(ctx, Main_ListRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) GetRoleAuth(ctx context.Context, in *GetRoleAuthReq, opts ...grpc.CallOption) (*GetRoleAuthResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoleAuthResp)
	err := c.cc.Invoke(ctx, Main_GetRoleAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) BindRole(ctx context.Context, in *BindRoleReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_BindRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) UnBindRole(ctx context.Context, in *UnBindRoleReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Main_UnBindRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServer is the server API for Main service.
// All implementations must embed UnimplementedMainServer
// for forward compatibility.
type MainServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*Empty, error)
	CreateAccount(context.Context, *CreateAccountReq) (*Empty, error)
	ListAccount(context.Context, *ListAccountReq) (*ListAccountResp, error)
	UpdateAccount(context.Context, *UpdateAccountReq) (*Empty, error)
	DeleteAccount(context.Context, *DeleteAccountReq) (*Empty, error)
	DeactiveAccount(context.Context, *DeleteAccountReq) (*Empty, error)
	ActiveAccount(context.Context, *DeleteAccountReq) (*Empty, error)
	CreateRole(context.Context, *CreateRoleReq) (*Empty, error)
	ListRole(context.Context, *ListRoleReq) (*ListRoleResp, error)
	GetRoleAuth(context.Context, *GetRoleAuthReq) (*GetRoleAuthResp, error)
	UpdateRole(context.Context, *UpdateRoleReq) (*Empty, error)
	DeleteRole(context.Context, *DeleteRoleReq) (*Empty, error)
	BindRole(context.Context, *BindRoleReq) (*Empty, error)
	UnBindRole(context.Context, *UnBindRoleReq) (*Empty, error)
	mustEmbedUnimplementedMainServer()
}

// UnimplementedMainServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMainServer struct{}

func (UnimplementedMainServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMainServer) Logout(context.Context, *LogoutReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedMainServer) CreateAccount(context.Context, *CreateAccountReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedMainServer) ListAccount(context.Context, *ListAccountReq) (*ListAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}
func (UnimplementedMainServer) UpdateAccount(context.Context, *UpdateAccountReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedMainServer) DeleteAccount(context.Context, *DeleteAccountReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedMainServer) DeactiveAccount(context.Context, *DeleteAccountReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactiveAccount not implemented")
}
func (UnimplementedMainServer) ActiveAccount(context.Context, *DeleteAccountReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveAccount not implemented")
}
func (UnimplementedMainServer) CreateRole(context.Context, *CreateRoleReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedMainServer) ListRole(context.Context, *ListRoleReq) (*ListRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedMainServer) GetRoleAuth(context.Context, *GetRoleAuthReq) (*GetRoleAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleAuth not implemented")
}
func (UnimplementedMainServer) UpdateRole(context.Context, *UpdateRoleReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedMainServer) DeleteRole(context.Context, *DeleteRoleReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedMainServer) BindRole(context.Context, *BindRoleReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindRole not implemented")
}
func (UnimplementedMainServer) UnBindRole(context.Context, *UnBindRoleReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBindRole not implemented")
}
func (UnimplementedMainServer) mustEmbedUnimplementedMainServer() {}
func (UnimplementedMainServer) testEmbeddedByValue()              {}

// UnsafeMainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServer will
// result in compilation errors.
type UnsafeMainServer interface {
	mustEmbedUnimplementedMainServer()
}

func RegisterMainServer(s grpc.ServiceRegistrar, srv MainServer) {
	// If the following call pancis, it indicates UnimplementedMainServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Main_ServiceDesc, srv)
}

func _Main_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CreateAccount(ctx, req.(*CreateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_ListAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).ListAccount(ctx, req.(*ListAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).UpdateAccount(ctx, req.(*UpdateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeleteAccount(ctx, req.(*DeleteAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_DeactiveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeactiveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_DeactiveAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeactiveAccount(ctx, req.(*DeleteAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_ActiveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).ActiveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_ActiveAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).ActiveAccount(ctx, req.(*DeleteAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).ListRole(ctx, req.(*ListRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_GetRoleAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).GetRoleAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_GetRoleAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).GetRoleAuth(ctx, req.(*GetRoleAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeleteRole(ctx, req.(*DeleteRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_BindRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).BindRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_BindRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).BindRole(ctx, req.(*BindRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_UnBindRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnBindRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).UnBindRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_UnBindRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).UnBindRole(ctx, req.(*UnBindRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Main_ServiceDesc is the grpc.ServiceDesc for Main service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Main_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Main",
	HandlerType: (*MainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Main_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Main_Logout_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Main_CreateAccount_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _Main_ListAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Main_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Main_DeleteAccount_Handler,
		},
		{
			MethodName: "DeactiveAccount",
			Handler:    _Main_DeactiveAccount_Handler,
		},
		{
			MethodName: "ActiveAccount",
			Handler:    _Main_ActiveAccount_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _Main_CreateRole_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _Main_ListRole_Handler,
		},
		{
			MethodName: "GetRoleAuth",
			Handler:    _Main_GetRoleAuth_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Main_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Main_DeleteRole_Handler,
		},
		{
			MethodName: "BindRole",
			Handler:    _Main_BindRole_Handler,
		},
		{
			MethodName: "UnBindRole",
			Handler:    _Main_UnBindRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
