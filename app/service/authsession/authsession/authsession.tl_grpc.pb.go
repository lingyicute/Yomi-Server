//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Papercraft Authors.
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: authsession.tl.proto

package authsession

import (
	context "context"
	mtproto "github.com/teamgram/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCAuthsession_AuthsessionGetAuthorizations_FullMethodName    = "/authsession.RPCAuthsession/authsession_getAuthorizations"
	RPCAuthsession_AuthsessionResetAuthorization_FullMethodName   = "/authsession.RPCAuthsession/authsession_resetAuthorization"
	RPCAuthsession_AuthsessionGetLayer_FullMethodName             = "/authsession.RPCAuthsession/authsession_getLayer"
	RPCAuthsession_AuthsessionGetLangPack_FullMethodName          = "/authsession.RPCAuthsession/authsession_getLangPack"
	RPCAuthsession_AuthsessionGetClient_FullMethodName            = "/authsession.RPCAuthsession/authsession_getClient"
	RPCAuthsession_AuthsessionGetLangCode_FullMethodName          = "/authsession.RPCAuthsession/authsession_getLangCode"
	RPCAuthsession_AuthsessionGetUserId_FullMethodName            = "/authsession.RPCAuthsession/authsession_getUserId"
	RPCAuthsession_AuthsessionGetPushSessionId_FullMethodName     = "/authsession.RPCAuthsession/authsession_getPushSessionId"
	RPCAuthsession_AuthsessionGetFutureSalts_FullMethodName       = "/authsession.RPCAuthsession/authsession_getFutureSalts"
	RPCAuthsession_AuthsessionQueryAuthKey_FullMethodName         = "/authsession.RPCAuthsession/authsession_queryAuthKey"
	RPCAuthsession_AuthsessionSetAuthKey_FullMethodName           = "/authsession.RPCAuthsession/authsession_setAuthKey"
	RPCAuthsession_AuthsessionBindAuthKeyUser_FullMethodName      = "/authsession.RPCAuthsession/authsession_bindAuthKeyUser"
	RPCAuthsession_AuthsessionUnbindAuthKeyUser_FullMethodName    = "/authsession.RPCAuthsession/authsession_unbindAuthKeyUser"
	RPCAuthsession_AuthsessionGetPermAuthKeyId_FullMethodName     = "/authsession.RPCAuthsession/authsession_getPermAuthKeyId"
	RPCAuthsession_AuthsessionBindTempAuthKey_FullMethodName      = "/authsession.RPCAuthsession/authsession_bindTempAuthKey"
	RPCAuthsession_AuthsessionSetClientSessionInfo_FullMethodName = "/authsession.RPCAuthsession/authsession_setClientSessionInfo"
	RPCAuthsession_AuthsessionGetAuthorization_FullMethodName     = "/authsession.RPCAuthsession/authsession_getAuthorization"
	RPCAuthsession_AuthsessionGetAuthStateData_FullMethodName     = "/authsession.RPCAuthsession/authsession_getAuthStateData"
	RPCAuthsession_AuthsessionSetLayer_FullMethodName             = "/authsession.RPCAuthsession/authsession_setLayer"
	RPCAuthsession_AuthsessionSetInitConnection_FullMethodName    = "/authsession.RPCAuthsession/authsession_setInitConnection"
)

