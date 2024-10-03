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

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (c *AuthsessionCore) AuthsessionGetPushSessionId(in *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
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

	sessionId := c.svcCtx.Dao.GetPushSessionId(
		c.ctx,
		in.GetUserId(),
		keyData.PermAuthKeyId,
		in.GetTokenType())
	if sessionId == 0 {
		c.Logger.Errorf("not found pushSessionId(%d,%d)", in.GetUserId(), keyData.PermAuthKeyId)
	}

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: sessionId,
	}).To_Int64(), nil
}
