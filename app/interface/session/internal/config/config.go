/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package config

import (
	"github.com/lingyicute/papercraft-server/pkg/conf"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AuthSession     zrpc.RpcClientConf
	StatusClient    zrpc.RpcClientConf
	GatewayClient   zrpc.RpcClientConf
	BFFProxyClients conf.BFFProxyClients
}

// Routine routine.
type Routine struct {
	Size uint64
	Chan uint64
}
