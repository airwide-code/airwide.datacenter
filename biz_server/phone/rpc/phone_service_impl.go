/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

// Before a voice call is ready, some preliminary actions have to be performed.
// The calling party needs to contact the party to be called and check whether it is ready to accept the call.
// Besides that, the parties have to negotiate the protocols to be used,
// learn the IP addresses of each other or of the Telegram relay servers to be used (so-called reflectors),
// and generate a one-time encryption key for this voice call with the aid of Diffie—Hellman key exchange.
// All of this is accomplished in parallel with the aid of several Telegram API methods and related notifications.
//

var (
	fingerprint uint64 = 12240908862933197005
)
const (
	PHONE_STATE_UNKNOWN = iota
	PHONE_STATE_REQUEST_CALL
)

type phoneCallState int

type phoneCallSession struct {
	id                    int64
	adminId               int32
	adminAccessHash       int64
	participantId         int32
	participantAccessHash int64
	date                  int32
	state                 int 		// phoneCallstate
	protocol              *mtproto.TLPhoneCallProtocol
	g_b                   []byte	// acceptCall
	g_a                   []byte	// confirm
}

// TODO(@benqi): 存储到redis里
var phoneCallSessionManager = make(map[int64]*phoneCallSession)

type PhoneServiceImpl struct {
}
