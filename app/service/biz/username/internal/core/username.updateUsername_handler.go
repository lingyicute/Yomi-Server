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
	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/username"
)

// UsernameUpdateUsername
// username.updateUsername peer_type:int peer_id:int username:string = Bool;
func (c *UsernameCore) UsernameUpdateUsername(in *username.TLUsernameUpdateUsername) (*mtproto.Bool, error) {
	_, _, err := c.svcCtx.Dao.UsernameDAO.Insert(c.ctx, &dataobject.UsernameDO{
		Username: in.Username,
		PeerType: in.PeerType,
		PeerId:   in.PeerId,
	})
	if err != nil {
		if sqlx.IsDuplicate(err) {
			return mtproto.BoolFalse, nil
		} else {
			c.Logger.Errorf("username.updateUsername - error: %v", err)
			return nil, err
		}
	}

	return mtproto.BoolTrue, nil
}
