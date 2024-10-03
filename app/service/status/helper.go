/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package status_helper

import (
	"github.com/lingyicute/papercraft-server/app/service/status/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/status/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/status/internal/svc"
)

type (
	Config         = config.Config
	ServiceContext = svc.ServiceContext
)

var (
	New               = service.New
	NewServiceContext = svc.NewServiceContext
)
