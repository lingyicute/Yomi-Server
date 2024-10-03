/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package svc

import (
	"github.com/papercraft/marmota/pkg/stores/sqlc"
	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMySQL(&c.Mysql)
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(db, sqlc.NewConn(db, c.Cache)),
	}
}