// RPCAuthsessionClient is the client API for RPCAuthsession service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCAuthsessionClient interface {
	AuthsessionGetAuthorizations(ctx context.Context, in *TLAuthsessionGetAuthorizations, opts ...grpc.CallOption) (*mtproto.Account_Authorizations, error)
	AuthsessionResetAuthorization(ctx context.Context, in *TLAuthsessionResetAuthorization, opts ...grpc.CallOption) (*Vector_Long, error)
	AuthsessionGetLayer(ctx context.Context, in *TLAuthsessionGetLayer, opts ...grpc.CallOption) (*mtproto.Int32, error)
	AuthsessionGetLangPack(ctx context.Context, in *TLAuthsessionGetLangPack, opts ...grpc.CallOption) (*mtproto.String, error)
	AuthsessionGetClient(ctx context.Context, in *TLAuthsessionGetClient, opts ...grpc.CallOption) (*mtproto.String, error)
	AuthsessionGetLangCode(ctx context.Context, in *TLAuthsessionGetLangCode, opts ...grpc.CallOption) (*mtproto.String, error)
	AuthsessionGetUserId(ctx context.Context, in *TLAuthsessionGetUserId, opts ...grpc.CallOption) (*mtproto.Int64, error)
	AuthsessionGetPushSessionId(ctx context.Context, in *TLAuthsessionGetPushSessionId, opts ...grpc.CallOption) (*mtproto.Int64, error)
	AuthsessionGetFutureSalts(ctx context.Context, in *TLAuthsessionGetFutureSalts, opts ...grpc.CallOption) (*mtproto.FutureSalts, error)
	AuthsessionQueryAuthKey(ctx context.Context, in *TLAuthsessionQueryAuthKey, opts ...grpc.CallOption) (*mtproto.AuthKeyInfo, error)
	AuthsessionSetAuthKey(ctx context.Context, in *TLAuthsessionSetAuthKey, opts ...grpc.CallOption) (*mtproto.Bool, error)
	AuthsessionBindAuthKeyUser(ctx context.Context, in *TLAuthsessionBindAuthKeyUser, opts ...grpc.CallOption) (*mtproto.Int64, error)
	AuthsessionUnbindAuthKeyUser(ctx context.Context, in *TLAuthsessionUnbindAuthKeyUser, opts ...grpc.CallOption) (*mtproto.Bool, error)
	AuthsessionGetPermAuthKeyId(ctx context.Context, in *TLAuthsessionGetPermAuthKeyId, opts ...grpc.CallOption) (*mtproto.Int64, error)
	AuthsessionBindTempAuthKey(ctx context.Context, in *TLAuthsessionBindTempAuthKey, opts ...grpc.CallOption) (*mtproto.Bool, error)
	AuthsessionSetClientSessionInfo(ctx context.Context, in *TLAuthsessionSetClientSessionInfo, opts ...grpc.CallOption) (*mtproto.Bool, error)
	AuthsessionGetAuthorization(ctx context.Context, in *TLAuthsessionGetAuthorization, opts ...grpc.CallOption) (*mtproto.Authorization, error)
	AuthsessionGetAuthStateData(ctx context.Context, in *TLAuthsessionGetAuthStateData, opts ...grpc.CallOption) (*AuthKeyStateData, error)
	AuthsessionSetLayer(ctx context.Context, in *TLAuthsessionSetLayer, opts ...grpc.CallOption) (*mtproto.Bool, error)
	AuthsessionSetInitConnection(ctx context.Context, in *TLAuthsessionSetInitConnection, opts ...grpc.CallOption) (*mtproto.Bool, error)
}

type rPCAuthsessionClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCAuthsessionClient(cc grpc.ClientConnInterface) RPCAuthsessionClient {
	return &rPCAuthsessionClient{cc}
}

