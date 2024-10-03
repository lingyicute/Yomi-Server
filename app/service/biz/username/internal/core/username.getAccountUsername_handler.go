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
	"context"
	"fmt"
	"github.com/papercraft/marmota/pkg/stores/sqlx"

	"github.com/lingyicute/papercraft-server/app/service/biz/username/internal/dal/dataobject"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/username"
)

// UsernameGetAccountUsername
// username.getAccountUsername user_id:int = UsernameData;
func (c *UsernameCore) UsernameGetAccountUsername(in *username.TLUsernameGetAccountUsername) (*username.UsernameData, error) {
	v := new(dataobject.UsernameDO)

	err := c.svcCtx.CachedConn.QueryRow(
		c.ctx,
		v,
		fmt.Sprintf("username_%d", in.GetUserId()),
		func(ctx context.Context, db *sqlx.DB, v interface{}) error {
			usernameDO, err2 := c.svcCtx.UsernameDAO.SelectByUserId(c.ctx, in.GetUserId())
			if err2 == nil {
				*(v.(*dataobject.UsernameDO)) = *usernameDO
			}
			return err2
		})
	if err != nil {
		return nil, err
	}

	return username.MakeTLUsernameData(&username.UsernameData{
		Username: v.Username,
	}).To_UsernameData(), nil
}
