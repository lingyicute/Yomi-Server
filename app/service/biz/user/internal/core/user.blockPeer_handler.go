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

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (c *UserCore) UserBlockPeer(in *user.TLUserBlockPeer) (*mtproto.Bool, error) {
	blocked := c.svcCtx.Dao.BlockUser(c.ctx, in.GetUserId(), in.GetPeerId())

	return mtproto.ToBool(blocked), nil
}
