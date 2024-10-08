/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Yomi.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package status_helper

import (
	"github.com/teamgram/teamgram-server/app/service/status/internal/config"
	"github.com/teamgram/teamgram-server/app/service/status/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/status/internal/svc"
)

type (
	Config         = config.Config
	ServiceContext = svc.ServiceContext
)

var (
	New               = service.New
	NewServiceContext = svc.NewServiceContext
)
