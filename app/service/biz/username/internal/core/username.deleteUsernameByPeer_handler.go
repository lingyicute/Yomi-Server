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

// UsernameDeleteUsernameByPeer
// username.deleteUsernameByPeer peer_type:int peer_id:int = Bool;
func (c *UsernameCore) UsernameDeleteUsernameByPeer(in *username.TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UsernameDAO.Delete2(c.ctx, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("username.deleteUsernameByPeer - error: %v")
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
