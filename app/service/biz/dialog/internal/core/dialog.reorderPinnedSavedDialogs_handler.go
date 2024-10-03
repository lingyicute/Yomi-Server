// Copyright 2024 Papercraft Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package core

import (
	"time"

	"github.com/papercraft/marmota/pkg/stores/sqlx"
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/dialog/dialog"
)

// DialogReorderPinnedSavedDialogs
// dialog.reorderPinnedSavedDialogs user_id:long force:Bool order:Vector<PeerUtil> = Bool;
func (c *DialogCore) DialogReorderPinnedSavedDialogs(in *dialog.TLDialogReorderPinnedSavedDialogs) (*mtproto.Bool, error) {
	var (
		userId      = in.GetUserId()
		force       = mtproto.FromBool(in.GetForce())
		orderPinned = time.Now().Unix()
	)

	sqlx.TxWrapper(
		c.ctx,
		c.svcCtx.Dao.DB,
		func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			if force {
				_, result.Err = c.svcCtx.Dao.SavedDialogsDAO.UpdateUserUnPinnedTx(
					tx,
					userId)
				if result.Err != nil {
					return
				}
			}

			for _, id := range in.Order {
				_, result.Err = c.svcCtx.Dao.SavedDialogsDAO.UpdateUserPeerPinnedTx(
					tx,
					orderPinned<<32,
					in.UserId,
					id.PeerType,
					id.PeerId)
				if result.Err != nil {
					return
				}
				orderPinned -= 1
			}
		})

	return mtproto.BoolTrue, nil
}
