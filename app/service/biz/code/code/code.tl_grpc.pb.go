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
// source: code.tl.proto

package code

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
	RPCCode_CodeCreatePhoneCode_FullMethodName     = "/code.RPCCode/code_createPhoneCode"
	RPCCode_CodeGetPhoneCode_FullMethodName        = "/code.RPCCode/code_getPhoneCode"
	RPCCode_CodeDeletePhoneCode_FullMethodName     = "/code.RPCCode/code_deletePhoneCode"
	RPCCode_CodeUpdatePhoneCodeData_FullMethodName = "/code.RPCCode/code_updatePhoneCodeData"
)

// RPCCodeClient is the client API for RPCCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCCodeClient interface {
	CodeCreatePhoneCode(ctx context.Context, in *TLCodeCreatePhoneCode, opts ...grpc.CallOption) (*PhoneCodeTransaction, error)
	CodeGetPhoneCode(ctx context.Context, in *TLCodeGetPhoneCode, opts ...grpc.CallOption) (*PhoneCodeTransaction, error)
	CodeDeletePhoneCode(ctx context.Context, in *TLCodeDeletePhoneCode, opts ...grpc.CallOption) (*mtproto.Bool, error)
	CodeUpdatePhoneCodeData(ctx context.Context, in *TLCodeUpdatePhoneCodeData, opts ...grpc.CallOption) (*mtproto.Bool, error)
}

type rPCCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCCodeClient(cc grpc.ClientConnInterface) RPCCodeClient {
	return &rPCCodeClient{cc}
}

func (c *rPCCodeClient) CodeCreatePhoneCode(ctx context.Context, in *TLCodeCreatePhoneCode, opts ...grpc.CallOption) (*PhoneCodeTransaction, error) {
	out := new(PhoneCodeTransaction)
	err := c.cc.Invoke(ctx, RPCCode_CodeCreatePhoneCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCodeClient) CodeGetPhoneCode(ctx context.Context, in *TLCodeGetPhoneCode, opts ...grpc.CallOption) (*PhoneCodeTransaction, error) {
	out := new(PhoneCodeTransaction)
	err := c.cc.Invoke(ctx, RPCCode_CodeGetPhoneCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCodeClient) CodeDeletePhoneCode(ctx context.Context, in *TLCodeDeletePhoneCode, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCCode_CodeDeletePhoneCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCodeClient) CodeUpdatePhoneCodeData(ctx context.Context, in *TLCodeUpdatePhoneCodeData, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCCode_CodeUpdatePhoneCodeData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCCodeServer is the server API for RPCCode service.
// All implementations should embed UnimplementedRPCCodeServer
// for forward compatibility
type RPCCodeServer interface {
	CodeCreatePhoneCode(context.Context, *TLCodeCreatePhoneCode) (*PhoneCodeTransaction, error)
	CodeGetPhoneCode(context.Context, *TLCodeGetPhoneCode) (*PhoneCodeTransaction, error)
	CodeDeletePhoneCode(context.Context, *TLCodeDeletePhoneCode) (*mtproto.Bool, error)
	CodeUpdatePhoneCodeData(context.Context, *TLCodeUpdatePhoneCodeData) (*mtproto.Bool, error)
}

// UnimplementedRPCCodeServer should be embedded to have forward compatible implementations.
type UnimplementedRPCCodeServer struct {
}

func (UnimplementedRPCCodeServer) CodeCreatePhoneCode(context.Context, *TLCodeCreatePhoneCode) (*PhoneCodeTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeCreatePhoneCode not implemented")
}
func (UnimplementedRPCCodeServer) CodeGetPhoneCode(context.Context, *TLCodeGetPhoneCode) (*PhoneCodeTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeGetPhoneCode not implemented")
}
func (UnimplementedRPCCodeServer) CodeDeletePhoneCode(context.Context, *TLCodeDeletePhoneCode) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeDeletePhoneCode not implemented")
}
func (UnimplementedRPCCodeServer) CodeUpdatePhoneCodeData(context.Context, *TLCodeUpdatePhoneCodeData) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeUpdatePhoneCodeData not implemented")
}

// UnsafeRPCCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCCodeServer will
// result in compilation errors.
type UnsafeRPCCodeServer interface {
	mustEmbedUnimplementedRPCCodeServer()
}

func RegisterRPCCodeServer(s grpc.ServiceRegistrar, srv RPCCodeServer) {
	s.RegisterService(&RPCCode_ServiceDesc, srv)
}

func _RPCCode_CodeCreatePhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLCodeCreatePhoneCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCodeServer).CodeCreatePhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCCode_CodeCreatePhoneCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCodeServer).CodeCreatePhoneCode(ctx, req.(*TLCodeCreatePhoneCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCode_CodeGetPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLCodeGetPhoneCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCodeServer).CodeGetPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCCode_CodeGetPhoneCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCodeServer).CodeGetPhoneCode(ctx, req.(*TLCodeGetPhoneCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCode_CodeDeletePhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLCodeDeletePhoneCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCodeServer).CodeDeletePhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCCode_CodeDeletePhoneCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCodeServer).CodeDeletePhoneCode(ctx, req.(*TLCodeDeletePhoneCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCode_CodeUpdatePhoneCodeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLCodeUpdatePhoneCodeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCodeServer).CodeUpdatePhoneCodeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCCode_CodeUpdatePhoneCodeData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCodeServer).CodeUpdatePhoneCodeData(ctx, req.(*TLCodeUpdatePhoneCodeData))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCCode_ServiceDesc is the grpc.ServiceDesc for RPCCode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCCode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.RPCCode",
	HandlerType: (*RPCCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "code_createPhoneCode",
			Handler:    _RPCCode_CodeCreatePhoneCode_Handler,
		},
		{
			MethodName: "code_getPhoneCode",
			Handler:    _RPCCode_CodeGetPhoneCode_Handler,
		},
		{
			MethodName: "code_deletePhoneCode",
			Handler:    _RPCCode_CodeDeletePhoneCode_Handler,
		},
		{
			MethodName: "code_updatePhoneCodeData",
			Handler:    _RPCCode_CodeUpdatePhoneCodeData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code.tl.proto",
}
