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

// MeshClient is the client API for Mesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshClient interface {
	CheckInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// gateway列表
	GetGatewayList(ctx context.Context, in *ListOption, opts ...grpc.CallOption) (*GatewayList, error)
	// 创建 Gateway
	CreateGateway(ctx context.Context, in *Gateway, opts ...grpc.CallOption) (*Response, error)
	// 更新 Gateway
	UpdateGateway(ctx context.Context, in *Gateway, opts ...grpc.CallOption) (*Response, error)
	// 删除gateway
	DeleteGateway(ctx context.Context, in *DeleteKind, opts ...grpc.CallOption) (*Response, error)
}

type meshClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshClient(cc grpc.ClientConnInterface) MeshClient {
	return &meshClient{cc}
}

func (c *meshClient) CheckInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.Mesh/CheckInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) GetGatewayList(ctx context.Context, in *ListOption, opts ...grpc.CallOption) (*GatewayList, error) {
	out := new(GatewayList)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.Mesh/GetGatewayList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) CreateGateway(ctx context.Context, in *Gateway, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.Mesh/CreateGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) UpdateGateway(ctx context.Context, in *Gateway, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.Mesh/UpdateGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) DeleteGateway(ctx context.Context, in *DeleteKind, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.Mesh/DeleteGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshServer is the server API for Mesh service.
// All implementations must embed UnimplementedMeshServer
// for forward compatibility
type MeshServer interface {
	CheckInfo(context.Context, *Request) (*Response, error)
	// gateway列表
	GetGatewayList(context.Context, *ListOption) (*GatewayList, error)
	// 创建 Gateway
	CreateGateway(context.Context, *Gateway) (*Response, error)
	// 更新 Gateway
	UpdateGateway(context.Context, *Gateway) (*Response, error)
	// 删除gateway
	DeleteGateway(context.Context, *DeleteKind) (*Response, error)
	mustEmbedUnimplementedMeshServer()
}

// UnimplementedMeshServer must be embedded to have forward compatible implementations.
type UnimplementedMeshServer struct {
}

func (UnimplementedMeshServer) CheckInfo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInfo not implemented")
}
func (UnimplementedMeshServer) GetGatewayList(context.Context, *ListOption) (*GatewayList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayList not implemented")
}
func (UnimplementedMeshServer) CreateGateway(context.Context, *Gateway) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGateway not implemented")
}
func (UnimplementedMeshServer) UpdateGateway(context.Context, *Gateway) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGateway not implemented")
}
func (UnimplementedMeshServer) DeleteGateway(context.Context, *DeleteKind) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGateway not implemented")
}
func (UnimplementedMeshServer) mustEmbedUnimplementedMeshServer() {}

// UnsafeMeshServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshServer will
// result in compilation errors.
type UnsafeMeshServer interface {
	mustEmbedUnimplementedMeshServer()
}

func RegisterMeshServer(s grpc.ServiceRegistrar, srv MeshServer) {
	s.RegisterService(&Mesh_ServiceDesc, srv)
}

func _Mesh_CheckInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).CheckInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.Mesh/CheckInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).CheckInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_GetGatewayList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).GetGatewayList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.Mesh/GetGatewayList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).GetGatewayList(ctx, req.(*ListOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_CreateGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Gateway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).CreateGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.Mesh/CreateGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).CreateGateway(ctx, req.(*Gateway))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_UpdateGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Gateway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).UpdateGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.Mesh/UpdateGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).UpdateGateway(ctx, req.(*Gateway))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_DeleteGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).DeleteGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.Mesh/DeleteGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).DeleteGateway(ctx, req.(*DeleteKind))
	}
	return interceptor(ctx, in, info, handler)
}

// Mesh_ServiceDesc is the grpc.ServiceDesc for Mesh service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mesh_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cratos.api.v1.Mesh",
	HandlerType: (*MeshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckInfo",
			Handler:    _Mesh_CheckInfo_Handler,
		},
		{
			MethodName: "GetGatewayList",
			Handler:    _Mesh_GetGatewayList_Handler,
		},
		{
			MethodName: "CreateGateway",
			Handler:    _Mesh_CreateGateway_Handler,
		},
		{
			MethodName: "UpdateGateway",
			Handler:    _Mesh_UpdateGateway_Handler,
		},
		{
			MethodName: "DeleteGateway",
			Handler:    _Mesh_DeleteGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mesh.proto",
}
