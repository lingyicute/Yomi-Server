// Copyright (c) 2024-present, Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package http

import (
	"github.com/lingyicute/papercraft-server/app/interface/httpserver/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c rest.RestConf) *rest.Server {
	srv := rest.MustNewServer(c)

	go func(s *rest.Server, c *svc.ServiceContext) {
		defer s.Stop()

		RegisterHandlers(s, c)

		s.Start()
	}(srv, ctx)

	return srv
}
