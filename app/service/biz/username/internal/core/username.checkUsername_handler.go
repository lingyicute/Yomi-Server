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
	"github.com/lingyicute/papercraft-server/app/service/biz/username/username"
)

// UsernameCheckUsername
// username.checkUsername username:string = UsernameExist;
func (c *UsernameCore) UsernameCheckUsername(in *username.TLUsernameCheckUsername) (*username.UsernameExist, error) {
	var (
		checked = usernameNotExisted
	)

	usernameDO, _ := c.svcCtx.Dao.UsernameDAO.SelectByUsername(c.ctx, in.Username)
	if usernameDO != nil {
		checked = usernameExisted
	}

	return checked, nil
}
