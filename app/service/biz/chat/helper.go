/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package chat_helper

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/dal/dao/mysql_dao"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/svc"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/plugin"
)

type (
	Config = config.Config
)

type (
	ChatParticipantsDAO = mysql_dao.ChatParticipantsDAO
	ChatParticipantsDO  = dataobject.ChatParticipantsDO
)

var (
	NewChatParticipantsDAO = mysql_dao.NewChatParticipantsDAO
)

func New(c Config, plugin plugin.ChatPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
