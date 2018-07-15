/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package document

import (
	"testing"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"fmt"
	"encoding/json"
)

func TestDocumentAttributes(t *testing.T) {
	attributes := &mtproto.DocumentAttributeList{}
	imageSize := &mtproto.TLDocumentAttributeImageSize{Data2: &mtproto.DocumentAttribute_Data{
		W: 512,
		H: 512,
	}}
	attributes.Attributes = append(attributes.Attributes, imageSize.To_DocumentAttribute())

	sticker := &mtproto.TLDocumentAttributeSticker{Data2: &mtproto.DocumentAttribute_Data{
		Alt: "😂",
		Stickerset: &mtproto.InputStickerSet{
			Constructor: mtproto.TLConstructor_CRC32_inputStickerSetID,
			Data2: &mtproto.InputStickerSet_Data{
				Id: 835404231795015689,
				AccessHash:987465871030319816,
			},
		},
	}}
	attributes.Attributes = append(attributes.Attributes, sticker.To_DocumentAttribute())

	fileName := &mtproto.TLDocumentAttributeFilename{Data2: &mtproto.DocumentAttribute_Data{
		FileName: "sticker.webp",
	}}

	attributes.Attributes = append(attributes.Attributes, fileName.To_DocumentAttribute())
	d, _ := json.Marshal(attributes)
	fmt.Println(string(d))
}
