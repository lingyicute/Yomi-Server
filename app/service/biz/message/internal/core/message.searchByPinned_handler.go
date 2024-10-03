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
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/message"
)

// MessageSearchByPinned
// message.searchByPinned user_id:long peer:PeerUtil = Vector<MessageBox>;
func (c *MessageCore) MessageSearchByPinned(in *message.TLMessageSearchByPinned) (*mtproto.MessageBoxList, error) {
	var (
		dialogId = mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
		boxList  = make([]*mtproto.MessageBox, 0)
	)

	switch in.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		c.svcCtx.Dao.MessagesDAO.SelectPinnedListWithCB(
			c.ctx,
			in.UserId,
			dialogId.A,
			dialogId.B,
			func(sz, i int, v *dataobject.MessagesDO) {
				boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
			})
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("message.searchByPinned blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	return mtproto.MakeTLMessageBoxList(&mtproto.MessageBoxList{
		BoxList: boxList,
	}).To_MessageBoxList(), nil
}
