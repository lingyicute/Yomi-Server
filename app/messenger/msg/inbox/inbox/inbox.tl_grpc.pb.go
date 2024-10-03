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
// source: inbox.tl.proto

package inbox

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
	RPCInbox_InboxEditUserMessageToInbox_FullMethodName     = "/inbox.RPCInbox/inbox_editUserMessageToInbox"
	RPCInbox_InboxEditChatMessageToInbox_FullMethodName     = "/inbox.RPCInbox/inbox_editChatMessageToInbox"
	RPCInbox_InboxDeleteMessagesToInbox_FullMethodName      = "/inbox.RPCInbox/inbox_deleteMessagesToInbox"
	RPCInbox_InboxDeleteUserHistoryToInbox_FullMethodName   = "/inbox.RPCInbox/inbox_deleteUserHistoryToInbox"
	RPCInbox_InboxDeleteChatHistoryToInbox_FullMethodName   = "/inbox.RPCInbox/inbox_deleteChatHistoryToInbox"
	RPCInbox_InboxReadUserMediaUnreadToInbox_FullMethodName = "/inbox.RPCInbox/inbox_readUserMediaUnreadToInbox"
	RPCInbox_InboxReadChatMediaUnreadToInbox_FullMethodName = "/inbox.RPCInbox/inbox_readChatMediaUnreadToInbox"
	RPCInbox_InboxUpdateHistoryReaded_FullMethodName        = "/inbox.RPCInbox/inbox_updateHistoryReaded"
	RPCInbox_InboxUpdatePinnedMessage_FullMethodName        = "/inbox.RPCInbox/inbox_updatePinnedMessage"
	RPCInbox_InboxUnpinAllMessages_FullMethodName           = "/inbox.RPCInbox/inbox_unpinAllMessages"
	RPCInbox_InboxSendUserMessageToInboxV2_FullMethodName   = "/inbox.RPCInbox/inbox_sendUserMessageToInboxV2"
	RPCInbox_InboxEditMessageToInboxV2_FullMethodName       = "/inbox.RPCInbox/inbox_editMessageToInboxV2"
	RPCInbox_InboxReadInboxHistory_FullMethodName           = "/inbox.RPCInbox/inbox_readInboxHistory"
	RPCInbox_InboxReadOutboxHistory_FullMethodName          = "/inbox.RPCInbox/inbox_readOutboxHistory"
	RPCInbox_InboxReadMediaUnreadToInboxV2_FullMethodName   = "/inbox.RPCInbox/inbox_readMediaUnreadToInboxV2"
)

