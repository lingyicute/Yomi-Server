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
	"github.com/papercraft/marmota/pkg/stores/sqlc"
	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/plugin"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
	Plugin plugin.MessagePlugin
}

// New new a dao and return.
func New(c config.Config, plugin plugin.MessagePlugin) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:      newMysqlDao(db, c.MessageSharding),
		CachedConn: sqlc.NewConn(db, c.Cache),
		Plugin:     plugin,
	}
}
