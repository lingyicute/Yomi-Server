// Copyright 2024 Papercraft Authors
//  All rights reserved.
//
// Author: @lingyicute
//

package media_client

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type MediaClientHelper struct {
	cli MediaClient
}

func NewMediaClientHelper(cli zrpc.Client) *MediaClientHelper {
	return &MediaClientHelper{
		cli: NewMediaClient(cli),
	}
}

func (m *MediaClientHelper) Client() MediaClient {
	return m.cli
}
