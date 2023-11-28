// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: barber/service_barber_shop.proto

package barber

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
	BarberShop_CreateBarber_FullMethodName          = "/pb.BarberShop/CreateBarber"
	BarberShop_LoginBarber_FullMethodName           = "/pb.BarberShop/LoginBarber"
	BarberShop_RefreshTokenBarber_FullMethodName    = "/pb.BarberShop/RefreshTokenBarber"
	BarberShop_UpdateBarber_FullMethodName          = "/pb.BarberShop/UpdateBarber"
	BarberShop_NewBarberShops_FullMethodName        = "/pb.BarberShop/NewBarberShops"
	BarberShop_FindBarberShopsNearby_FullMethodName = "/pb.BarberShop/FindBarberShopsNearby"
	BarberShop_NewServiceCategory_FullMethodName    = "/pb.BarberShop/NewServiceCategory"
	BarberShop_NewServices_FullMethodName           = "/pb.BarberShop/NewServices"
)

// BarberShopClient is the client API for BarberShop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BarberShopClient interface {
	CreateBarber(ctx context.Context, in *CreateBarberRequest, opts ...grpc.CallOption) (*CreateBarberResponse, error)
	LoginBarber(ctx context.Context, in *LoginBarberRequest, opts ...grpc.CallOption) (*LoginBarberResponse, error)
	RefreshTokenBarber(ctx context.Context, in *RefreshTokenBarberRequest, opts ...grpc.CallOption) (*RefreshTokenBarberResponse, error)
	UpdateBarber(ctx context.Context, in *UpdateBarberRequest, opts ...grpc.CallOption) (*UpdateBarberResponse, error)
	NewBarberShops(ctx context.Context, in *CreateBarberShopsRequest, opts ...grpc.CallOption) (*CreateBarberShopsResponse, error)
	FindBarberShopsNearby(ctx context.Context, in *FindBarberShopsNearbyRequest, opts ...grpc.CallOption) (*FindBarberShopsNearbyResponse, error)
	NewServiceCategory(ctx context.Context, in *CreateServiceCategoryRequest, opts ...grpc.CallOption) (*CreateServiceCategoryResponse, error)
	NewServices(ctx context.Context, in *CreateServicesRequest, opts ...grpc.CallOption) (*CreateServicesResponse, error)
}

type barberShopClient struct {
	cc grpc.ClientConnInterface
}

func NewBarberShopClient(cc grpc.ClientConnInterface) BarberShopClient {
	return &barberShopClient{cc}
}

