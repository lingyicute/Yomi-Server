/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package core

import (
	"github.com/lingyicute/papercraft-server/app/service/media/media"
)

// MediaGetPhotoSizeListList
// media.getPhotoSizeListList id_list:Vector<long> = Vector<PhotoSizeList>;
func (c *MediaCore) MediaGetPhotoSizeListList(in *media.TLMediaGetPhotoSizeListList) (*media.Vector_PhotoSizeList, error) {
	szListList := c.svcCtx.Dao.GetPhotoSizeListList(c.ctx, in.GetIdList())
	photoSizeListList := &media.Vector_PhotoSizeList{
		Datas: make([]*media.PhotoSizeList, 0, len(szListList)),
	}

	for szId, szList := range szListList {
		photoSizeListList.Datas = append(photoSizeListList.Datas, &media.PhotoSizeList{
			SizeId: szId,
			Sizes:  szList,
			DcId:   1,
		})
	}

	return photoSizeListList, nil
}
