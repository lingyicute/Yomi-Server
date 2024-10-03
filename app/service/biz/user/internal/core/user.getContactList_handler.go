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

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (c *UserCore) UserGetContactList(in *user.TLUserGetContactList) (*user.Vector_ContactData, error) {
	rValList := &user.Vector_ContactData{}

	rValList.Datas = c.svcCtx.Dao.GetUserContactList(c.ctx, in.GetUserId())
	if rValList.Datas == nil {
		rValList.Datas = []*mtproto.ContactData{}
	}

	return rValList, nil
}
