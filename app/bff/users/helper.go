/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package users_helper

import (
	"github.com/lingyicute/papercraft-server/app/bff/users/internal/config"
	"github.com/lingyicute/papercraft-server/app/bff/users/internal/plugin"
	"github.com/lingyicute/papercraft-server/app/bff/users/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/bff/users/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.StoryPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
