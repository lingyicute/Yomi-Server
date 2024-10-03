// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dao

import (
	"path"
	"strings"

	"github.com/papercraft/proto/mtproto"
)

var (
	emptyPhoto = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
)

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	if ext == "" {
		ext = ".partial"
	}
	return strings.ToLower(ext)
}

func makePhotoEmpty(id int64) *mtproto.Photo {
	return mtproto.MakeTLPhotoEmpty(&mtproto.Photo{
		Id: id,
	}).To_Photo()
}
