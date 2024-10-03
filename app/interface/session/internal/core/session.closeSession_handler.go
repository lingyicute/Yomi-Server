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
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/interface/session/session"
)

// SessionCloseSession
// session.closeSession client:SessionClientEvent = Bool;
func (c *SessionCore) SessionCloseSession(in *session.TLSessionCloseSession) (*mtproto.Bool, error) {
	var (
		cli = in.GetClient()
	)

	if cli == nil {
		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("session.closeSession - error: %v", err)
		return nil, err
	}

	mainAuth := c.svcCtx.MainAuthMgr.GetMainAuthWrapper(cli.PermAuthKeyId)
	if mainAuth == nil {
		c.Logger.Errorf("session.closeSession - not found sessList by keyId: %d", cli)
	} else {
		mainAuth.SessionClientClosed(c.ctx, int(cli.KeyType), cli.AuthKeyId, cli.ServerId, cli.SessionId)
	}

	return mtproto.BoolTrue, nil
}
