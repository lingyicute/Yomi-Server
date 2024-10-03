/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package qrcode_helper

import (
	"github.com/lingyicute/papercraft-server/app/bff/qrcode/internal/config"
	"github.com/lingyicute/papercraft-server/app/bff/qrcode/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/bff/qrcode/internal/svc"
	"github.com/lingyicute/papercraft-server/app/bff/qrcode/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.QrcodePlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
