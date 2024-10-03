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
	"github.com/lingyicute/papercraft-server/app/service/biz/username/username"
)

// UsernameDeleteUsername
// username.deleteUsername username:string = Bool;
func (c *UsernameCore) UsernameDeleteUsername(in *username.TLUsernameDeleteUsername) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UsernameDAO.Delete(c.ctx, in.Username)
	if err != nil {
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
