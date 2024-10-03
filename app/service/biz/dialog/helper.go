/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package dialog_helper

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/internal/dal/dao/mysql_dao"
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}

type (
	DialogsDAO      = mysql_dao.DialogsDAO
	DialogsDO       = dataobject.DialogsDO
	SavedDialogsDAO = mysql_dao.SavedDialogsDAO
	SavedDialogsDO  = dataobject.SavedDialogsDO
)

var (
	NewDialogsDAO      = mysql_dao.NewDialogsDAO
	NewSavedDialogsDAO = mysql_dao.NewSavedDialogsDAO
)
