// Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
//  All rights reserved.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package dialog

type (
	DialogPinnedExtList []*DialogPinnedExt
)

func (m DialogPinnedExtList) Add(peerType int32, peerId int64, order int64) DialogPinnedExtList {
	return append(m, &DialogPinnedExt{
		Order:    order,
		PeerType: peerType,
		PeerId:   peerId,
	})
}

func (m DialogPinnedExtList) Len() int {
	return len(m)
}
func (m DialogPinnedExtList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m DialogPinnedExtList) Less(i, j int) bool {
	// TODO: if date[i] == date[j]
	return m[i].Order < m[j].Order
}
