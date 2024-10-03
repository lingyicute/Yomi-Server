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
	"github.com/lingyicute/papercraft-server/app/service/media/media"
)

// MediaGetDocumentList
// media.getDocumentList id_list:Vector<long> = Vector<Document>;
func (c *MediaCore) MediaGetDocumentList(in *media.TLMediaGetDocumentList) (*media.Vector_Document, error) {
	documents := c.svcCtx.Dao.GetDocumentListByIdList(c.ctx, in.IdList)

	return &media.Vector_Document{
		Datas: documents,
	}, nil
}
