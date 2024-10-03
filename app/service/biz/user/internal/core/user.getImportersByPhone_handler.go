// Copyright 2022 Papercraft Authors
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
	"github.com/lingyicute/papercraft-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/user/user"
)

// UserGetImportersByPhone
// user.getImportersByPhone phone:string = Vector<InputContact>;
func (c *UserCore) UserGetImportersByPhone(in *user.TLUserGetImportersByPhone) (*user.Vector_InputContact, error) {
	contacts := &user.Vector_InputContact{
		Datas: make([]*mtproto.InputContact, 0),
	}

	c.svcCtx.Dao.UnregisteredContactsDAO.SelectImportersByPhoneWithCB(
		c.ctx,
		in.Phone,
		func(sz, i int, v *dataobject.UnregisteredContactsDO) {
			contacts.Datas = append(contacts.Datas, mtproto.MakeTLInputPhoneContact(&mtproto.InputContact{
				ClientId:  v.ImporterUserId,
				Phone:     "",
				FirstName: v.ImportFirstName,
				LastName:  v.ImportLastName,
			}).To_InputContact())
		})

	return contacts, nil
}
