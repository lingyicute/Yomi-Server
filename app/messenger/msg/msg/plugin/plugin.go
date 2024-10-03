// Copyright 2022 Papercraft Authors
//  All rights reserved.
//
// Author: @lingyicute
//

package plugin

import (
	"context"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/username/username"
)

type MsgPlugin interface {
	ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error
	UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*mtproto.Peer, error)
	GetWebpagePreview(ctx context.Context, url string) (*mtproto.WebPage, error)
}
