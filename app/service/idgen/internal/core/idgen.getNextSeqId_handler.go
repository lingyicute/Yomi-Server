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

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (c *IdgenCore) IdgenGetNextSeqId(in *idgen.TLIdgenGetNextSeqId) (*mtproto.Int64, error) {
	id, err := c.svcCtx.Dao.KV.IncrbyCtx(c.ctx, in.GetKey(), 1)
	if err != nil {
		c.Logger.Errorf("dgen.getNextSeqId(%s) error: %v", in.GetKey(), err)
		return nil, err
	}

	return &mtproto.Int64{
		V: id,
	}, nil
}
