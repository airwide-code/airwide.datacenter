/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
    "github.com/golang/glog"
    "github.com/airwide-code/airwide.datacenter/mtproto"
    "golang.org/x/net/context"
    "github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
    "github.com/airwide-code/airwide.datacenter/baselib/logger"
    "github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
    "github.com/airwide-code/airwide.datacenter/biz/base"
    message2 "github.com/airwide-code/airwide.datacenter/biz/core/message"
    "time"
    "github.com/airwide-code/airwide.datacenter/biz/core/user"
)

func makeOutboxMessageBySendMultiMedia(authKeyId int64, fromId int32, peer *base.PeerUtil, request *mtproto.TLMessagesSendMultiMedia) ([]*mtproto.TLMessage, []int64) {
    multi_media := request.GetMultiMedia()
    messages := make([]*mtproto.TLMessage, 0, len(multi_media))
    randomIdList := make([]int64, 0, len(multi_media))
    grouped_id := base.NextSnowflakeId()
    for _, media := range multi_media {
        message := &mtproto.TLMessage{ Data2: &mtproto.Message_Data{
            Out:          true,
            Silent:       request.GetSilent(),
            FromId:       fromId,
            ToId:         peer.ToPeer(),
            ReplyToMsgId: request.GetReplyToMsgId(),
            Media: 		  makeMediaByInputMedia(authKeyId, media.GetData2().GetMedia()),
            // Entities:     media.GetData2()
            // ReplyMarkup: media.GetData2().GetReplyMarkup(),
            Date:         int32(time.Now().Unix()),
            GroupedId:    grouped_id,
        }}

        messages = append(messages, message)
        randomIdList = append(randomIdList, media.GetData2().GetRandomId())
    }

    return messages, randomIdList
}

func makeUpdateNewMessageListUpdates(selfUserId int32, boxList message2.MessageBoxList) *mtproto.TLUpdates {
    var messages []*mtproto.Message = make([]*mtproto.Message, 0, len(boxList))
    for _, box := range boxList {
        messages = append(messages, box.Message)
    }

    userIdList, _, _ := message2.PickAllIDListByMessages(messages)
    userList := user.GetUsersBySelfAndIDList(selfUserId, userIdList)
    updateNewList := make([]*mtproto.Update, 0, len(messages))
    for _, m := range messages {
        updateNewList = append(updateNewList, &mtproto.Update{
                Constructor: mtproto.TLConstructor_CRC32_updateNewMessage,
                Data2: &mtproto.Update_Data{
                    Message_1: m,
            }})
    }
    return &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
        Updates: updateNewList,
        Users:   userList,
        Date:    int32(time.Now().Unix()),
        Seq:     0,
    }}
}

// messages.sendMultiMedia#2095512f flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int multi_media:Vector<InputSingleMedia> = Updates;
func (s *MessagesServiceImpl) MessagesSendMultiMedia(ctx context.Context, request *mtproto.TLMessagesSendMultiMedia) (*mtproto.Updates, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("messages.sendMultiMedia#2095512f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    //md := grpc_util.RpcMetadataFromIncoming(ctx)
    //glog.Infof("messages.sendMedia#c8f16791 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): ???
    // request.NoWebpage
    // request.Background

    // peer
    var (
        peer *base.PeerUtil
        err error
    )

    if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
        err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
        glog.Error("messages.sendMedia#c8f16791 - invalid peer", err)
        return nil, err
    }
    // TODO(@benqi): check user or channels's access_hash

    // peer = helper.FromInputPeer2(md.UserId, request.GetPeer())
    if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
        peer = &base.PeerUtil{
            PeerType: base.PEER_USER,
            PeerId:   md.UserId,
        }
    } else {
        peer = base.FromInputPeer(request.GetPeer())
    }

    /////////////////////////////////////////////////////////////////////////////////////
    // 发件箱
    // sendMessageToOutbox
    outboxMessages, randomIdList := makeOutboxMessageBySendMultiMedia(md.AuthId, md.UserId, peer, request)
    var messageOutboxList message2.MessageBoxList
    for i := 0; i < len(outboxMessages); i++ {
        messageOutbox := message2.CreateMessageOutboxByNew(md.UserId, peer, randomIdList[i], outboxMessages[i].To_Message(), func(messageId int32) {
            // 更新会话信息
            user.CreateOrUpdateByOutbox(md.UserId, peer.PeerType, peer.PeerId, messageId, outboxMessages[i].GetMentioned(), request.GetClearDraft())
        })
        messageOutboxList = append(messageOutboxList, messageOutbox)
    }

    syncUpdates := makeUpdateNewMessageListUpdates(md.UserId, messageOutboxList)
    state, err := sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, md.UserId, syncUpdates.To_Updates())
    if err != nil {
        return nil, err
    }

    reply := SetupUpdatesState(state, syncUpdates)
    updateList := []*mtproto.Update{}
    for i := 0; i < len(messageOutboxList); i++ {
        updateMessageID := &mtproto.TLUpdateMessageID{Data2: &mtproto.Update_Data{
            Id_4:     messageOutboxList[i].MessageId,
            RandomId: randomIdList[i],
        }}
        updateList = append(updateList, updateMessageID.To_Update())
    }
    updateList = append(updateList, reply.GetUpdates()...)

    reply.SetUpdates(updateList)

    /////////////////////////////////////////////////////////////////////////////////////
    // 收件箱
    if request.GetPeer().GetConstructor() != mtproto.TLConstructor_CRC32_inputPeerSelf {
        // var inBoxes message2.MessageBoxList
        var inBoxeMap = map[int32][]*message2.MessageBox{}
        for i := 0; i < len(outboxMessages); i++ {
            inBoxes, _ := messageOutboxList[i].InsertMessageToInbox(md.UserId, peer, func(inBoxUserId, messageId int32) {
                // 更新会话信息
                switch peer.PeerType {
                case base.PEER_USER:
                    user.CreateOrUpdateByInbox(inBoxUserId, peer.PeerType, md.UserId, messageId, outboxMessages[i].GetMentioned())
                case base.PEER_CHAT, base.PEER_CHANNEL:
                    user.CreateOrUpdateByInbox(inBoxUserId, peer.PeerType, peer.PeerId, messageId, outboxMessages[i].GetMentioned())
                }
            })

            for j := 0; j < len(inBoxes); j++ {
                if boxList, ok := inBoxeMap[inBoxes[j].UserId]; !ok {
                    inBoxeMap[inBoxes[j].UserId] = []*message2.MessageBox{inBoxes[j]}
                } else {
                    boxList = append(boxList, inBoxes[j])
                    inBoxeMap[inBoxes[j].UserId] = boxList
                }
            }
        }

        for k, v := range  inBoxeMap {

            syncUpdates = makeUpdateNewMessageListUpdates(k, v)
            sync_client.GetSyncClient().PushToUserUpdatesData(k, syncUpdates.To_Updates())
        }
    }

    glog.Infof("messages.sendMultiMedia#2095512f - reply: %s", logger.JsonDebugData(reply))
    return reply.To_Updates(), nil
}
