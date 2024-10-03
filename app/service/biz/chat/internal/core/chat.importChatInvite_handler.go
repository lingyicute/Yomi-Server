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
	"time"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/chat"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatImportChatInvite
// chat.importChatInvite self_id:long hash:string = MutableChat;
func (c *ChatCore) ChatImportChatInvite(in *chat.TLChatImportChatInvite) (*mtproto.MutableChat, error) {
	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, in.Hash)
	if err != nil {
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		err = mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}

	if chatInviteDO.ExpireDate != 0 && time.Now().Unix() > chatInviteDO.ExpireDate {
		err = mtproto.ErrInviteHashExpired
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}
	if chatInviteDO.UsageLimit > 0 {
		sz := c.svcCtx.Dao.CommonDAO.CalcSize(c.ctx, "chat_invite_participants", map[string]interface{}{
			"link": chatInviteDO.Link,
		})

		if sz >= int(chatInviteDO.UsageLimit) {
			err = mtproto.ErrInviteHashExpired
			c.Logger.Errorf("chat.importChatInvite - error: %v", err)
			return nil, err
		}
	}

	chat2, err := c.ChatAddChatUser(&chat.TLChatAddChatUser{
		ChatId:    chatInviteDO.ChatId,
		InviterId: chatInviteDO.AdminId,
		UserId:    in.SelfId,
	})
	if err != nil {
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.ChatInviteParticipantsDAO.Insert(c.ctx, &dataobject.ChatInviteParticipantsDO{
		ChatId: chatInviteDO.ChatId,
		Link:   in.Hash,
		UserId: in.SelfId,
		Date2:  time.Now().Unix(),
	})

	return chat2, nil
}
