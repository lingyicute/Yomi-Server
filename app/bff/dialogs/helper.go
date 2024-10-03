/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package dialogs_helper

import (
	"github.com/lingyicute/papercraft-server/app/bff/dialogs/internal/config"
	"github.com/lingyicute/papercraft-server/app/bff/dialogs/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/bff/dialogs/internal/svc"
	"github.com/lingyicute/papercraft-server/app/bff/dialogs/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.DialogsPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
