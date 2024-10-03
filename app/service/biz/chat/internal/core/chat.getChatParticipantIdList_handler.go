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
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatGetChatParticipantIdList
// chat.getChatParticipantIdList chat_id:long = Vector<long>;
func (c *ChatCore) ChatGetChatParticipantIdList(in *chat.TLChatGetChatParticipantIdList) (*chat.Vector_Long, error) {
	var (
		idList = make([]int64, 0)
	)

	c.svcCtx.Dao.ChatParticipantsDAO.SelectListWithCB(
		c.ctx,
		in.ChatId,
		func(sz, i int, v *dataobject.ChatParticipantsDO) {
			if v.State != mtproto.ChatMemberStateNormal {
				return
			}
			idList = append(idList, v.UserId)
		})

	return &chat.Vector_Long{
		Datas: idList,
	}, nil
}
