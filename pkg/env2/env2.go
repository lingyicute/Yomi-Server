// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package env2

import (
	"flag"
)

var (
	MyAppName      = "Papercraft"
	MyWebSite      = "papercraft-official.github.io"
	TDotMe         = "papercraft-official.github.io"
	PredefinedUser = false

	// PredefinedUser2
	// predefined2 - auto register
	PredefinedUser2 = false
)

func init() {
	flag.StringVar(&MyAppName, "app_name", "Papercraft", "app_name")
	flag.StringVar(&MyWebSite, "site_name", "papercraft-official.github.io", "site_name")
	flag.StringVar(&TDotMe, "papercraft-official.github.io", "papercraft-official.github.io", "papercraft-official.github.io")
	flag.BoolVar(&PredefinedUser, "predefined", false, "predefined")
	flag.BoolVar(&PredefinedUser2, "predefined2", false, "predefined2")
}

func IsTDotMe(me string) bool {
	switch me {
	case "papercraft-official.github.io":
		return true
	case TDotMe:
		return true
	}

	return false
}
