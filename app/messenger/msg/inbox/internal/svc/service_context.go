/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package svc

import (
	kafka "github.com/papercraft/marmota/pkg/mq"
	"github.com/papercraft/marmota/pkg/net/rpcx"
	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/inbox/internal/config"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/internal/dao"
	sync_client "github.com/lingyicute/papercraft-server/app/messenger/sync/client"
	// channel_client "github.com/lingyicute/papercraft-server/app/service/biz/channel/client"
	chat_client "github.com/lingyicute/papercraft-server/app/service/biz/chat/client"
	dialog_client "github.com/lingyicute/papercraft-server/app/service/biz/dialog/client"
	user_client "github.com/lingyicute/papercraft-server/app/service/biz/user/client"
	idgen_client "github.com/lingyicute/papercraft-server/app/service/idgen/client"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMySQL(&c.Mysql)

	dao := &dao.Dao{
		Mysql:        dao.NewMysqlDao(db, c.MessageSharding),
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		UserClient:   user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:   chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		SyncClient:   sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.SyncClient)),
		DialogClient: dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
	}
	if c.BotSyncClient != nil {
		dao.BotSyncClient = sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.BotSyncClient))
	}

	return &ServiceContext{
		Config: c,
		Dao:    dao,
	}
}
