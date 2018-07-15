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
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
	"time"
)

// help.getConfig#c4f9186b = Config;
func (s *HelpServiceImpl) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("help.getConfig#c4f9186b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 设置Reply对象累死人了, 得想个办法实现model和mtproto的自动转换
	helpConfig := mtproto.NewTLConfig()
	// &mtproto.TLConfig{}
	helpConfig.SetPhonecallsEnabled(config.PhonecallsEnabled)
	helpConfig.SetDefaultP2PContacts(config.DefaultP2pContacts)
	now := int32(time.Now().Unix())
	helpConfig.SetDate(now)
	helpConfig.SetExpires(now + EXPIRES_TIMEOUT)
	if config.TestMode == true {
		// mtproto.NewTLBoolTrue().To_Bool()
		helpConfig.SetTestMode(mtproto.ToBool(true))
		// MakeBool(new(mtproto.TLBoolTrue))
	} else {
		helpConfig.SetTestMode(mtproto.ToBool(false))
		// MakeBool(new(mtproto.TLBoolFalse))
	}
	helpConfig.SetThisDc(config.ThisDc)
	// = config.ThisDc
	dcOptions := make([]*mtproto.DcOption, 0, len(config.DcOptions))
	for _, opt := range config.DcOptions {
		dcOption := mtproto.NewTLDcOption()
		dcOption.SetIpv6(opt.Ipv6)
		dcOption.SetMediaOnly(opt.MediaOnly)
		dcOption.SetTcpoOnly(opt.TcpoOnly)
		dcOption.SetCdn(opt.Cdn)
		dcOption.SetStatic(opt.Static)
		dcOption.SetId(opt.Id)
		dcOption.SetIpAddress(opt.IpAddress)
		dcOption.SetPort(opt.Port)
		dcOptions = append(dcOptions, dcOption.To_DcOption())
		// helpConfig.SetDcOptions = append(helpConfig.DcOptions, mtproto.MakeDcOption(dcOption))
	}
	helpConfig.SetDcOptions(dcOptions)

	helpConfig.SetChatSizeMax(config.ChatSizeMax)
	helpConfig.SetMegagroupSizeMax(config.MegagroupSizeMax)
	helpConfig.SetForwardedCountMax(config.ForwardedCountMax)
	helpConfig.SetOnlineUpdatePeriodMs(config.OnlineUpdatePeriodMs)
	helpConfig.SetOfflineBlurTimeoutMs(config.OfflineBlurTimeoutMs)
	helpConfig.SetOfflineIdleTimeoutMs(config.OfflineIdleTimeoutMs)
	helpConfig.SetOnlineCloudTimeoutMs(config.OnlineCloudTimeoutMs)
	helpConfig.SetNotifyCloudDelayMs(config.NotifyCloudDelayMs)
	helpConfig.SetNotifyDefaultDelayMs(config.NotifyDefaultDelayMs)
	helpConfig.SetChatBigSize(config.ChatBigSize)
	helpConfig.SetPushChatPeriodMs(config.PushChatPeriodMs)
	helpConfig.SetPushChatLimit(config.PushChatLimit)
	helpConfig.SetSavedGifsLimit(config.SavedGifsLimit)
	helpConfig.SetEditTimeLimit(config.EditTimeLimit)
	helpConfig.SetRatingEDecay(config.RatingEDecay)
	helpConfig.SetStickersRecentLimit(config.StickersRecentLimit)
	helpConfig.SetStickersFavedLimit(config.StickersFavedLimit)
	helpConfig.SetChannelsReadMediaPeriod(config.ChannelsReadMediaPeriod)
	helpConfig.SetTmpSessions(config.TmpSessions)
	helpConfig.SetPinnedDialogsCountMax(config.PinnedDialogsCountMax)
	helpConfig.SetCallReceiveTimeoutMs(config.CallReceiveTimeoutMs)
	helpConfig.SetCallRingTimeoutMs(config.CallRingTimeoutMs)
	helpConfig.SetCallConnectTimeoutMs(config.CallConnectTimeoutMs)
	helpConfig.SetCallPacketTimeoutMs(config.CallPacketTimeoutMs)
	helpConfig.SetMeUrlPrefix(config.MeUrlPrefix)
	helpConfig.SetSuggestedLangCode(config.SuggestedLangCode)
	helpConfig.SetLangPackVersion(config.LangPackVersion)

	disabledFeatures := make([]*mtproto.DisabledFeature, 0, len(config.DisabledFeatures))
	for _, disabled := range config.DisabledFeatures {
		disabledFeature := mtproto.NewTLDisabledFeature()
		disabledFeature.SetFeature(disabled.Feature)
		disabledFeature.SetDescription(disabled.Description)
		disabledFeatures = append(disabledFeatures, disabledFeature.To_DisabledFeature())
		// helpConfig.DisabledFeatures = append(helpConfig.DisabledFeatures, mtproto.MakeDisabledFeature(disabledFeature))
	}
	helpConfig.SetDisabledFeatures(disabledFeatures)

	reply := helpConfig.To_Config()
	glog.Infof("help.getConfig#c4f9186b - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}
