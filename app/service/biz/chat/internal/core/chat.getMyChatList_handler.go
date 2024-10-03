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
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/chat"
)

// ChatGetMyChatList
// chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;
func (c *ChatCore) ChatGetMyChatList(in *chat.TLChatGetMyChatList) (*chat.Vector_MutableChat, error) {
	var (
		chatList = make([]*mtproto.MutableChat, 0)
	)

	//
	if mtproto.FromBool(in.IsCreator) {
		c.svcCtx.Dao.ChatParticipantsDAO.SelectMyAdminListWithCB(
			c.ctx,
			in.UserId,
			func(sz, i int, v int64) {
				chat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, v, in.UserId)
				if err != nil {
					c.Logger.Errorf("chat.getMyChatList - error: %v", err)
				} else if chat != nil {
					chatList = append(chatList, chat)
				}
			})
	} else {
		// TODO:
	}

	return &chat.Vector_MutableChat{
		Datas: chatList,
	}, nil
}
