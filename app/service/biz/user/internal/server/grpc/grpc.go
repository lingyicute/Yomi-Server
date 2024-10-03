/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package grpc

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/svc"
	userpb "github.com/lingyicute/papercraft-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		userpb.RegisterRPCUserServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
