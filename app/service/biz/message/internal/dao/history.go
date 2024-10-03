// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dao

import (
	"context"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/message/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

// GetOffsetIdBackwardHistoryMessages offset
func (d *Dao) GetOffsetIdBackwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectBackwardByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
		// logx.WithContext(ctx).Infof("GetOffsetIdBackwardHistoryMessages: %v", rList)
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetIdForwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectForwardByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateBackwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectBackwardByOffsetDateLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			int64(offsetDate),
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateForwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectForwardByOffsetDateLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			int64(offsetDate),
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

// GetOffsetIdBackwardUnreadMentions GetOffsetIdBackwardUnreadMentions
func (d *Dao) GetOffsetIdBackwardUnreadMentions(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectBackwardUnreadMentionsByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")
	}
	return
}

func (d *Dao) GetOffsetIdForwardUnreadMentions(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectForwardUnreadMentionsByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")
	}
	return
}