func (c *rPCAuthsessionClient) AuthsessionGetAuthorizations(ctx context.Context, in *TLAuthsessionGetAuthorizations, opts ...grpc.CallOption) (*mtproto.Account_Authorizations, error) {
	out := new(mtproto.Account_Authorizations)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetAuthorizations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionResetAuthorization(ctx context.Context, in *TLAuthsessionResetAuthorization, opts ...grpc.CallOption) (*Vector_Long, error) {
	out := new(Vector_Long)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionResetAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetLayer(ctx context.Context, in *TLAuthsessionGetLayer, opts ...grpc.CallOption) (*mtproto.Int32, error) {
	out := new(mtproto.Int32)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetLayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetLangPack(ctx context.Context, in *TLAuthsessionGetLangPack, opts ...grpc.CallOption) (*mtproto.String, error) {
	out := new(mtproto.String)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetLangPack_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetClient(ctx context.Context, in *TLAuthsessionGetClient, opts ...grpc.CallOption) (*mtproto.String, error) {
	out := new(mtproto.String)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetLangCode(ctx context.Context, in *TLAuthsessionGetLangCode, opts ...grpc.CallOption) (*mtproto.String, error) {
	out := new(mtproto.String)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetLangCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetUserId(ctx context.Context, in *TLAuthsessionGetUserId, opts ...grpc.CallOption) (*mtproto.Int64, error) {
	out := new(mtproto.Int64)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetPushSessionId(ctx context.Context, in *TLAuthsessionGetPushSessionId, opts ...grpc.CallOption) (*mtproto.Int64, error) {
	out := new(mtproto.Int64)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetPushSessionId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetFutureSalts(ctx context.Context, in *TLAuthsessionGetFutureSalts, opts ...grpc.CallOption) (*mtproto.FutureSalts, error) {
	out := new(mtproto.FutureSalts)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetFutureSalts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionQueryAuthKey(ctx context.Context, in *TLAuthsessionQueryAuthKey, opts ...grpc.CallOption) (*mtproto.AuthKeyInfo, error) {
	out := new(mtproto.AuthKeyInfo)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionQueryAuthKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionSetAuthKey(ctx context.Context, in *TLAuthsessionSetAuthKey, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionSetAuthKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionBindAuthKeyUser(ctx context.Context, in *TLAuthsessionBindAuthKeyUser, opts ...grpc.CallOption) (*mtproto.Int64, error) {
	out := new(mtproto.Int64)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionBindAuthKeyUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionUnbindAuthKeyUser(ctx context.Context, in *TLAuthsessionUnbindAuthKeyUser, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionUnbindAuthKeyUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetPermAuthKeyId(ctx context.Context, in *TLAuthsessionGetPermAuthKeyId, opts ...grpc.CallOption) (*mtproto.Int64, error) {
	out := new(mtproto.Int64)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetPermAuthKeyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionBindTempAuthKey(ctx context.Context, in *TLAuthsessionBindTempAuthKey, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionBindTempAuthKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionSetClientSessionInfo(ctx context.Context, in *TLAuthsessionSetClientSessionInfo, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionSetClientSessionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetAuthorization(ctx context.Context, in *TLAuthsessionGetAuthorization, opts ...grpc.CallOption) (*mtproto.Authorization, error) {
	out := new(mtproto.Authorization)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionGetAuthStateData(ctx context.Context, in *TLAuthsessionGetAuthStateData, opts ...grpc.CallOption) (*AuthKeyStateData, error) {
	out := new(AuthKeyStateData)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionGetAuthStateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionSetLayer(ctx context.Context, in *TLAuthsessionSetLayer, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionSetLayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthsessionClient) AuthsessionSetInitConnection(ctx context.Context, in *TLAuthsessionSetInitConnection, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCAuthsession_AuthsessionSetInitConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCAuthsessionServer is the server API for RPCAuthsession service.
// All implementations should embed UnimplementedRPCAuthsessionServer
// for forward compatibility
type RPCAuthsessionServer interface {
	AuthsessionGetAuthorizations(context.Context, *TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error)
	AuthsessionResetAuthorization(context.Context, *TLAuthsessionResetAuthorization) (*Vector_Long, error)
	AuthsessionGetLayer(context.Context, *TLAuthsessionGetLayer) (*mtproto.Int32, error)
	AuthsessionGetLangPack(context.Context, *TLAuthsessionGetLangPack) (*mtproto.String, error)
	AuthsessionGetClient(context.Context, *TLAuthsessionGetClient) (*mtproto.String, error)
	AuthsessionGetLangCode(context.Context, *TLAuthsessionGetLangCode) (*mtproto.String, error)
	AuthsessionGetUserId(context.Context, *TLAuthsessionGetUserId) (*mtproto.Int64, error)
	AuthsessionGetPushSessionId(context.Context, *TLAuthsessionGetPushSessionId) (*mtproto.Int64, error)
	AuthsessionGetFutureSalts(context.Context, *TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error)
	AuthsessionQueryAuthKey(context.Context, *TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error)
	AuthsessionSetAuthKey(context.Context, *TLAuthsessionSetAuthKey) (*mtproto.Bool, error)
	AuthsessionBindAuthKeyUser(context.Context, *TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error)
	AuthsessionUnbindAuthKeyUser(context.Context, *TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error)
	AuthsessionGetPermAuthKeyId(context.Context, *TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error)
	AuthsessionBindTempAuthKey(context.Context, *TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error)
	AuthsessionSetClientSessionInfo(context.Context, *TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error)
	AuthsessionGetAuthorization(context.Context, *TLAuthsessionGetAuthorization) (*mtproto.Authorization, error)
	AuthsessionGetAuthStateData(context.Context, *TLAuthsessionGetAuthStateData) (*AuthKeyStateData, error)
	AuthsessionSetLayer(context.Context, *TLAuthsessionSetLayer) (*mtproto.Bool, error)
	AuthsessionSetInitConnection(context.Context, *TLAuthsessionSetInitConnection) (*mtproto.Bool, error)
}

// UnimplementedRPCAuthsessionServer should be embedded to have forward compatible implementations.
type UnimplementedRPCAuthsessionServer struct {
}

func (UnimplementedRPCAuthsessionServer) AuthsessionGetAuthorizations(context.Context, *TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetAuthorizations not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionResetAuthorization(context.Context, *TLAuthsessionResetAuthorization) (*Vector_Long, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionResetAuthorization not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetLayer(context.Context, *TLAuthsessionGetLayer) (*mtproto.Int32, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetLayer not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetLangPack(context.Context, *TLAuthsessionGetLangPack) (*mtproto.String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetLangPack not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetClient(context.Context, *TLAuthsessionGetClient) (*mtproto.String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetClient not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetLangCode(context.Context, *TLAuthsessionGetLangCode) (*mtproto.String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetLangCode not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetUserId(context.Context, *TLAuthsessionGetUserId) (*mtproto.Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetUserId not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetPushSessionId(context.Context, *TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetPushSessionId not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetFutureSalts(context.Context, *TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetFutureSalts not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionQueryAuthKey(context.Context, *TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionQueryAuthKey not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionSetAuthKey(context.Context, *TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionSetAuthKey not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionBindAuthKeyUser(context.Context, *TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionBindAuthKeyUser not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionUnbindAuthKeyUser(context.Context, *TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionUnbindAuthKeyUser not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetPermAuthKeyId(context.Context, *TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetPermAuthKeyId not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionBindTempAuthKey(context.Context, *TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionBindTempAuthKey not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionSetClientSessionInfo(context.Context, *TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionSetClientSessionInfo not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetAuthorization(context.Context, *TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetAuthorization not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionGetAuthStateData(context.Context, *TLAuthsessionGetAuthStateData) (*AuthKeyStateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionGetAuthStateData not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionSetLayer(context.Context, *TLAuthsessionSetLayer) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionSetLayer not implemented")
}
func (UnimplementedRPCAuthsessionServer) AuthsessionSetInitConnection(context.Context, *TLAuthsessionSetInitConnection) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthsessionSetInitConnection not implemented")
}

// UnsafeRPCAuthsessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCAuthsessionServer will
// result in compilation errors.
type UnsafeRPCAuthsessionServer interface {
	mustEmbedUnimplementedRPCAuthsessionServer()
}

func RegisterRPCAuthsessionServer(s grpc.ServiceRegistrar, srv RPCAuthsessionServer) {
	s.RegisterService(&RPCAuthsession_ServiceDesc, srv)
}

func _RPCAuthsession_AuthsessionGetAuthorizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetAuthorizations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetAuthorizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetAuthorizations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetAuthorizations(ctx, req.(*TLAuthsessionGetAuthorizations))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionResetAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionResetAuthorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionResetAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionResetAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionResetAuthorization(ctx, req.(*TLAuthsessionResetAuthorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetLayer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetLayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetLayer(ctx, req.(*TLAuthsessionGetLayer))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetLangPack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetLangPack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetLangPack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetLangPack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetLangPack(ctx, req.(*TLAuthsessionGetLangPack))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetClient(ctx, req.(*TLAuthsessionGetClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetLangCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetLangCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetLangCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetLangCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetLangCode(ctx, req.(*TLAuthsessionGetLangCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetUserId(ctx, req.(*TLAuthsessionGetUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetPushSessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetPushSessionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetPushSessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetPushSessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetPushSessionId(ctx, req.(*TLAuthsessionGetPushSessionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetFutureSalts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetFutureSalts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetFutureSalts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetFutureSalts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetFutureSalts(ctx, req.(*TLAuthsessionGetFutureSalts))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionQueryAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionQueryAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionQueryAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionQueryAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionQueryAuthKey(ctx, req.(*TLAuthsessionQueryAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionSetAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionSetAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionSetAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionSetAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionSetAuthKey(ctx, req.(*TLAuthsessionSetAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionBindAuthKeyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionBindAuthKeyUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionBindAuthKeyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionBindAuthKeyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionBindAuthKeyUser(ctx, req.(*TLAuthsessionBindAuthKeyUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionUnbindAuthKeyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionUnbindAuthKeyUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionUnbindAuthKeyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionUnbindAuthKeyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionUnbindAuthKeyUser(ctx, req.(*TLAuthsessionUnbindAuthKeyUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetPermAuthKeyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetPermAuthKeyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetPermAuthKeyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetPermAuthKeyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetPermAuthKeyId(ctx, req.(*TLAuthsessionGetPermAuthKeyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionBindTempAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionBindTempAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionBindTempAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionBindTempAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionBindTempAuthKey(ctx, req.(*TLAuthsessionBindTempAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionSetClientSessionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionSetClientSessionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionSetClientSessionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionSetClientSessionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionSetClientSessionInfo(ctx, req.(*TLAuthsessionSetClientSessionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetAuthorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetAuthorization(ctx, req.(*TLAuthsessionGetAuthorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionGetAuthStateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionGetAuthStateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionGetAuthStateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionGetAuthStateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionGetAuthStateData(ctx, req.(*TLAuthsessionGetAuthStateData))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionSetLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionSetLayer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionSetLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionSetLayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionSetLayer(ctx, req.(*TLAuthsessionSetLayer))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthsession_AuthsessionSetInitConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthsessionSetInitConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthsessionServer).AuthsessionSetInitConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCAuthsession_AuthsessionSetInitConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthsessionServer).AuthsessionSetInitConnection(ctx, req.(*TLAuthsessionSetInitConnection))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCAuthsession_ServiceDesc is the grpc.ServiceDesc for RPCAuthsession service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCAuthsession_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authsession.RPCAuthsession",
	HandlerType: (*RPCAuthsessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "authsession_getAuthorizations",
			Handler:    _RPCAuthsession_AuthsessionGetAuthorizations_Handler,
		},
		{
			MethodName: "authsession_resetAuthorization",
			Handler:    _RPCAuthsession_AuthsessionResetAuthorization_Handler,
		},
		{
			MethodName: "authsession_getLayer",
			Handler:    _RPCAuthsession_AuthsessionGetLayer_Handler,
		},
		{
			MethodName: "authsession_getLangPack",
			Handler:    _RPCAuthsession_AuthsessionGetLangPack_Handler,
		},
		{
			MethodName: "authsession_getClient",
			Handler:    _RPCAuthsession_AuthsessionGetClient_Handler,
		},
		{
			MethodName: "authsession_getLangCode",
			Handler:    _RPCAuthsession_AuthsessionGetLangCode_Handler,
		},
		{
			MethodName: "authsession_getUserId",
			Handler:    _RPCAuthsession_AuthsessionGetUserId_Handler,
		},
		{
			MethodName: "authsession_getPushSessionId",
			Handler:    _RPCAuthsession_AuthsessionGetPushSessionId_Handler,
		},
		{
			MethodName: "authsession_getFutureSalts",
			Handler:    _RPCAuthsession_AuthsessionGetFutureSalts_Handler,
		},
		{
			MethodName: "authsession_queryAuthKey",
			Handler:    _RPCAuthsession_AuthsessionQueryAuthKey_Handler,
		},
		{
			MethodName: "authsession_setAuthKey",
			Handler:    _RPCAuthsession_AuthsessionSetAuthKey_Handler,
		},
		{
			MethodName: "authsession_bindAuthKeyUser",
			Handler:    _RPCAuthsession_AuthsessionBindAuthKeyUser_Handler,
		},
		{
			MethodName: "authsession_unbindAuthKeyUser",
			Handler:    _RPCAuthsession_AuthsessionUnbindAuthKeyUser_Handler,
		},
		{
			MethodName: "authsession_getPermAuthKeyId",
			Handler:    _RPCAuthsession_AuthsessionGetPermAuthKeyId_Handler,
		},
		{
			MethodName: "authsession_bindTempAuthKey",
			Handler:    _RPCAuthsession_AuthsessionBindTempAuthKey_Handler,
		},
		{
			MethodName: "authsession_setClientSessionInfo",
			Handler:    _RPCAuthsession_AuthsessionSetClientSessionInfo_Handler,
		},
		{
			MethodName: "authsession_getAuthorization",
			Handler:    _RPCAuthsession_AuthsessionGetAuthorization_Handler,
		},
		{
			MethodName: "authsession_getAuthStateData",
			Handler:    _RPCAuthsession_AuthsessionGetAuthStateData_Handler,
		},
		{
			MethodName: "authsession_setLayer",
			Handler:    _RPCAuthsession_AuthsessionSetLayer_Handler,
		},
		{
			MethodName: "authsession_setInitConnection",
			Handler:    _RPCAuthsession_AuthsessionSetInitConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authsession.tl.proto",
}
