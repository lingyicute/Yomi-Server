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

// MediaUploadEncryptedFile
// media.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
func (c *MediaCore) MediaUploadEncryptedFile(in *media.TLMediaUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	// TODO: not impl
	c.Logger.Errorf("media.uploadEncryptedFile blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
