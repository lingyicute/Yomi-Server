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
	"github.com/lingyicute/papercraft-server/app/messenger/msg/inbox/inbox"
	"github.com/lingyicute/papercraft-server/app/messenger/msg/msg/msg"
	"github.com/lingyicute/papercraft-server/app/messenger/sync/sync"
)

// MsgDeletePhoneCallHistory
// msg.deletePhoneCallHistory user_id:long auth_key_id:long revoke:Bool = messages.AffectedFoundMessages;
func (c *MsgCore) MsgDeletePhoneCallHistory(in *msg.TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	var (
		pts, ptsCount int32
		msgIdList     []int32
		msgDataIdList []int64
		err           error
	)

	if msgIdList, msgDataIdList, err = c.svcCtx.Dao.DeletePhoneCallHistory(c.ctx, in.UserId); err != nil {
		c.Logger.Errorf("DeleteMessages - %v", err)
		return nil, err
	}

	pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.UserId, len(msgDataIdList))
	ptsCount = int32(len(msgDataIdList))

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:        in.UserId,
			PermAuthKeyId: in.AuthKeyId,
			Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
				Messages:  msgIdList,
				Pts_INT32: pts,
				PtsCount:  ptsCount,
			}).To_Update()),
		})

	if in.Revoke {
		c.svcCtx.Dao.InboxClient.InboxDeleteMessagesToInbox(
			c.ctx,
			&inbox.TLInboxDeleteMessagesToInbox{
				FromId: in.UserId,
				Id:     msgDataIdList,
			})
	}

	return mtproto.MakeTLMessagesAffectedFoundMessages(&mtproto.Messages_AffectedFoundMessages{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
		Messages: msgIdList,
	}).To_Messages_AffectedFoundMessages(), nil
}
