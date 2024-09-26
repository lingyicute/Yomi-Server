/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package privacysettingshelper

import (
	"github.com/teamgram/teamgram-server/app/bff/privacysettings/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/privacysettings/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/privacysettings/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
