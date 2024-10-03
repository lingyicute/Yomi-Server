// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dao

import (
	"github.com/papercraft/proto/mtproto"
)

const (
	saltTimeout = 30 * 60 // salt timeout
)

func removeAllNil(vList []*mtproto.Authorization) []*mtproto.Authorization {
	for i := 0; i < len(vList); {
		if vList[i] != nil {
			i++
			continue
		}

		if i < len(vList)-1 {
			copy(vList[i:], vList[i+1:])
		}

		vList[len(vList)-1] = nil
		vList = vList[:len(vList)-1]
	}

	return vList
}
