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
	"github.com/lingyicute/papercraft-server/app/service/biz/user/user"
)

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (c *UserCore) UserGetProfilePhotos(in *user.TLUserGetProfilePhotos) (*user.Vector_Long, error) {
	idList, err := c.svcCtx.Dao.UserProfilePhotosDAO.SelectList(c.ctx, in.UserId)
	if err != nil {
		c.Logger.Errorf("user.getProfilePhotos - error: %v", err)
	}

	return &user.Vector_Long{
		Datas: idList,
	}, nil
}
