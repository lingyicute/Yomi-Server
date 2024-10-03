/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package httpserverhelper

import (
	"github.com/lingyicute/papercraft-server/app/interface/httpserver/internal/server"
	"github.com/zeromicro/go-zero/zrpc"
)

func init() {
	zrpc.DontLogClientContentForMethod("/session.RPCSession/session_sendHttpDataToSession")
}

type (
	Server = server.Server
)
