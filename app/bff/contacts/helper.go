/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package contacts_helper

import (
	"github.com/lingyicute/papercraft-server/app/bff/contacts/internal/config"
	"github.com/lingyicute/papercraft-server/app/bff/contacts/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/bff/contacts/internal/svc"
	"github.com/lingyicute/papercraft-server/app/bff/contacts/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.ContactsPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
