// Copyright 2024 Papercraft Authors
//  All rights reserved.
//
// Author: @lingyicute
//

package dao

import (
	"context"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/username"
)

func (d *Dao) ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error {
	return mtproto.ErrMethodNotImpl
}

func (d *Dao) UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*mtproto.Peer, error) {
	return d.UsernameClient.UsernameResolveUsername(ctx, in)
}

func (d *Dao) GetWebpagePreview(ctx context.Context, url string) (*mtproto.WebPage, error) {
	return nil, mtproto.ErrMethodNotImpl
}
