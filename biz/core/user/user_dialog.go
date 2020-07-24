/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package user

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/golang/glog"
	base2 "github.com/airwide-code/airwide.datacenter/baselib/base"
	"time"
	"encoding/json"
)

func dialogDOToDialog(dialogDO* dataobject.UserDialogsDO) *mtproto.TLDialog {
	dialog := mtproto.NewTLDialog()
	// dialogData := dialog.GetData2()
	// draftIdList := make([]int32, 0)

	dialog.SetPinned(dialogDO.IsPinned == 1)
	dialog.SetPeer(base.ToPeerByTypeAndID(dialogDO.PeerType, dialogDO.PeerId))
	if dialogDO.PeerType == base.PEER_CHANNEL {
		// TODO(@benqi): only channel has pts
		// dialog.SetPts(messageBoxsDO.Pts)
		// peerChannlIdList = append(peerChannlIdList, dialogDO.PeerId)
		dialog.SetPts(dialogDO.Pts)
	}

	dialog.SetTopMessage(dialogDO.TopMessage)
	dialog.SetReadInboxMaxId(dialogDO.ReadInboxMaxId)
	dialog.SetReadOutboxMaxId(dialogDO.ReadOutboxMaxId)
	dialog.SetUnreadCount(dialogDO.UnreadCount)
	dialog.SetUnreadMentionsCount(dialogDO.UnreadMentionsCount)

	if dialogDO.DraftType == 2 {
		draft := &mtproto.DraftMessage{}
		err := json.Unmarshal([]byte(dialogDO.DraftMessageData), &draft)
		if err == nil {
			dialog.SetDraft(draft)
		}
	}

	// NotifySettings
	peerNotifySettings := mtproto.NewTLPeerNotifySettings()
	peerNotifySettings.SetShowPreviews(dialogDO.ShowPreviews == 1)
	peerNotifySettings.SetSilent(dialogDO.Silent == 1)
	peerNotifySettings.SetMuteUntil(dialogDO.MuteUntil)
	peerNotifySettings.SetSound(dialogDO.Sound)
	dialog.SetNotifySettings(peerNotifySettings.To_PeerNotifySettings())
	return dialog
}

func dialogDOListToDialogList(dialogDOList []dataobject.UserDialogsDO) (dialogs []*mtproto.Dialog) {
	draftIdList := make([]int32, 0)
	for _, dialogDO := range dialogDOList {
		if dialogDO.DraftId > 0 {
			draftIdList = append(draftIdList, dialogDO.DraftId)
		}
		dialogs = append(dialogs, dialogDOToDialog(&dialogDO).To_Dialog())
	}

	// TODO(@benqi): fetch draft message list
	return
}

func GetDialogsByOffsetId(userId int32, isPinned bool, offsetId int32, limit int32) (dialogs []*mtproto.Dialog) {
	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPinnedAndOffset(
		userId, base2.BoolToInt8(isPinned), offsetId, limit)
	glog.Infof("GetDialogsByOffsetId - dialogDOList: {%v}, query: {userId: %d, isPinned: %v, offeetId: %d, limit: %d ", dialogDOList, userId, isPinned, offsetId, limit)

	dialogs = dialogDOListToDialogList(dialogDOList)
	return
}

//func (m *dialogModel) GetDialogsByOffsetDate(userId int32, excludePinned bool, offsetData int32, limit int32) (dialogs []*mtproto.TLDialog) {
//	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectDialogsByPinnedAndOffsetDate(
//		userId, base2.BoolToInt8(!excludePinned), offsetData, limit)
//	for _, dialogDO := range dialogDOList {
//		dialogs = append(dialogs, dialogDOToDialog(&dialogDO))
//	}
//	return
//}

func GetDialogsByUserIDAndType(userId, peerType int32) (dialogs []*mtproto.Dialog) {
	dialogsDAO := dao.GetUserDialogsDAO(dao.DB_SLAVE)

	var dialogDOList []dataobject.UserDialogsDO
	if peerType == base.PEER_UNKNOWN || peerType == base.PEER_EMPTY {
		dialogDOList = dialogsDAO.SelectDialogsByUserID(userId)
		glog.Infof("SelectDialogsByUserID(%d) - {%v}", userId, dialogDOList)
	} else {
		dialogDOList = dialogsDAO.SelectDialogsByPeerType(userId, int8(peerType))
		glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogDOList)
	}

	dialogs = dialogDOListToDialogList(dialogDOList)
	// glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogs)
	return
}

