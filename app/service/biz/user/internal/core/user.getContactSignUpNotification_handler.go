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
	"github.com/lingyicute/papercraft-server/app/service/biz/user/user"
)

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (c *UserCore) UserGetContactSignUpNotification(in *user.TLUserGetContactSignUpNotification) (*mtproto.Bool, error) {
	var (
		rV = false
	)

	if do, err := c.svcCtx.Dao.UserSettingsDAO.SelectByKey(c.ctx, in.UserId, "contactSignUpNotification"); do != nil {
		if do.Value == "true" {
			rV = true
		}
	} else if err == nil {
		rV = true
	} else {
		c.Logger.Errorf("user.getContactSignUpNotification - error: %v", err)
	}

	return mtproto.ToBool(rV), nil
}
