/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package mtproto

// import "time"

////////////////////////////////////////////////////////////////////////////////
func ToBool(b bool) *Bool {
	if b {
		return NewTLBoolTrue().To_Bool()
	} else {
		return NewTLBoolFalse().To_Bool()
	}
}

func FromBool(b *Bool) bool {
	return TLConstructor_CRC32_boolTrue == b.GetConstructor()
}

/*
//////////////////////////////////////////////////////////////////////////////////
// 太麻烦了
func GetUserIdListByChatParticipants(participants *TLChatParticipants) []int32 {
	chatUserIdList := []int32{}

	// TODO(@benqi):  nil check
	for _, participant := range participants.GetParticipants() {
		switch participant.Payload.(type) {
		case *ChatParticipant_ChatParticipant:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipant().GetUserId())
		case *ChatParticipant_ChatParticipantAdmin:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantAdmin().GetUserId())
		case *ChatParticipant_ChatParticipantCreator:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantCreator().GetUserId())
		}
	}
	return chatUserIdList
}

func (this *InputMedia) ToMessageMedia() (*MessageMedia) {
	switch this.Payload.(type) {
	case InputMedia_InputMediaUploadedPhoto:
		imedia := this.GetInputMediaUploadedPhoto()
		_ = TLInputMediaUploadedPhoto{}
		media := &TLMessageMediaPhoto{
			TtlSeconds: imedia.TtlSeconds,
			Caption: imedia.Caption,
		}

		p := &TLPhoto{
			HasStickers: len(imedia.Stickers) > 0,
			Date:int32(time.Now().Unix()),
		}

		switch imedia.GetFile().Payload.(type) {
		case *InputFile_InputFile:
			f := imedia.GetFile().GetInputFile()
			p.Id = f.Id
			p.AccessHash = 1 // f.Md5Checksum
			photoSize := &TLPhotoSize{
			}
			p.Sizes = append(p.Sizes, photoSize.ToPhotoSize())
		case *InputFile_InputFileBig:
			f := imedia.GetFile().GetInputFileBig()
			p.Id = f.Id
			p.AccessHash = 1 // f.Md5Checksum
			_ = TLInputFile{}
		}

		media.Photo = p.ToPhoto()
		return media.ToMessageMedia()
	}

	return nil
}
*/
