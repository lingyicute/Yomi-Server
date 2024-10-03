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

// UserDeletePeerSettings
// user.deletePeerSettings user_id:int peer_type:int peer_id:int = Bool;
func (c *UserCore) UserDeletePeerSettings(in *user.TLUserDeletePeerSettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.DeleteUserPeerSettings(c.ctx, in.UserId, in.PeerType, in.PeerId)

	if err != nil {
		c.Logger.Errorf("user.deletePeerSettings - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
