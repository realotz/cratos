// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package namespace

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NamespacesClient is the client API for Namespaces service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespacesClient interface {
	// Namespace列表
	List(ctx context.Context, in *ListOption, opts ...grpc.CallOption) (*NamespaceList, error)
	// 获取Namespace
	Get(ctx context.Context, in *GetKind, opts ...grpc.CallOption) (*v1.Namespace, error)
	// 创建 Namespace
	Create(ctx context.Context, in *v1.Namespace, opts ...grpc.CallOption) (*Response, error)
	// 更新 Namespace
	Update(ctx context.Context, in *v1.Namespace, opts ...grpc.CallOption) (*Response, error)
	// 删除Namespace
	Delete(ctx context.Context, in *DeleteKind, opts ...grpc.CallOption) (*Response, error)
}

type namespacesClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespacesClient(cc grpc.ClientConnInterface) NamespacesClient {
	return &namespacesClient{cc}
}

func (c *namespacesClient) List(ctx context.Context, in *ListOption, opts ...grpc.CallOption) (*NamespaceList, error) {
	out := new(NamespaceList)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.namespace.Namespaces/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Get(ctx context.Context, in *GetKind, opts ...grpc.CallOption) (*v1.Namespace, error) {
	out := new(v1.Namespace)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.namespace.Namespaces/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Create(ctx context.Context, in *v1.Namespace, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.namespace.Namespaces/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Update(ctx context.Context, in *v1.Namespace, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.namespace.Namespaces/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) Delete(ctx context.Context, in *DeleteKind, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cratos.api.v1.namespace.Namespaces/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespacesServer is the server API for Namespaces service.
// All implementations must embed UnimplementedNamespacesServer
// for forward compatibility
type NamespacesServer interface {
	// Namespace列表
	List(context.Context, *ListOption) (*NamespaceList, error)
	// 获取Namespace
	Get(context.Context, *GetKind) (*v1.Namespace, error)
	// 创建 Namespace
	Create(context.Context, *v1.Namespace) (*Response, error)
	// 更新 Namespace
	Update(context.Context, *v1.Namespace) (*Response, error)
	// 删除Namespace
	Delete(context.Context, *DeleteKind) (*Response, error)
	mustEmbedUnimplementedNamespacesServer()
}

// UnimplementedNamespacesServer must be embedded to have forward compatible implementations.
type UnimplementedNamespacesServer struct {
}

func (UnimplementedNamespacesServer) List(context.Context, *ListOption) (*NamespaceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNamespacesServer) Get(context.Context, *GetKind) (*v1.Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNamespacesServer) Create(context.Context, *v1.Namespace) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNamespacesServer) Update(context.Context, *v1.Namespace) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNamespacesServer) Delete(context.Context, *DeleteKind) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNamespacesServer) mustEmbedUnimplementedNamespacesServer() {}

// UnsafeNamespacesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespacesServer will
// result in compilation errors.
type UnsafeNamespacesServer interface {
	mustEmbedUnimplementedNamespacesServer()
}

func RegisterNamespacesServer(s grpc.ServiceRegistrar, srv NamespacesServer) {
	s.RegisterService(&Namespaces_ServiceDesc, srv)
}

func _Namespaces_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.namespace.Namespaces/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).List(ctx, req.(*ListOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.namespace.Namespaces/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Get(ctx, req.(*GetKind))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.namespace.Namespaces/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Create(ctx, req.(*v1.Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.namespace.Namespaces/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Update(ctx, req.(*v1.Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cratos.api.v1.namespace.Namespaces/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Delete(ctx, req.(*DeleteKind))
	}
	return interceptor(ctx, in, info, handler)
}

// Namespaces_ServiceDesc is the grpc.ServiceDesc for Namespaces service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Namespaces_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cratos.api.v1.namespace.Namespaces",
	HandlerType: (*NamespacesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Namespaces_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Namespaces_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Namespaces_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Namespaces_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Namespaces_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namespace.proto",
}