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
	"github.com/lingyicute/papercraft-server/app/interface/gateway/gateway"
	"github.com/lingyicute/papercraft-server/app/interface/gnetway/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/interface/gnetway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(svcCtx *svc.ServiceContext, c zrpc.RpcServerConf, srv gateway.RPCGatewayServer) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		gateway.RegisterRPCGatewayServer(grpcServer, service.New(svcCtx, srv))
	})
	logx.Must(err)
	return s
}
