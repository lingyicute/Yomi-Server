/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package authorization_helper

import (
	"github.com/lingyicute/papercraft-server/app/bff/authorization/internal/config"
	"github.com/lingyicute/papercraft-server/app/bff/authorization/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/bff/authorization/internal/svc"
	"github.com/lingyicute/papercraft-server/app/bff/authorization/plugin"
	"github.com/lingyicute/papercraft-server/pkg/code"
)

type (
	Config               = config.Config
	AuthorizationService = service.Service
)

func New(c Config, code2 code.VerifyCodeInterface, plugin plugin.AuthorizationPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, code2, plugin))
}
