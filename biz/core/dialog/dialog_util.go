/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dialog

import "github.com/airwide-code/airwide.datacenter/mtproto"

type dialogItems struct {
	MessageIdList        []int32
	ChannelMessageIdMap  map[int32]int32
	UserIdList           []int32
	ChatIdList           []int32
	ChannelIdList        []int32
}

func makeDialogItems() *dialogItems {
	return &dialogItems{
		MessageIdList: make([]int32, 0),
		ChannelMessageIdMap: make(map[int32]int32, 0),
		UserIdList: make([]int32, 0),
		ChatIdList: make([]int32, 0),
		ChannelIdList: make([]int32, 0),
	}
}

func PickAllIDListByDialogs2(dialogs []*mtproto.Dialog) (items *dialogItems) {
	items = makeDialogItems()

	for _, d := range dialogs {
		dialog := d.To_Dialog()
		p := dialog.GetPeer()

		// TODO(@benqi): 先假设只有PEER_USER
		switch p.GetConstructor() {
		case mtproto.TLConstructor_CRC32_peerUser:
			items.MessageIdList = append(items.MessageIdList, dialog.GetTopMessage())
			items.UserIdList = append(items.UserIdList, p.GetData2().GetUserId())
		case mtproto.TLConstructor_CRC32_peerChat:
			items.MessageIdList = append(items.MessageIdList, dialog.GetTopMessage())
			items.ChatIdList = append(items.ChatIdList, p.GetData2().GetChatId())
		case mtproto.TLConstructor_CRC32_peerChannel:
			items.ChannelMessageIdMap[p.GetData2().GetChannelId()] = dialog.GetTopMessage()
			items.ChannelIdList = append(items.ChannelIdList, p.GetData2().GetChannelId())
		}
	}

	return
}
