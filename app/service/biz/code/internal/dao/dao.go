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
	"github.com/lingyicute/papercraft-server/app/service/biz/code/internal/config"

	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	kv kv.Store
}

func New(c config.Config) *Dao {
	return &Dao{
		kv: kv.NewStore(c.KV),
	}
}
