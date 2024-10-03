/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package message_helper

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/dal/dao/mysql_dao"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/plugin"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.MessagePlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}

type (
	MessagesDAO = mysql_dao.MessagesDAO
	MessagesDO  = dataobject.MessagesDO
	// ChannelMessagesDAO   = mysql_dao.ChannelMessagesDAO
	// ChannelMessagesDO    = dataobject.ChannelMessagesDO
	// ScheduledMessagesDAO = mysql_dao.ScheduledMessagesDAO
	// ScheduledMessagesDO  = dataobject.ScheduledMessagesDO
)

var (
	NewMessagesDAO = mysql_dao.NewMessagesDAO
	// NewChannelMessagesDAO   = mysql_dao.NewChannelMessagesDAO
	// NewScheduledMessagesDAO = mysql_dao.NewScheduledMessagesDAO
)
