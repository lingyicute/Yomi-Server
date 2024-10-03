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
	"github.com/lingyicute/papercraft-server/app/service/biz/message/message"
)

// MessageUpdatePinnedMessageId
// message.updatePinnedMessageId user_id:long peer_type:int peer_id:long id:int pinned:Bool = Bool;
func (c *MessageCore) MessageUpdatePinnedMessageId(in *message.TLMessageUpdatePinnedMessageId) (*mtproto.Bool, error) {
	switch in.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		_, err := c.svcCtx.Dao.MessagesDAO.UpdatePinned(c.ctx, mtproto.FromBool(in.Pinned), in.UserId, in.Id)
		if err != nil {
			c.Logger.Errorf("message.updatePinnedMessageId - error: %v", err)
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("message.updatePinnedMessageId blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	return mtproto.BoolTrue, nil
}
