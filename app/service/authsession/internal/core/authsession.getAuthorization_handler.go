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

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (c *AuthsessionCore) AuthsessionGetAuthorization(in *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
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

	authorization, err := c.svcCtx.Dao.GetAuthorization(c.ctx, keyData.PermAuthKeyId)
	if err != nil {
		c.Logger.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	}

	return authorization, nil
}
