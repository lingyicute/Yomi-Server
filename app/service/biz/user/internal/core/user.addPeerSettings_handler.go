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

// UserAddPeerSettings
// user.addPeerSettings user_id:int peer_type:int peer_id:int settings:PeerSettings = Bool;
func (c *UserCore) UserAddPeerSettings(in *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.SetUserPeerSettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetPeerId(),
		in.GetSettings())

	if err != nil {
		c.Logger.Errorf("user.addPeerSettings - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
