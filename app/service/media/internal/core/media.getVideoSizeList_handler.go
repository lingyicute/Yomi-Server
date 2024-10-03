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

// MediaGetVideoSizeList
// media.getVideoSizeList size_id = VideoSizeList;
func (c *MediaCore) MediaGetVideoSizeList(in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error) {
	cData, err := c.svcCtx.Dao.GetCachePhotoData(c.ctx, in.GetSizeId())
	if err != nil {
		c.Logger.Errorf("media.getVideoSizeList - error: %v", err)
		return nil, err
	}

	return cData.ToVideoSizeList(), nil
}
