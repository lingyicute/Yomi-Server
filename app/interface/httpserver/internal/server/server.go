/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present, Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package server

import (
	"flag"

	"github.com/lingyicute/papercraft-server/app/interface/httpserver/internal/config"
	"github.com/lingyicute/papercraft-server/app/interface/httpserver/internal/server/http"
	"github.com/lingyicute/papercraft-server/app/interface/httpserver/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/httpserver.yaml", "the config file")

type Server struct {
	httpSrv *rest.Server
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	ctx := svc.NewServiceContext(c)

	s.httpSrv = http.New(ctx, c.Http)

	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
}
