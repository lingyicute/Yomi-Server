/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package user_helper

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/dao"
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/svc"
)

type (
	Dao           = dao.Dao
	CacheUserData = dao.CacheUserData
)

var (
	GenCacheUserDataCacheKey = dao.GenCacheUserDataCacheKey
)

type (
	Config  = config.Config
	Service = service.Service
)

var (
	NewServiceContext = svc.NewServiceContext
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
