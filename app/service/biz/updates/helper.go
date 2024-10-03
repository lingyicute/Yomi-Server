/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package updates_helper

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/updates/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/updates/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/biz/updates/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
