/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package gateway_client

import (
	"context"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/interface/gateway/gateway"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type GatewayClient interface {
	GatewaySendDataToGateway(ctx context.Context, in *gateway.TLGatewaySendDataToGateway) (*mtproto.Bool, error)
}

type defaultGatewayClient struct {
	cli zrpc.Client
}

func NewGatewayClient(cli zrpc.Client) GatewayClient {
	return &defaultGatewayClient{
		cli: cli,
	}
}

// GatewaySendDataToGateway
// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (m *defaultGatewayClient) GatewaySendDataToGateway(ctx context.Context, in *gateway.TLGatewaySendDataToGateway) (*mtproto.Bool, error) {
	client := gateway.NewRPCGatewayClient(m.cli.Conn())
	return client.GatewaySendDataToGateway(ctx, in)
}
