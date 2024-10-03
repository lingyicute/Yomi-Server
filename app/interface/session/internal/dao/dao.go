/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package dao

import (
	"context"

	"github.com/papercraft/marmota/pkg/cache"
	"github.com/papercraft/marmota/pkg/net/ip"
	"github.com/papercraft/proto/mtproto"
	"github.com/papercraft/proto/mtproto/rpc/metadata"
	bff_proxy_client "github.com/lingyicute/papercraft-server/app/bff/bff/client"
	"github.com/lingyicute/papercraft-server/app/interface/session/internal/config"
	authsession_client "github.com/lingyicute/papercraft-server/app/service/authsession/client"
	status_client "github.com/lingyicute/papercraft-server/app/service/status/client"

	"github.com/zeromicro/go-zero/zrpc"
)

type Dao struct {
	cache *cache.LRUCache
	authsession_client.AuthsessionClient
	status_client.StatusClient
	*bff_proxy_client.BFFProxyClient
	eGateServers map[string]*Gateway
	MyServerId   string
	*RpcShardingManager
}

func New(c config.Config) *Dao {
	myServerId := ip.FigureOutListenOn(c.ListenOn)
	d := &Dao{
		cache:              cache.NewLRUCache(1024 * 1024 * 1024),
		AuthsessionClient:  authsession_client.NewAuthsessionClient(zrpc.MustNewClient(c.AuthSession)),
		BFFProxyClient:     bff_proxy_client.NewBFFProxyClients(c.BFFProxyClients.Clients, c.BFFProxyClients.IDMap),
		StatusClient:       status_client.NewStatusClient(zrpc.MustNewClient(c.StatusClient)),
		eGateServers:       make(map[string]*Gateway),
		MyServerId:         myServerId,
		RpcShardingManager: NewRpcShardingManager(myServerId, c.Etcd),
	}

	d.watchGateway(c.GatewayClient)

	return d
}

func (d *Dao) InvokeContext(ctx context.Context, rpcMetaData *metadata.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	return d.BFFProxyClient.InvokeContext(ctx, rpcMetaData, object)
}