func (c *barberShopClient) CreateBarber(ctx context.Context, in *CreateBarberRequest, opts ...grpc.CallOption) (*CreateBarberResponse, error) {
	out := new(CreateBarberResponse)
	err := c.cc.Invoke(ctx, BarberShop_CreateBarber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) LoginBarber(ctx context.Context, in *LoginBarberRequest, opts ...grpc.CallOption) (*LoginBarberResponse, error) {
	out := new(LoginBarberResponse)
	err := c.cc.Invoke(ctx, BarberShop_LoginBarber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) RefreshTokenBarber(ctx context.Context, in *RefreshTokenBarberRequest, opts ...grpc.CallOption) (*RefreshTokenBarberResponse, error) {
	out := new(RefreshTokenBarberResponse)
	err := c.cc.Invoke(ctx, BarberShop_RefreshTokenBarber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) UpdateBarber(ctx context.Context, in *UpdateBarberRequest, opts ...grpc.CallOption) (*UpdateBarberResponse, error) {
	out := new(UpdateBarberResponse)
	err := c.cc.Invoke(ctx, BarberShop_UpdateBarber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) NewBarberShops(ctx context.Context, in *CreateBarberShopsRequest, opts ...grpc.CallOption) (*CreateBarberShopsResponse, error) {
	out := new(CreateBarberShopsResponse)
	err := c.cc.Invoke(ctx, BarberShop_NewBarberShops_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) FindBarberShopsNearby(ctx context.Context, in *FindBarberShopsNearbyRequest, opts ...grpc.CallOption) (*FindBarberShopsNearbyResponse, error) {
	out := new(FindBarberShopsNearbyResponse)
	err := c.cc.Invoke(ctx, BarberShop_FindBarberShopsNearby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) NewServiceCategory(ctx context.Context, in *CreateServiceCategoryRequest, opts ...grpc.CallOption) (*CreateServiceCategoryResponse, error) {
	out := new(CreateServiceCategoryResponse)
	err := c.cc.Invoke(ctx, BarberShop_NewServiceCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barberShopClient) NewServices(ctx context.Context, in *CreateServicesRequest, opts ...grpc.CallOption) (*CreateServicesResponse, error) {
	out := new(CreateServicesResponse)
	err := c.cc.Invoke(ctx, BarberShop_NewServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BarberShopServer is the server API for BarberShop service.
// All implementations must embed UnimplementedBarberShopServer
// for forward compatibility
type BarberShopServer interface {
	CreateBarber(context.Context, *CreateBarberRequest) (*CreateBarberResponse, error)
	LoginBarber(context.Context, *LoginBarberRequest) (*LoginBarberResponse, error)
	RefreshTokenBarber(context.Context, *RefreshTokenBarberRequest) (*RefreshTokenBarberResponse, error)
	UpdateBarber(context.Context, *UpdateBarberRequest) (*UpdateBarberResponse, error)
	NewBarberShops(context.Context, *CreateBarberShopsRequest) (*CreateBarberShopsResponse, error)
	FindBarberShopsNearby(context.Context, *FindBarberShopsNearbyRequest) (*FindBarberShopsNearbyResponse, error)
	NewServiceCategory(context.Context, *CreateServiceCategoryRequest) (*CreateServiceCategoryResponse, error)
	NewServices(context.Context, *CreateServicesRequest) (*CreateServicesResponse, error)
	mustEmbedUnimplementedBarberShopServer()
}

// UnimplementedBarberShopServer must be embedded to have forward compatible implementations.
type UnimplementedBarberShopServer struct {
}

func (UnimplementedBarberShopServer) CreateBarber(context.Context, *CreateBarberRequest) (*CreateBarberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBarber not implemented")
}
func (UnimplementedBarberShopServer) LoginBarber(context.Context, *LoginBarberRequest) (*LoginBarberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginBarber not implemented")
}
func (UnimplementedBarberShopServer) RefreshTokenBarber(context.Context, *RefreshTokenBarberRequest) (*RefreshTokenBarberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokenBarber not implemented")
}
func (UnimplementedBarberShopServer) UpdateBarber(context.Context, *UpdateBarberRequest) (*UpdateBarberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBarber not implemented")
}
func (UnimplementedBarberShopServer) NewBarberShops(context.Context, *CreateBarberShopsRequest) (*CreateBarberShopsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBarberShops not implemented")
}
func (UnimplementedBarberShopServer) FindBarberShopsNearby(context.Context, *FindBarberShopsNearbyRequest) (*FindBarberShopsNearbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBarberShopsNearby not implemented")
}
func (UnimplementedBarberShopServer) NewServiceCategory(context.Context, *CreateServiceCategoryRequest) (*CreateServiceCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewServiceCategory not implemented")
}
func (UnimplementedBarberShopServer) NewServices(context.Context, *CreateServicesRequest) (*CreateServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewServices not implemented")
}
func (UnimplementedBarberShopServer) mustEmbedUnimplementedBarberShopServer() {}

// UnsafeBarberShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BarberShopServer will
// result in compilation errors.
type UnsafeBarberShopServer interface {
	mustEmbedUnimplementedBarberShopServer()
}

func RegisterBarberShopServer(s grpc.ServiceRegistrar, srv BarberShopServer) {
	s.RegisterService(&BarberShop_ServiceDesc, srv)
}

func _BarberShop_CreateBarber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBarberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).CreateBarber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_CreateBarber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).CreateBarber(ctx, req.(*CreateBarberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_LoginBarber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginBarberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).LoginBarber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_LoginBarber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).LoginBarber(ctx, req.(*LoginBarberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_RefreshTokenBarber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenBarberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).RefreshTokenBarber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_RefreshTokenBarber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).RefreshTokenBarber(ctx, req.(*RefreshTokenBarberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_UpdateBarber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBarberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).UpdateBarber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_UpdateBarber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).UpdateBarber(ctx, req.(*UpdateBarberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_NewBarberShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBarberShopsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).NewBarberShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_NewBarberShops_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).NewBarberShops(ctx, req.(*CreateBarberShopsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_FindBarberShopsNearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBarberShopsNearbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).FindBarberShopsNearby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_FindBarberShopsNearby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).FindBarberShopsNearby(ctx, req.(*FindBarberShopsNearbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_NewServiceCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).NewServiceCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_NewServiceCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).NewServiceCategory(ctx, req.(*CreateServiceCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarberShop_NewServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarberShopServer).NewServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarberShop_NewServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarberShopServer).NewServices(ctx, req.(*CreateServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BarberShop_ServiceDesc is the grpc.ServiceDesc for BarberShop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BarberShop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BarberShop",
	HandlerType: (*BarberShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBarber",
			Handler:    _BarberShop_CreateBarber_Handler,
		},
		{
			MethodName: "LoginBarber",
			Handler:    _BarberShop_LoginBarber_Handler,
		},
		{
			MethodName: "RefreshTokenBarber",
			Handler:    _BarberShop_RefreshTokenBarber_Handler,
		},
		{
			MethodName: "UpdateBarber",
			Handler:    _BarberShop_UpdateBarber_Handler,
		},
		{
			MethodName: "NewBarberShops",
			Handler:    _BarberShop_NewBarberShops_Handler,
		},
		{
			MethodName: "FindBarberShopsNearby",
			Handler:    _BarberShop_FindBarberShopsNearby_Handler,
		},
		{
			MethodName: "NewServiceCategory",
			Handler:    _BarberShop_NewServiceCategory_Handler,
		},
		{
			MethodName: "NewServices",
			Handler:    _BarberShop_NewServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "barber/service_barber_shop.proto",
}