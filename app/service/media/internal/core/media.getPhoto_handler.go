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
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/media/media"
)

// MediaGetPhoto
// media.getPhoto photo_id:long = Photo;
func (c *MediaCore) MediaGetPhoto(in *media.TLMediaGetPhoto) (*mtproto.Photo, error) {
	cData, err := c.svcCtx.Dao.GetCachePhotoData(c.ctx, in.GetPhotoId())
	if err != nil {
		c.Logger.Errorf("media.getPhotoFileData(%d) - error: %v", in.GetPhotoId(), err)
		return nil, err
	}

	return cData.ToPhoto(), nil
}
