package login

import (
	config "monaco/config"
	"monaco/defs"
	"monaco/logger"
	"monaco/request"
	"net/url"
)

func Wx(queryMap url.Values) string {
	wxConf := config.Config.Wechat
	reqParams := request.Parameters{
		"appid":      wxConf.Appid,
		"secret":     wxConf.Secret,
		"grant_type": wxConf.GrantType,
		"js_code":    queryMap.Get("code"),
	}

	resp, err := request.GET(wxConf.WechatURL, reqParams)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return ""
	}
	logger.Debug(resp)
	return string(resp)
}
