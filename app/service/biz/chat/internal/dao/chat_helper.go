// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dao

import (
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/chat/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

func (d *Dao) MakeImmutableChatByDO(chatsDO *dataobject.ChatsDO) (chat *mtproto.ImmutableChat) {
	chat = &mtproto.ImmutableChat{
		Id:                     chatsDO.Id,
		Creator:                chatsDO.CreatorUserId,
		Title:                  chatsDO.Title,
		Photo:                  mtproto.MakeTLPhotoEmpty(nil).To_Photo(),
		Deactivated:            chatsDO.Deactivated,
		CallActive:             false,
		CallNotEmpty:           false,
		Noforwards:             chatsDO.Noforwards,
		ParticipantsCount:      chatsDO.ParticipantCount,
		Date:                   chatsDO.Date,
		Version:                chatsDO.Version,
		MigratedTo:             nil,
		DefaultBannedRights:    mtproto.BannedRights(chatsDO.DefaultBannedRights).ToChatBannedRights(),
		CanSetUsername:         false,
		About:                  chatsDO.About,
		ExportedInvite:         nil,
		BotInfo:                nil,
		Call:                   nil,
		AvailableReactionsType: chatsDO.AvailableReactionsType,
		AvailableReactions:     nil,
		TtlPeriod:              chatsDO.TtlPeriod,
	}

	if chatsDO.MigratedToId != 0 && chatsDO.MigratedToAccessHash != 0 {
		chat.MigratedTo = mtproto.MakeTLInputChannel(&mtproto.InputChannel{
			ChannelId:  chatsDO.MigratedToId,
			AccessHash: chatsDO.MigratedToAccessHash,
		}).To_InputChannel()
	}
	//
	////// chat_photo && photo
	//if chatsDO.PhotoId != 0 {
	//	chat.Photo, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
	//		PhotoId: chatsDO.PhotoId,
	//	})
	//}
	//if chat.Photo == nil {
	//	chat.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	//}

	chat.ExportedInvite = nil // model.ExportedChatInviteEmpty

	if chatsDO.AvailableReactionsType == mtproto.ChatReactionsTypeSome {
		jsonx.UnmarshalFromString(chatsDO.AvailableReactions, &chat.AvailableReactions)
	}

	return
}

func (d *Dao) MakeImmutableChatParticipant(chatParticipantsDO *dataobject.ChatParticipantsDO) (participant *mtproto.ImmutableChatParticipant) {
	participant = &mtproto.ImmutableChatParticipant{
		Id:              chatParticipantsDO.Id,
		ChatId:          chatParticipantsDO.ChatId,
		UserId:          chatParticipantsDO.UserId,
		State:           chatParticipantsDO.State,
		ParticipantType: chatParticipantsDO.ParticipantType,
		Link:            chatParticipantsDO.Link,
		InviterUserId:   chatParticipantsDO.InviterUserId,
		InvitedAt:       chatParticipantsDO.InvitedAt,
		KickedAt:        chatParticipantsDO.KickedAt,
		LeftAt:          chatParticipantsDO.LeftAt,
		AdminRights:     nil,
		IsBot:           chatParticipantsDO.IsBot,
		Date:            chatParticipantsDO.Date2,
	}

	if participant.ParticipantType == mtproto.ChatMemberAdmin {
		participant.AdminRights = mtproto.MakeDefaultChatAdminRights()
	}

	return
}
