/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package core

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/updates/updates"
)

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (c *UpdatesCore) UpdatesGetChannelDifferenceV2(in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	c.Logger.Errorf("updates.getChannelDifferenceV2 blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")

	return updates.MakeTLChannelDifference(&updates.ChannelDifference{
		Final:        false,
		Pts:          in.Pts,
		NewMessages:  nil,
		OtherUpdates: nil,
	}).To_ChannelDifference(), nil
}
