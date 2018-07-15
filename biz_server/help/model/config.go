/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package model

// TODO(@benqi): 配置中心管理配置
// dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
type DcOption struct {
	Ipv6      bool
	MediaOnly bool
	TcpoOnly  bool
	Cdn       bool
	Static    bool
	Id        int32
	IpAddress string
	Port      int32
}

type DisabledFeature struct {
	Feature     string
	Description string
}

type Config struct {
	PhonecallsEnabled       bool
	DefaultP2pContacts      bool
	Date                    int32
	Expires                 int32
	TestMode                bool
	ThisDc                  int32
	DcOptions               []DcOption
	ChatSizeMax             int32
	MegagroupSizeMax        int32
	ForwardedCountMax       int32
	OnlineUpdatePeriodMs    int32
	OfflineBlurTimeoutMs    int32
	OfflineIdleTimeoutMs    int32
	OnlineCloudTimeoutMs    int32
	NotifyCloudDelayMs      int32
	NotifyDefaultDelayMs    int32
	ChatBigSize             int32
	PushChatPeriodMs        int32
	PushChatLimit           int32
	SavedGifsLimit          int32
	EditTimeLimit           int32
	RatingEDecay            int32
	StickersRecentLimit     int32
	StickersFavedLimit      int32
	ChannelsReadMediaPeriod int32
	TmpSessions             int32
	PinnedDialogsCountMax   int32
	CallReceiveTimeoutMs    int32
	CallRingTimeoutMs       int32
	CallConnectTimeoutMs    int32
	CallPacketTimeoutMs     int32
	MeUrlPrefix             string
	SuggestedLangCode       string
	LangPackVersion         int32
	DisabledFeatures        []DisabledFeature
}
