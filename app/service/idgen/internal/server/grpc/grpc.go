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
	"github.com/lingyicute/papercraft-server/app/service/idgen/idgen"
	"github.com/lingyicute/papercraft-server/app/service/idgen/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/idgen/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		idgen.RegisterRPCIdgenServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
