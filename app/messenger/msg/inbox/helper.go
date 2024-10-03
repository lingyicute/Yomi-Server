/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package inbox_helper

import (
	kafka "github.com/papercraft/marmota/pkg/mq"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/inbox/internal/config"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/inbox/internal/core"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/inbox/internal/server/mq"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/inbox/internal/svc"
)

type (
	Config         = config.Config
	ServiceContext = svc.ServiceContext
	InboxCore      = core.InboxCore
)

var (
	NewServiceContext = svc.NewServiceContext
	NewInboxCore      = core.New
)

func New(c Config) *kafka.ConsumerGroup {
	return mq.New(svc.NewServiceContext(c), c.InboxConsumer)
}
