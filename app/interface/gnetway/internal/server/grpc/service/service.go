/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package service

import (
	"github.com/lingyicute/papercraft-server/app/interface/gateway/gateway"
	"github.com/lingyicute/papercraft-server/app/interface/gnetway/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
	gateway.RPCGatewayServer
}

func New(ctx *svc.ServiceContext, srv gateway.RPCGatewayServer) *Service {
	return &Service{
		svcCtx:           ctx,
		RPCGatewayServer: srv,
	}
}
