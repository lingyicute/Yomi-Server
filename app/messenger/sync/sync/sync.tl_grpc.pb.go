//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Papercraft Authors.
//  All rights reserved.
//
// Author: @lingyicute

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: sync.tl.proto

package sync

import (
	context "context"
	mtproto "github.com/papercraft/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCSync_SyncUpdatesMe_FullMethodName        = "/sync.RPCSync/sync_updatesMe"
	RPCSync_SyncUpdatesNotMe_FullMethodName     = "/sync.RPCSync/sync_updatesNotMe"
	RPCSync_SyncPushUpdates_FullMethodName      = "/sync.RPCSync/sync_pushUpdates"
	RPCSync_SyncPushUpdatesIfNot_FullMethodName = "/sync.RPCSync/sync_pushUpdatesIfNot"
	RPCSync_SyncPushBotUpdates_FullMethodName   = "/sync.RPCSync/sync_pushBotUpdates"
	RPCSync_SyncPushRpcResult_FullMethodName    = "/sync.RPCSync/sync_pushRpcResult"
	RPCSync_SyncBroadcastUpdates_FullMethodName = "/sync.RPCSync/sync_broadcastUpdates"
)

// RPCSyncClient is the client API for RPCSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCSyncClient interface {
	SyncUpdatesMe(ctx context.Context, in *TLSyncUpdatesMe, opts ...grpc.CallOption) (*mtproto.Void, error)
	SyncUpdatesNotMe(ctx context.Context, in *TLSyncUpdatesNotMe, opts ...grpc.CallOption) (*mtproto.Void, error)
	SyncPushUpdates(ctx context.Context, in *TLSyncPushUpdates, opts ...grpc.CallOption) (*mtproto.Void, error)
	SyncPushUpdatesIfNot(ctx context.Context, in *TLSyncPushUpdatesIfNot, opts ...grpc.CallOption) (*mtproto.Void, error)
	SyncPushBotUpdates(ctx context.Context, in *TLSyncPushBotUpdates, opts ...grpc.CallOption) (*mtproto.Void, error)
	SyncPushRpcResult(ctx context.Context, in *TLSyncPushRpcResult, opts ...grpc.CallOption) (*mtproto.Void, error)
	SyncBroadcastUpdates(ctx context.Context, in *TLSyncBroadcastUpdates, opts ...grpc.CallOption) (*mtproto.Void, error)
}

type rPCSyncClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCSyncClient(cc grpc.ClientConnInterface) RPCSyncClient {
	return &rPCSyncClient{cc}
}

func (c *rPCSyncClient) SyncUpdatesMe(ctx context.Context, in *TLSyncUpdatesMe, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncUpdatesMe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncUpdatesNotMe(ctx context.Context, in *TLSyncUpdatesNotMe, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncUpdatesNotMe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncPushUpdates(ctx context.Context, in *TLSyncPushUpdates, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncPushUpdates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncPushUpdatesIfNot(ctx context.Context, in *TLSyncPushUpdatesIfNot, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncPushUpdatesIfNot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncPushBotUpdates(ctx context.Context, in *TLSyncPushBotUpdates, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncPushBotUpdates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncPushRpcResult(ctx context.Context, in *TLSyncPushRpcResult, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncPushRpcResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncBroadcastUpdates(ctx context.Context, in *TLSyncBroadcastUpdates, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCSync_SyncBroadcastUpdates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCSyncServer is the server API for RPCSync service.
// All implementations should embed UnimplementedRPCSyncServer
// for forward compatibility
type RPCSyncServer interface {
	SyncUpdatesMe(context.Context, *TLSyncUpdatesMe) (*mtproto.Void, error)
	SyncUpdatesNotMe(context.Context, *TLSyncUpdatesNotMe) (*mtproto.Void, error)
	SyncPushUpdates(context.Context, *TLSyncPushUpdates) (*mtproto.Void, error)
	SyncPushUpdatesIfNot(context.Context, *TLSyncPushUpdatesIfNot) (*mtproto.Void, error)
	SyncPushBotUpdates(context.Context, *TLSyncPushBotUpdates) (*mtproto.Void, error)
	SyncPushRpcResult(context.Context, *TLSyncPushRpcResult) (*mtproto.Void, error)
	SyncBroadcastUpdates(context.Context, *TLSyncBroadcastUpdates) (*mtproto.Void, error)
}

// UnimplementedRPCSyncServer should be embedded to have forward compatible implementations.
type UnimplementedRPCSyncServer struct {
}

func (UnimplementedRPCSyncServer) SyncUpdatesMe(context.Context, *TLSyncUpdatesMe) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncUpdatesMe not implemented")
}
func (UnimplementedRPCSyncServer) SyncUpdatesNotMe(context.Context, *TLSyncUpdatesNotMe) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncUpdatesNotMe not implemented")
}
func (UnimplementedRPCSyncServer) SyncPushUpdates(context.Context, *TLSyncPushUpdates) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPushUpdates not implemented")
}
func (UnimplementedRPCSyncServer) SyncPushUpdatesIfNot(context.Context, *TLSyncPushUpdatesIfNot) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPushUpdatesIfNot not implemented")
}
func (UnimplementedRPCSyncServer) SyncPushBotUpdates(context.Context, *TLSyncPushBotUpdates) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPushBotUpdates not implemented")
}
func (UnimplementedRPCSyncServer) SyncPushRpcResult(context.Context, *TLSyncPushRpcResult) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPushRpcResult not implemented")
}
func (UnimplementedRPCSyncServer) SyncBroadcastUpdates(context.Context, *TLSyncBroadcastUpdates) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncBroadcastUpdates not implemented")
}

// UnsafeRPCSyncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCSyncServer will
// result in compilation errors.
type UnsafeRPCSyncServer interface {
	mustEmbedUnimplementedRPCSyncServer()
}

func RegisterRPCSyncServer(s grpc.ServiceRegistrar, srv RPCSyncServer) {
	s.RegisterService(&RPCSync_ServiceDesc, srv)
}

func _RPCSync_SyncUpdatesMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncUpdatesMe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdatesMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncUpdatesMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdatesMe(ctx, req.(*TLSyncUpdatesMe))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncUpdatesNotMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncUpdatesNotMe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdatesNotMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncUpdatesNotMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdatesNotMe(ctx, req.(*TLSyncUpdatesNotMe))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncPushUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncPushUpdates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncPushUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncPushUpdates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncPushUpdates(ctx, req.(*TLSyncPushUpdates))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncPushUpdatesIfNot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncPushUpdatesIfNot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncPushUpdatesIfNot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncPushUpdatesIfNot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncPushUpdatesIfNot(ctx, req.(*TLSyncPushUpdatesIfNot))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncPushBotUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncPushBotUpdates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncPushBotUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncPushBotUpdates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncPushBotUpdates(ctx, req.(*TLSyncPushBotUpdates))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncPushRpcResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncPushRpcResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncPushRpcResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncPushRpcResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncPushRpcResult(ctx, req.(*TLSyncPushRpcResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncBroadcastUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSyncBroadcastUpdates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncBroadcastUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSync_SyncBroadcastUpdates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncBroadcastUpdates(ctx, req.(*TLSyncBroadcastUpdates))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCSync_ServiceDesc is the grpc.ServiceDesc for RPCSync service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCSync_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sync.RPCSync",
	HandlerType: (*RPCSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sync_updatesMe",
			Handler:    _RPCSync_SyncUpdatesMe_Handler,
		},
		{
			MethodName: "sync_updatesNotMe",
			Handler:    _RPCSync_SyncUpdatesNotMe_Handler,
		},
		{
			MethodName: "sync_pushUpdates",
			Handler:    _RPCSync_SyncPushUpdates_Handler,
		},
		{
			MethodName: "sync_pushUpdatesIfNot",
			Handler:    _RPCSync_SyncPushUpdatesIfNot_Handler,
		},
		{
			MethodName: "sync_pushBotUpdates",
			Handler:    _RPCSync_SyncPushBotUpdates_Handler,
		},
		{
			MethodName: "sync_pushRpcResult",
			Handler:    _RPCSync_SyncPushRpcResult_Handler,
		},
		{
			MethodName: "sync_broadcastUpdates",
			Handler:    _RPCSync_SyncBroadcastUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync.tl.proto",
}
