// Copyright 2024 Papercraft Authors
//  All rights reserved.
//
// Author: @lingyicute
//

package dao

import (
	"github.com/lingyicute/papercraft-server/app/interface/gnetway/internal/config"
)

type Dao struct {
	*ShardingSessionClient
}

func New(c config.Config) *Dao {
	return &Dao{
		ShardingSessionClient: NewShardingSessionClient(c),
	}
}