func GetPinnedDialogs(userId int32) (dialogs []*mtproto.Dialog) {
	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectPinnedDialogs(userId)
	dialogs = dialogDOListToDialogList(dialogDOList)
	return
}

func GetPeersDialogs(selfId int32, peers []*mtproto.InputPeer) (dialogs []*mtproto.Dialog) {
	for _, peer := range peers {
		peerUtil := base.FromInputPeer2(selfId, peer)
		dialogDO := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPeer(selfId, int8(peerUtil.PeerType), peerUtil.PeerId)
		if dialogDO != nil {
			dialogs = append(dialogs, dialogDOToDialog(dialogDO).To_Dialog())
		}
	}
	return
}

// 发件箱
func CreateOrUpdateByOutbox(userId, peerType int32, peerId int32, topMessage int32, unreadMentions, clearDraft bool) {
	var (
		main = dao.GetUserDialogsDAO(dao.DB_MASTER)
		affectedRows = int64(0)
		date = int32(time.Now().Unix())
	)

	if clearDraft && unreadMentions {
		affectedRows = main.UpdateTopMessageAndMentionsAndClearDraft(topMessage, date, userId, int8(peerType), peerId)
	} else if clearDraft && !unreadMentions {
		affectedRows = main.UpdateTopMessageAndClearDraft(topMessage, date, userId, int8(peerType), peerId)
	} else if !clearDraft && unreadMentions {
		affectedRows = main.UpdateTopMessageAndMentions(topMessage, date, userId, int8(peerType), peerId)
	} else {
		affectedRows = main.UpdateTopMessage(topMessage, date, userId, int8(peerType), peerId)
	}

	if affectedRows == 0 {
		// 创建会话
		dialog := &dataobject.UserDialogsDO{}
		dialog.UserId = userId
		dialog.PeerType = int8(peerType)
		dialog.PeerId = peerId
		if unreadMentions {
			dialog.UnreadMentionsCount = 1
		} else {
			dialog.UnreadMentionsCount = 0
		}
		dialog.UnreadCount = 0
		dialog.TopMessage = topMessage
		dialog.CreatedAt = base2.NowFormatYMDHMS()
		dialog.Date2 = date
		main.Insert(dialog)
	}
	return
}

// 收件箱
func CreateOrUpdateByInbox(userId, peerType int32, peerId int32, topMessage int32, unreadMentions bool) {
	var (
		main = dao.GetUserDialogsDAO(dao.DB_MASTER)
		affectedRows = int64(0)
		date = int32(time.Now().Unix())
	)

	if unreadMentions {
		affectedRows = main.UpdateTopMessageAndUnreadAndMentions(topMessage, date, userId, int8(peerType), peerId)
	} else {
		affectedRows = main.UpdateTopMessageAndUnread(topMessage, date, userId, int8(peerType), peerId)
	}

	glog.Info("createOrUpdateByInbox - ", affectedRows)
	if affectedRows == 0 {
		// 创建会话
		dialog := &dataobject.UserDialogsDO{}
		dialog.UserId = userId
		dialog.PeerType = int8(peerType)
		dialog.PeerId = peerId
		if unreadMentions {
			dialog.UnreadMentionsCount = 1
		} else {
			dialog.UnreadMentionsCount = 0
		}
		dialog.UnreadCount = 1
		dialog.TopMessage = topMessage
		dialog.CreatedAt = base2.NowFormatYMDHMS()
		dialog.DraftMessageData = ""
		dialog.Date2 = date
		main.Insert(dialog)
	} else {

	}
	return
}

func SaveDraftMessage(userId int32, peerType int32, peerId int32, message *mtproto.DraftMessage) {
	var (
		main = dao.GetUserDialogsDAO(dao.DB_MASTER)
	)

	draft, _ := json.Marshal(message)
	main.SaveDraft(string(draft), userId, int8(peerType), peerId)
}
