// Copyright 2022 Papercraft Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dao

import (
	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/lingyicute/papercraft-server/app/service/media/internal/dal/dao/mysql_dao"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.DocumentsDAO
	*mysql_dao.PhotosDAO
	*mysql_dao.PhotoSizesDAO
	*mysql_dao.VideoSizesDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:            db,
		DocumentsDAO:  mysql_dao.NewDocumentsDAO(db),
		PhotosDAO:     mysql_dao.NewPhotosDAO(db),
		PhotoSizesDAO: mysql_dao.NewPhotoSizesDAO(db),
		VideoSizesDAO: mysql_dao.NewVideoSizesDAO(db),
		CommonDAO:     sqlx.NewCommonDAO(db),
	}
}
