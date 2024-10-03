/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package main

import (
	"github.com/papercraft/marmota/pkg/commands"

	"github.com/lingyicute/papercraft-server/app/service/biz/username/internal/server"
)

func main() {
	commands.Run(server.New())
}
