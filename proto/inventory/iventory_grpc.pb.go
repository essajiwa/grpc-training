// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/inventory/iventory.proto

package inventory

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

// InventoryInfoClient is the client API for InventoryInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryInfoClient interface {
	GetStock(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Inventory, error)
	AddStock(ctx context.Context, in *Inventory, opts ...grpc.CallOption) (*Inventory, error)
}

type inventoryInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryInfoClient(cc grpc.ClientConnInterface) InventoryInfoClient {
	return &inventoryInfoClient{cc}
}

func (c *inventoryInfoClient) GetStock(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Inventory, error) {
	out := new(Inventory)
	err := c.cc.Invoke(ctx, "/inventory.InventoryInfo/GetStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryInfoClient) AddStock(ctx context.Context, in *Inventory, opts ...grpc.CallOption) (*Inventory, error) {
	out := new(Inventory)
	err := c.cc.Invoke(ctx, "/inventory.InventoryInfo/AddStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryInfoServer is the server API for InventoryInfo service.
// All implementations must embed UnimplementedInventoryInfoServer
// for forward compatibility
type InventoryInfoServer interface {
	GetStock(context.Context, *ProductID) (*Inventory, error)
	AddStock(context.Context, *Inventory) (*Inventory, error)
	mustEmbedUnimplementedInventoryInfoServer()
}

// UnimplementedInventoryInfoServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryInfoServer struct {
}

func (UnimplementedInventoryInfoServer) GetStock(context.Context, *ProductID) (*Inventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedInventoryInfoServer) AddStock(context.Context, *Inventory) (*Inventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStock not implemented")
}
func (UnimplementedInventoryInfoServer) mustEmbedUnimplementedInventoryInfoServer() {}

// UnsafeInventoryInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryInfoServer will
// result in compilation errors.
type UnsafeInventoryInfoServer interface {
	mustEmbedUnimplementedInventoryInfoServer()
}

func RegisterInventoryInfoServer(s grpc.ServiceRegistrar, srv InventoryInfoServer) {
	s.RegisterService(&InventoryInfo_ServiceDesc, srv)
}

func _InventoryInfo_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryInfoServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.InventoryInfo/GetStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryInfoServer).GetStock(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryInfo_AddStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Inventory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryInfoServer).AddStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.InventoryInfo/AddStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryInfoServer).AddStock(ctx, req.(*Inventory))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryInfo_ServiceDesc is the grpc.ServiceDesc for InventoryInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryInfo",
	HandlerType: (*InventoryInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStock",
			Handler:    _InventoryInfo_GetStock_Handler,
		},
		{
			MethodName: "AddStock",
			Handler:    _InventoryInfo_AddStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/inventory/iventory.proto",
}