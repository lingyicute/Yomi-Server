/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package config

import (
	kafka "github.com/papercraft/marmota/pkg/mq"
	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	InboxConsumer   kafka.KafkaConsumerConf
	Mysql           sqlx.Config
	KV              kv.KvConf
	IdgenClient     zrpc.RpcClientConf
	UserClient      zrpc.RpcClientConf
	ChatClient      zrpc.RpcClientConf
	DialogClient    zrpc.RpcClientConf
	SyncClient      *kafka.KafkaProducerConf
	BotSyncClient   *kafka.KafkaProducerConf `json:",optional"`
	MessageSharding int                      `json:",default=1"`
}
