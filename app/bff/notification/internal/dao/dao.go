// Copyright 2022 Papercraft Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dao

import (
	kafka "github.com/papercraft/marmota/pkg/mq"
	"github.com/papercraft/marmota/pkg/net/rpcx"
	"github.com/lingyicute/papercraft-server/app/bff/notification/internal/config"
	sync_client "github.com/lingyicute/papercraft-server/app/messenger/sync/client"
	chat_client "github.com/lingyicute/papercraft-server/app/service/biz/chat/client"
	user_client "github.com/lingyicute/papercraft-server/app/service/biz/user/client"
)

type Dao struct {
	user_client.UserClient
	chat_client.ChatClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient: user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient: chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		SyncClient: sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
