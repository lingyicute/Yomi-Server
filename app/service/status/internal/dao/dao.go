/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package dao

import (
	"github.com/papercraft/marmota/pkg/stores/kv"
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/status/internal/config"
)

type Dao struct {
	KV kv.Store
	mtproto.RPCNotificationClient
}

func New(c config.Config) *Dao {
	d := &Dao{
		KV: kv.NewStore(c.Status),
	}

	return d
}
