// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package code

import (
	"github.com/papercraft/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ToAuthSentCode
// ///////////////////////////////////////////////////////////////////////////////////////////////////
// TODO: 如果手机号已经注册，检查是否有其他设备在线，有则使用sentCodeTypeApp
//
//	否则使用sentCodeTypeSms
//
// TODO: 有则使用sentCodeTypeFlashCall和entCodeTypeCall？？
func (m *PhoneCodeTransaction) ToAuthSentCode() *mtproto.Auth_SentCode {
	// TODO: only use sms

	authSentCode := mtproto.MakeTLAuthSentCode(&mtproto.Auth_SentCode{
		Type:          makeAuthSentCodeType(m.SentCodeType, len(m.PhoneCode), m.FlashCallPattern),
		PhoneCodeHash: m.PhoneCodeHash,
		NextType:      makeAuthCodeType(m.NextCodeType),
		Timeout:       &wrapperspb.Int32Value{Value: 60}, // TODO: 默认60s
	}).To_Auth_SentCode()
	if m.SentCodeType == CodeTypeApp {
		authSentCode.Timeout = nil
	}
	return authSentCode
}
