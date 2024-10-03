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
	"github.com/lingyicute/papercraft-server/app/service/biz/code/code"
)

// CodeGetPhoneCode
// code.getPhoneCode auth_key_id:long phone:string phone_code_hash:string = PhoneCodeTransaction;
func (c *CodeCore) CodeGetPhoneCode(in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.svcCtx.Dao.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", in.AuthKeyId, in.Phone, err)
		err = mtproto.ErrPhoneCodeExpired
		return nil, err
	} else if codeData == nil {
		c.Logger.Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", in.AuthKeyId, in.Phone, err)
		err = mtproto.ErrPhoneCodeExpired
		return nil, err
	} else if codeData.PhoneCodeHash != in.PhoneCodeHash {
		c.Logger.Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s, phoneCodeHash: %s} error - invalid phoneCodeHash",
			in.AuthKeyId,
			in.Phone,
			in.PhoneCodeHash)
		err = mtproto.ErrPhoneCodeInvalid
		return nil, err
	}

	return codeData, nil
}
