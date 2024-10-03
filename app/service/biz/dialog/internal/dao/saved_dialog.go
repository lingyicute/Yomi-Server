// Copyright 2024 Papercraft Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/internal/dal/dataobject"
)

func (d *Dao) MakeSavedDialog(do *dataobject.SavedDialogsDO) *mtproto.SavedDialog {
	return mtproto.MakeTLSavedDialog(&mtproto.SavedDialog{
		Pinned:     do.Pinned > 0,
		Peer:       mtproto.MakePeer(do.PeerType, do.PeerId),
		TopMessage: do.TopMessage,
	}).To_SavedDialog()
}
