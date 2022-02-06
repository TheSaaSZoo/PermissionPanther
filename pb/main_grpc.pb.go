// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// PermissionPantherClient is the client API for PermissionPanther service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionPantherClient interface {
	// Checks whether a permission exists, and at what recursion level. If there is an explicit `deny` permission then group checking will be aborted.
	CheckDirectPermission(ctx context.Context, in *CheckDirectReq, opts ...grpc.CallOption) (*CheckDirectRes, error)
	// Lists all the permissions an entity has, optionally specify permissions to filter on
	ListEntityRelations(ctx context.Context, in *ListEntityRelationsReq, opts ...grpc.CallOption) (*RelationsResponse, error)
	// List all relations for an object, optoinally specify permissions to filter on
	ListObjectRelations(ctx context.Context, in *ListObjectRelationsReq, opts ...grpc.CallOption) (*RelationsResponse, error)
	// Will set a permission for an entity on an object. If the permission already exists it is a no-op.
	SetPermission(ctx context.Context, in *RelationReq, opts ...grpc.CallOption) (*Applied, error)
	// Will remove a permission for an entity on an object. If the permission does not exist it is a no-op.
	RemovePermission(ctx context.Context, in *RelationReq, opts ...grpc.CallOption) (*Applied, error)
}

type permissionPantherClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionPantherClient(cc grpc.ClientConnInterface) PermissionPantherClient {
	return &permissionPantherClient{cc}
}

func (c *permissionPantherClient) CheckDirectPermission(ctx context.Context, in *CheckDirectReq, opts ...grpc.CallOption) (*CheckDirectRes, error) {
	out := new(CheckDirectRes)
	err := c.cc.Invoke(ctx, "/PermissionPanther/CheckDirectPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPantherClient) ListEntityRelations(ctx context.Context, in *ListEntityRelationsReq, opts ...grpc.CallOption) (*RelationsResponse, error) {
	out := new(RelationsResponse)
	err := c.cc.Invoke(ctx, "/PermissionPanther/ListEntityRelations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPantherClient) ListObjectRelations(ctx context.Context, in *ListObjectRelationsReq, opts ...grpc.CallOption) (*RelationsResponse, error) {
	out := new(RelationsResponse)
	err := c.cc.Invoke(ctx, "/PermissionPanther/ListObjectRelations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPantherClient) SetPermission(ctx context.Context, in *RelationReq, opts ...grpc.CallOption) (*Applied, error) {
	out := new(Applied)
	err := c.cc.Invoke(ctx, "/PermissionPanther/SetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPantherClient) RemovePermission(ctx context.Context, in *RelationReq, opts ...grpc.CallOption) (*Applied, error) {
	out := new(Applied)
	err := c.cc.Invoke(ctx, "/PermissionPanther/RemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionPantherServer is the server API for PermissionPanther service.
// All implementations must embed UnimplementedPermissionPantherServer
// for forward compatibility
type PermissionPantherServer interface {
	// Checks whether a permission exists, and at what recursion level. If there is an explicit `deny` permission then group checking will be aborted.
	CheckDirectPermission(context.Context, *CheckDirectReq) (*CheckDirectRes, error)
	// Lists all the permissions an entity has, optionally specify permissions to filter on
	ListEntityRelations(context.Context, *ListEntityRelationsReq) (*RelationsResponse, error)
	// List all relations for an object, optoinally specify permissions to filter on
	ListObjectRelations(context.Context, *ListObjectRelationsReq) (*RelationsResponse, error)
	// Will set a permission for an entity on an object. If the permission already exists it is a no-op.
	SetPermission(context.Context, *RelationReq) (*Applied, error)
	// Will remove a permission for an entity on an object. If the permission does not exist it is a no-op.
	RemovePermission(context.Context, *RelationReq) (*Applied, error)
	mustEmbedUnimplementedPermissionPantherServer()
}

// UnimplementedPermissionPantherServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionPantherServer struct {
}

func (UnimplementedPermissionPantherServer) CheckDirectPermission(context.Context, *CheckDirectReq) (*CheckDirectRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDirectPermission not implemented")
}
func (UnimplementedPermissionPantherServer) ListEntityRelations(context.Context, *ListEntityRelationsReq) (*RelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntityRelations not implemented")
}
func (UnimplementedPermissionPantherServer) ListObjectRelations(context.Context, *ListObjectRelationsReq) (*RelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjectRelations not implemented")
}
func (UnimplementedPermissionPantherServer) SetPermission(context.Context, *RelationReq) (*Applied, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPermission not implemented")
}
func (UnimplementedPermissionPantherServer) RemovePermission(context.Context, *RelationReq) (*Applied, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermission not implemented")
}
func (UnimplementedPermissionPantherServer) mustEmbedUnimplementedPermissionPantherServer() {}

// UnsafePermissionPantherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionPantherServer will
// result in compilation errors.
type UnsafePermissionPantherServer interface {
	mustEmbedUnimplementedPermissionPantherServer()
}

func RegisterPermissionPantherServer(s grpc.ServiceRegistrar, srv PermissionPantherServer) {
	s.RegisterService(&PermissionPanther_ServiceDesc, srv)
}

func _PermissionPanther_CheckDirectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDirectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPantherServer).CheckDirectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PermissionPanther/CheckDirectPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPantherServer).CheckDirectPermission(ctx, req.(*CheckDirectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPanther_ListEntityRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntityRelationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPantherServer).ListEntityRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PermissionPanther/ListEntityRelations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPantherServer).ListEntityRelations(ctx, req.(*ListEntityRelationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPanther_ListObjectRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectRelationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPantherServer).ListObjectRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PermissionPanther/ListObjectRelations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPantherServer).ListObjectRelations(ctx, req.(*ListObjectRelationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPanther_SetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPantherServer).SetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PermissionPanther/SetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPantherServer).SetPermission(ctx, req.(*RelationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPanther_RemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPantherServer).RemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PermissionPanther/RemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPantherServer).RemovePermission(ctx, req.(*RelationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionPanther_ServiceDesc is the grpc.ServiceDesc for PermissionPanther service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionPanther_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PermissionPanther",
	HandlerType: (*PermissionPantherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckDirectPermission",
			Handler:    _PermissionPanther_CheckDirectPermission_Handler,
		},
		{
			MethodName: "ListEntityRelations",
			Handler:    _PermissionPanther_ListEntityRelations_Handler,
		},
		{
			MethodName: "ListObjectRelations",
			Handler:    _PermissionPanther_ListObjectRelations_Handler,
		},
		{
			MethodName: "SetPermission",
			Handler:    _PermissionPanther_SetPermission_Handler,
		},
		{
			MethodName: "RemovePermission",
			Handler:    _PermissionPanther_RemovePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/main.proto",
}
