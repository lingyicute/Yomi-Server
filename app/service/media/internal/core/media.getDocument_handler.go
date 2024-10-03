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
	"github.com/lingyicute/papercraft-server/app/service/media/media"
)

// MediaGetDocument
// media.getDocument id:long = Document;
func (c *MediaCore) MediaGetDocument(in *media.TLMediaGetDocument) (*mtproto.Document, error) {
	document := c.svcCtx.Dao.GetDocumentById(c.ctx, in.GetId())

	return document, nil
}
