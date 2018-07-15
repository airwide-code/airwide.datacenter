/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package model

import "github.com/airwide-code/airwide.datacenter/mtproto"

type LangPacks struct {
	LangCode    string
	Version     int32
	Strings     []*mtproto.LangPackString_Data
	StringPluralizeds []*mtproto.LangPackString_Data
	StringDeleteds []*mtproto.LangPackString_Data
}
