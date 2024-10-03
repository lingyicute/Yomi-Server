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
	"github.com/lingyicute/papercraft-server/app/service/idgen/idgen"
)

// IdgenNextId
// idgen.nextId = Int64;
func (c *IdgenCore) IdgenNextId(in *idgen.TLIdgenNextId) (*mtproto.Int64, error) {
	_ = in

	return &mtproto.Int64{
		V: c.svcCtx.Dao.Node.Generate().Int64(),
	}, nil
}