// RPCInboxClient is the client API for RPCInbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCInboxClient interface {
	InboxEditUserMessageToInbox(ctx context.Context, in *TLInboxEditUserMessageToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxEditChatMessageToInbox(ctx context.Context, in *TLInboxEditChatMessageToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxDeleteMessagesToInbox(ctx context.Context, in *TLInboxDeleteMessagesToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxDeleteUserHistoryToInbox(ctx context.Context, in *TLInboxDeleteUserHistoryToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxDeleteChatHistoryToInbox(ctx context.Context, in *TLInboxDeleteChatHistoryToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxReadUserMediaUnreadToInbox(ctx context.Context, in *TLInboxReadUserMediaUnreadToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxReadChatMediaUnreadToInbox(ctx context.Context, in *TLInboxReadChatMediaUnreadToInbox, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxUpdateHistoryReaded(ctx context.Context, in *TLInboxUpdateHistoryReaded, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxUpdatePinnedMessage(ctx context.Context, in *TLInboxUpdatePinnedMessage, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxUnpinAllMessages(ctx context.Context, in *TLInboxUnpinAllMessages, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxSendUserMessageToInboxV2(ctx context.Context, in *TLInboxSendUserMessageToInboxV2, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxEditMessageToInboxV2(ctx context.Context, in *TLInboxEditMessageToInboxV2, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxReadInboxHistory(ctx context.Context, in *TLInboxReadInboxHistory, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxReadOutboxHistory(ctx context.Context, in *TLInboxReadOutboxHistory, opts ...grpc.CallOption) (*mtproto.Void, error)
	InboxReadMediaUnreadToInboxV2(ctx context.Context, in *TLInboxReadMediaUnreadToInboxV2, opts ...grpc.CallOption) (*mtproto.Void, error)
}

type rPCInboxClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCInboxClient(cc grpc.ClientConnInterface) RPCInboxClient {
	return &rPCInboxClient{cc}
}

func (c *rPCInboxClient) InboxEditUserMessageToInbox(ctx context.Context, in *TLInboxEditUserMessageToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxEditUserMessageToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxEditChatMessageToInbox(ctx context.Context, in *TLInboxEditChatMessageToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxEditChatMessageToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxDeleteMessagesToInbox(ctx context.Context, in *TLInboxDeleteMessagesToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxDeleteMessagesToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxDeleteUserHistoryToInbox(ctx context.Context, in *TLInboxDeleteUserHistoryToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxDeleteUserHistoryToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxDeleteChatHistoryToInbox(ctx context.Context, in *TLInboxDeleteChatHistoryToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxDeleteChatHistoryToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxReadUserMediaUnreadToInbox(ctx context.Context, in *TLInboxReadUserMediaUnreadToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxReadUserMediaUnreadToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxReadChatMediaUnreadToInbox(ctx context.Context, in *TLInboxReadChatMediaUnreadToInbox, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxReadChatMediaUnreadToInbox_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxUpdateHistoryReaded(ctx context.Context, in *TLInboxUpdateHistoryReaded, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxUpdateHistoryReaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxUpdatePinnedMessage(ctx context.Context, in *TLInboxUpdatePinnedMessage, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxUpdatePinnedMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxUnpinAllMessages(ctx context.Context, in *TLInboxUnpinAllMessages, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxUnpinAllMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxSendUserMessageToInboxV2(ctx context.Context, in *TLInboxSendUserMessageToInboxV2, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxSendUserMessageToInboxV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxEditMessageToInboxV2(ctx context.Context, in *TLInboxEditMessageToInboxV2, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxEditMessageToInboxV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxReadInboxHistory(ctx context.Context, in *TLInboxReadInboxHistory, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxReadInboxHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxReadOutboxHistory(ctx context.Context, in *TLInboxReadOutboxHistory, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxReadOutboxHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInboxClient) InboxReadMediaUnreadToInboxV2(ctx context.Context, in *TLInboxReadMediaUnreadToInboxV2, opts ...grpc.CallOption) (*mtproto.Void, error) {
	out := new(mtproto.Void)
	err := c.cc.Invoke(ctx, RPCInbox_InboxReadMediaUnreadToInboxV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCInboxServer is the server API for RPCInbox service.
// All implementations should embed UnimplementedRPCInboxServer
// for forward compatibility
type RPCInboxServer interface {
	InboxEditUserMessageToInbox(context.Context, *TLInboxEditUserMessageToInbox) (*mtproto.Void, error)
	InboxEditChatMessageToInbox(context.Context, *TLInboxEditChatMessageToInbox) (*mtproto.Void, error)
	InboxDeleteMessagesToInbox(context.Context, *TLInboxDeleteMessagesToInbox) (*mtproto.Void, error)
	InboxDeleteUserHistoryToInbox(context.Context, *TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error)
	InboxDeleteChatHistoryToInbox(context.Context, *TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error)
	InboxReadUserMediaUnreadToInbox(context.Context, *TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error)
	InboxReadChatMediaUnreadToInbox(context.Context, *TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error)
	InboxUpdateHistoryReaded(context.Context, *TLInboxUpdateHistoryReaded) (*mtproto.Void, error)
	InboxUpdatePinnedMessage(context.Context, *TLInboxUpdatePinnedMessage) (*mtproto.Void, error)
	InboxUnpinAllMessages(context.Context, *TLInboxUnpinAllMessages) (*mtproto.Void, error)
	InboxSendUserMessageToInboxV2(context.Context, *TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error)
	InboxEditMessageToInboxV2(context.Context, *TLInboxEditMessageToInboxV2) (*mtproto.Void, error)
	InboxReadInboxHistory(context.Context, *TLInboxReadInboxHistory) (*mtproto.Void, error)
	InboxReadOutboxHistory(context.Context, *TLInboxReadOutboxHistory) (*mtproto.Void, error)
	InboxReadMediaUnreadToInboxV2(context.Context, *TLInboxReadMediaUnreadToInboxV2) (*mtproto.Void, error)
}

// UnimplementedRPCInboxServer should be embedded to have forward compatible implementations.
type UnimplementedRPCInboxServer struct {
}

func (UnimplementedRPCInboxServer) InboxEditUserMessageToInbox(context.Context, *TLInboxEditUserMessageToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxEditUserMessageToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxEditChatMessageToInbox(context.Context, *TLInboxEditChatMessageToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxEditChatMessageToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxDeleteMessagesToInbox(context.Context, *TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxDeleteMessagesToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxDeleteUserHistoryToInbox(context.Context, *TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxDeleteUserHistoryToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxDeleteChatHistoryToInbox(context.Context, *TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxDeleteChatHistoryToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxReadUserMediaUnreadToInbox(context.Context, *TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxReadUserMediaUnreadToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxReadChatMediaUnreadToInbox(context.Context, *TLInboxReadChatMediaUnreadToInbox) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxReadChatMediaUnreadToInbox not implemented")
}
func (UnimplementedRPCInboxServer) InboxUpdateHistoryReaded(context.Context, *TLInboxUpdateHistoryReaded) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxUpdateHistoryReaded not implemented")
}
func (UnimplementedRPCInboxServer) InboxUpdatePinnedMessage(context.Context, *TLInboxUpdatePinnedMessage) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxUpdatePinnedMessage not implemented")
}
func (UnimplementedRPCInboxServer) InboxUnpinAllMessages(context.Context, *TLInboxUnpinAllMessages) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxUnpinAllMessages not implemented")
}
func (UnimplementedRPCInboxServer) InboxSendUserMessageToInboxV2(context.Context, *TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxSendUserMessageToInboxV2 not implemented")
}
func (UnimplementedRPCInboxServer) InboxEditMessageToInboxV2(context.Context, *TLInboxEditMessageToInboxV2) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxEditMessageToInboxV2 not implemented")
}
func (UnimplementedRPCInboxServer) InboxReadInboxHistory(context.Context, *TLInboxReadInboxHistory) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxReadInboxHistory not implemented")
}
func (UnimplementedRPCInboxServer) InboxReadOutboxHistory(context.Context, *TLInboxReadOutboxHistory) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxReadOutboxHistory not implemented")
}
func (UnimplementedRPCInboxServer) InboxReadMediaUnreadToInboxV2(context.Context, *TLInboxReadMediaUnreadToInboxV2) (*mtproto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboxReadMediaUnreadToInboxV2 not implemented")
}

// UnsafeRPCInboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCInboxServer will
// result in compilation errors.
type UnsafeRPCInboxServer interface {
	mustEmbedUnimplementedRPCInboxServer()
}

func RegisterRPCInboxServer(s grpc.ServiceRegistrar, srv RPCInboxServer) {
	s.RegisterService(&RPCInbox_ServiceDesc, srv)
}

func _RPCInbox_InboxEditUserMessageToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxEditUserMessageToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxEditUserMessageToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxEditUserMessageToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxEditUserMessageToInbox(ctx, req.(*TLInboxEditUserMessageToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxEditChatMessageToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxEditChatMessageToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxEditChatMessageToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxEditChatMessageToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxEditChatMessageToInbox(ctx, req.(*TLInboxEditChatMessageToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxDeleteMessagesToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxDeleteMessagesToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxDeleteMessagesToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxDeleteMessagesToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxDeleteMessagesToInbox(ctx, req.(*TLInboxDeleteMessagesToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxDeleteUserHistoryToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxDeleteUserHistoryToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxDeleteUserHistoryToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxDeleteUserHistoryToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxDeleteUserHistoryToInbox(ctx, req.(*TLInboxDeleteUserHistoryToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxDeleteChatHistoryToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxDeleteChatHistoryToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxDeleteChatHistoryToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxDeleteChatHistoryToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxDeleteChatHistoryToInbox(ctx, req.(*TLInboxDeleteChatHistoryToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxReadUserMediaUnreadToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxReadUserMediaUnreadToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxReadUserMediaUnreadToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxReadUserMediaUnreadToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxReadUserMediaUnreadToInbox(ctx, req.(*TLInboxReadUserMediaUnreadToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxReadChatMediaUnreadToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxReadChatMediaUnreadToInbox)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxReadChatMediaUnreadToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxReadChatMediaUnreadToInbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxReadChatMediaUnreadToInbox(ctx, req.(*TLInboxReadChatMediaUnreadToInbox))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxUpdateHistoryReaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxUpdateHistoryReaded)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxUpdateHistoryReaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxUpdateHistoryReaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxUpdateHistoryReaded(ctx, req.(*TLInboxUpdateHistoryReaded))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxUpdatePinnedMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxUpdatePinnedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxUpdatePinnedMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxUpdatePinnedMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxUpdatePinnedMessage(ctx, req.(*TLInboxUpdatePinnedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxUnpinAllMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxUnpinAllMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxUnpinAllMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxUnpinAllMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxUnpinAllMessages(ctx, req.(*TLInboxUnpinAllMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxSendUserMessageToInboxV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxSendUserMessageToInboxV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxSendUserMessageToInboxV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxSendUserMessageToInboxV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxSendUserMessageToInboxV2(ctx, req.(*TLInboxSendUserMessageToInboxV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxEditMessageToInboxV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxEditMessageToInboxV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxEditMessageToInboxV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxEditMessageToInboxV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxEditMessageToInboxV2(ctx, req.(*TLInboxEditMessageToInboxV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxReadInboxHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxReadInboxHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxReadInboxHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxReadInboxHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxReadInboxHistory(ctx, req.(*TLInboxReadInboxHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxReadOutboxHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxReadOutboxHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxReadOutboxHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxReadOutboxHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxReadOutboxHistory(ctx, req.(*TLInboxReadOutboxHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInbox_InboxReadMediaUnreadToInboxV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLInboxReadMediaUnreadToInboxV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInboxServer).InboxReadMediaUnreadToInboxV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCInbox_InboxReadMediaUnreadToInboxV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInboxServer).InboxReadMediaUnreadToInboxV2(ctx, req.(*TLInboxReadMediaUnreadToInboxV2))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCInbox_ServiceDesc is the grpc.ServiceDesc for RPCInbox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCInbox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inbox.RPCInbox",
	HandlerType: (*RPCInboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "inbox_editUserMessageToInbox",
			Handler:    _RPCInbox_InboxEditUserMessageToInbox_Handler,
		},
		{
			MethodName: "inbox_editChatMessageToInbox",
			Handler:    _RPCInbox_InboxEditChatMessageToInbox_Handler,
		},
		{
			MethodName: "inbox_deleteMessagesToInbox",
			Handler:    _RPCInbox_InboxDeleteMessagesToInbox_Handler,
		},
		{
			MethodName: "inbox_deleteUserHistoryToInbox",
			Handler:    _RPCInbox_InboxDeleteUserHistoryToInbox_Handler,
		},
		{
			MethodName: "inbox_deleteChatHistoryToInbox",
			Handler:    _RPCInbox_InboxDeleteChatHistoryToInbox_Handler,
		},
		{
			MethodName: "inbox_readUserMediaUnreadToInbox",
			Handler:    _RPCInbox_InboxReadUserMediaUnreadToInbox_Handler,
		},
		{
			MethodName: "inbox_readChatMediaUnreadToInbox",
			Handler:    _RPCInbox_InboxReadChatMediaUnreadToInbox_Handler,
		},
		{
			MethodName: "inbox_updateHistoryReaded",
			Handler:    _RPCInbox_InboxUpdateHistoryReaded_Handler,
		},
		{
			MethodName: "inbox_updatePinnedMessage",
			Handler:    _RPCInbox_InboxUpdatePinnedMessage_Handler,
		},
		{
			MethodName: "inbox_unpinAllMessages",
			Handler:    _RPCInbox_InboxUnpinAllMessages_Handler,
		},
		{
			MethodName: "inbox_sendUserMessageToInboxV2",
			Handler:    _RPCInbox_InboxSendUserMessageToInboxV2_Handler,
		},
		{
			MethodName: "inbox_editMessageToInboxV2",
			Handler:    _RPCInbox_InboxEditMessageToInboxV2_Handler,
		},
		{
			MethodName: "inbox_readInboxHistory",
			Handler:    _RPCInbox_InboxReadInboxHistory_Handler,
		},
		{
			MethodName: "inbox_readOutboxHistory",
			Handler:    _RPCInbox_InboxReadOutboxHistory_Handler,
		},
		{
			MethodName: "inbox_readMediaUnreadToInboxV2",
			Handler:    _RPCInbox_InboxReadMediaUnreadToInboxV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inbox.tl.proto",
}
