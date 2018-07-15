/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package message

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"encoding/json"
	"time"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
	base2 "github.com/airwide-code/airwide.datacenter/baselib/base"
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/golang/glog"
	"fmt"
)

type ChannelMessageBox struct {
	SenderUserId        int32
	ChannelId           int32
	ChannelMessageBoxId int32
	MessageId           int64
	RandomId            int64
	Message             *mtproto.Message
}

type ChannelBoxCreated func(int32)

func CreateChannelMessageBoxByNew(fromId, channelId int32, clientRandomId int64, message2 *mtproto.Message, cb ChannelBoxCreated) (box *ChannelMessageBox) {
	now := int32(time.Now().Unix())
	boxId := int32(update2.NextChannelMessageBoxId(base2.Int32ToString(channelId)))
	messageDatasDO := &dataobject.MessageDatasDO{
		DialogId:     int64(-channelId),
		MessageId:    base.NextSnowflakeId(),
		SenderUserId: fromId,
		PeerType:     int8(base.PEER_CHANNEL),
		PeerId:       channelId,
		RandomId:     clientRandomId,
		Date:         now,
		Deleted:      0,
	}

	channelMessageBoxesDO := &dataobject.ChannelMessageBoxesDO{
		SenderUserId:        fromId,
		ChannelId:           channelId,
		ChannelMessageBoxId: boxId,
		MessageId:           messageDatasDO.MessageId,
		Date:                now,
	}

	switch message2.GetConstructor() {
	case mtproto.TLConstructor_CRC32_messageEmpty:
		messageDatasDO.MessageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.TLConstructor_CRC32_message:
		messageDatasDO.MessageType = MESSAGE_TYPE_MESSAGE
		message := message2.To_Message()

		// mentioned = message.GetMentioned()
		message.SetId(channelMessageBoxesDO.ChannelMessageBoxId)
	case mtproto.TLConstructor_CRC32_messageService:
		messageDatasDO.MessageType = MESSAGE_TYPE_MESSAGE_SERVICE
		message := message2.To_MessageService()

		// mentioned = message.GetMentioned()
		message.SetId(channelMessageBoxesDO.ChannelMessageBoxId)
	}

	messageData, _ := json.Marshal(message2)
	messageDatasDO.MessageData = string(messageData)

	//// TODO(@benqi): pocess clientRandomId dup
	dao.GetMessageDatasDAO(dao.DB_MASTER).Insert(messageDatasDO)
	dao.GetChannelMessageBoxesDAO(dao.DB_MASTER).Insert(channelMessageBoxesDO)

	box = &ChannelMessageBox{
		SenderUserId:          fromId,
		ChannelId:  channelId,
		ChannelMessageBoxId: boxId,
		MessageId: channelMessageBoxesDO.MessageId,
		RandomId:        clientRandomId,
		Message:         message2,
	}

	if cb != nil {
		cb(box.ChannelMessageBoxId)
	}

	return
}

func doToChannelMessage(do *dataobject.MessageDatasDO) (*mtproto.Message, error) {
	message := &mtproto.Message{
		Data2: &mtproto.Message_Data{},
	}

	switch do.MessageType {
	case MESSAGE_TYPE_MESSAGE_EMPTY:
		message.Constructor = mtproto.TLConstructor_CRC32_messageEmpty
		// message = message2
	case MESSAGE_TYPE_MESSAGE:
		// err := proto.Unmarshal(messageDO.MessageData, message)
		err := json.Unmarshal([]byte(do.MessageData), message)
		if err != nil {
			glog.Errorf("messageDOToMessage - Unmarshal message(%d)error: %v", do.Id, err)
			return nil, err
		} else {
			message.Constructor = mtproto.TLConstructor_CRC32_message
		}
	case MESSAGE_TYPE_MESSAGE_SERVICE:
		err := json.Unmarshal([]byte(do.MessageData), message)
		if err != nil {
			glog.Errorf("messageDOToMessage - Unmarshal message(%d)error: %v", do.Id, err)
			return nil, err
		} else {
			message.Constructor = mtproto.TLConstructor_CRC32_messageService
		}
	default:
		err := fmt.Errorf("messageDOToMessage - Invalid messageType, db's data error, message(%d)", do.Id)
		glog.Error(err)
		return nil, err
	}

	return message, nil
}

func GetChannelMessage(channelId int32, id int32) (message *mtproto.Message) {
	do := dao.GetChannelMessageBoxesDAO(dao.DB_SLAVE).SelectByMessageId(channelId, id)
	if do == nil {
		return
	}

	messageDO := dao.GetMessageDatasDAO(dao.DB_SLAVE).SelectByMessageId(do.MessageId)
	message, _ = doToChannelMessage(messageDO)
	return
}

func GetChannelMessageList(channelId int32, idList []int32) (messages []*mtproto.Message) {
	if len(idList) == 0 {
		messages = []*mtproto.Message{}
	} else {
		doList := dao.GetChannelMessageBoxesDAO(dao.DB_SLAVE).SelectByMessageIdList(channelId, idList)
		messages = make([]*mtproto.Message, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			// TODO(@benqi): check data
			messageDO := dao.GetMessageDatasDAO(dao.DB_SLAVE).SelectByMessageId(doList[i].MessageId)
			if messageDO == nil {
				continue
			}
			m, _ := doToChannelMessage(messageDO)
			if m != nil {
				messages = append(messages, m)
			}
		}
	}
	return
}
