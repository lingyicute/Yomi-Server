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

// MessageUnPinAllMessages
// message.unPinAllMessages user_id:long peer_type:int peer_id:long = Vector<int>;
func (c *MessageCore) MessageUnPinAllMessages(in *message.TLMessageUnPinAllMessages) (*message.Vector_Int, error) {
	var (
		idList []int32
	)

	switch in.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		dialogId := mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
		idList, _ = c.svcCtx.Dao.MessagesDAO.SelectPinnedMessageIdList(c.ctx, in.UserId, dialogId.A, dialogId.B)
		if len(idList) > 0 {
			c.svcCtx.Dao.MessagesDAO.UpdateUnPinnedByIdList(c.ctx, in.UserId, idList)
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("message.unPinAllMessages blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	return &message.Vector_Int{
		Datas: idList,
	}, nil
}
