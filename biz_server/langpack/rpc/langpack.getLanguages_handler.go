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
)



// languages config
/**************************

[[LangPackLanguages]]
Name = "English"
NativeName = "English"
LangCode = "en"

[[LangPackLanguages]]
Name = "German"
NativeName = "Deutsch"
LangCode = "de"

[[LangPackLanguages]]
Name = "Dutch"
NativeName = "Nederlands"
LangCode = "nl"

[[LangPackLanguages]]
Name = "Spanish"
NativeName = "Español"
LangCode = "es"

[[LangPackLanguages]]
Name = "Italian"
NativeName = "Italiano"
LangCode = "it"

[[LangPackLanguages]]
Name = "Portuguese (Brazil)"
NativeName = "Português (Brasil)"
LangCode = "pt-br"

[[LangPackLanguages]]
Name = "Korean"
NativeName = "한국어"
LangCode = "ko"

[[LangPackLanguages]]
Name = "Malay"
NativeName = "Bahasa Melayu"
LangCode = "ms"

[[LangPackLanguages]]
Name = "Russian"
NativeName = "Русский"
LangCode = "ru"

[[LangPackLanguages]]
Name = "French"
NativeName = "Français"
LangCode = "fr"

[[LangPackLanguages]]
Name = "Ukrainian"
NativeName = "Українська"
LangCode = "uk"

**************************************************/

// langpack.getLanguages#800fd57d = Vector<LangPackLanguage>;
func (s *LangpackServiceImpl) LangpackGetLanguages(ctx context.Context, request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("langpack.getLanguages#800fd57d - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Add other language
	language := &mtproto.TLLangPackLanguage{Data2: &mtproto.LangPackLanguage_Data{
		Name:       "English",
		NativeName: "English",
		LangCode:   "en",
	}}

	languages := &mtproto.Vector_LangPackLanguage{}
	languages.Datas = append(languages.Datas, language.To_LangPackLanguage())

	glog.Infof("langpack.getLanguages#800fd57d - reply: %s", logger.JsonDebugData(languages))
	return languages, nil
}
