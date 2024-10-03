/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package main

import (
	"github.com/papercraft/marmota/pkg/commands"

	"github.com/lingyicute/papercraft-server/app/bff/premium/internal/server"
)

func main() {
	commands.Run(server.New())
}
