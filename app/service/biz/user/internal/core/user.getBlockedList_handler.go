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
	"github.com/lingyicute/papercraft-server/app/service/biz/user/user"
)

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (c *UserCore) UserGetBlockedList(in *user.TLUserGetBlockedList) (*user.Vector_PeerBlocked, error) {
	// TODO: recache
	doList, _ := c.svcCtx.Dao.UserPeerBlocksDAO.SelectList(c.ctx, in.UserId, in.Limit)
	bockedList := &user.Vector_PeerBlocked{
		Datas: make([]*mtproto.PeerBlocked, 0, len(doList)),
	}

	for _, do := range doList {
		switch do.PeerType {
		case mtproto.PEER_USER:
			blocked := mtproto.MakeTLPeerBlocked(&mtproto.PeerBlocked{
				PeerId: mtproto.MakePeerUser(do.PeerId),
				Date:   int32(do.Date),
			})
			bockedList.Datas = append(bockedList.Datas, blocked.To_PeerBlocked())
		case mtproto.PEER_CHAT:
			blocked := mtproto.MakeTLPeerBlocked(&mtproto.PeerBlocked{
				PeerId: mtproto.MakePeerChat(do.PeerId),
				Date:   int32(do.Date),
			})
			bockedList.Datas = append(bockedList.Datas, blocked.To_PeerBlocked())
		case mtproto.PEER_CHANNEL:
			blocked := mtproto.MakeTLPeerBlocked(&mtproto.PeerBlocked{
				PeerId: mtproto.MakePeerChannel(do.PeerId),
				Date:   int32(do.Date),
			})
			bockedList.Datas = append(bockedList.Datas, blocked.To_PeerBlocked())
		}
	}

	return bockedList, nil
}
