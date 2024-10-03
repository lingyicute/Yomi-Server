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

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (c *AuthsessionCore) AuthsessionGetUserId(in *authsession.TLAuthsessionGetUserId) (*mtproto.Int64, error) {
	var (
		inKeyId = in.GetAuthKeyId()
	)

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, inKeyId)
	if err != nil {
		c.Logger.Errorf("queryAuthKeyV2(%d) is error: %v", inKeyId, err)
		return nil, err
	} else if keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("queryAuthKeyV2(%d) - PermAuthKeyId is empty", inKeyId)
		return nil, mtproto.ErrAuthKeyPermEmpty
	}

	userId := c.svcCtx.Dao.GetAuthKeyUserId(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: userId,
	}).To_Int64(), nil
}
