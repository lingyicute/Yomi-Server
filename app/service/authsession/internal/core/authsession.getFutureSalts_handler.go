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
	"github.com/lingyicute/papercraft-server/app/service/authsession/authsession"
)

const (
	kDefaultSaltNum = 32
)

// AuthsessionGetFutureSalts
// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (c *AuthsessionCore) AuthsessionGetFutureSalts(in *authsession.TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	num := in.GetNum()
	if num == 0 {
		num = kDefaultSaltNum
	}
	futureSalts, err := c.svcCtx.Dao.GetFutureSalts(c.ctx, in.GetAuthKeyId(), num)
	if err != nil {
		c.Logger.Errorf("session.getFutureSalts - %v", err)
		return nil, err
	}

	return futureSalts.To_FutureSalts(), nil
}
