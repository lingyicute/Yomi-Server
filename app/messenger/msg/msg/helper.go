/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package msg_helper

import (
	"github.com/lingyicute/papercraft-server/app/messenger/msg/msg/internal/config"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/msg/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/msg/internal/svc"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/msg/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.MsgPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
