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

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (c *MediaCore) MediaGetPhotoSizeList(in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	cData, err := c.svcCtx.Dao.GetCachePhotoData(c.ctx, in.GetSizeId())
	if err != nil {
		c.Logger.Errorf("media.getPhotoSizeList - error: %v", err)
		return nil, err
	}

	return cData.ToPhotoSizeList(), nil
}
