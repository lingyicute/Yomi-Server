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

// UserIsBot
// user.isBot id:long = Bool;
func (c *UserCore) UserIsBot(in *user.TLUserIsBot) (*mtproto.Bool, error) {
	userData := c.svcCtx.Dao.GetCacheUserData(c.ctx, in.GetId())
	if userData == nil {
		c.Logger.Errorf("user.isBot - error: invalid user(%d)", in.GetId())
		return nil, mtproto.ErrUserIdInvalid
	}

	return mtproto.ToBool(userData.GetUserData().GetBot() != nil), nil
}
