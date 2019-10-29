package login

import (
	"encoding/json"
	config "monaco/config"
	"monaco/defs"
	"monaco/logger"
	"monaco/request"
	"net/url"
)

//GetTokenRes get tokrn response
type GetTokenRes struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//Qywx user login
func Qywx(queryMap url.Values) string {
	qywxConf := config.Config.QyWechat

	//企业内部开发 获取token
	//https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ID&corpsecret=SECRECT
	reqTokenParam := request.Parameters{
		"corpid":     qywxConf.Corpid,
		"corpsecret": qywxConf.Secret,
	}
	resp, err := request.GET(qywxConf.GetTokenURL, reqTokenParam)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return ""
	}
	getTokenRes := GetTokenRes{}
	json.Unmarshal([]byte(resp), &getTokenRes)

	token := getTokenRes.AccessToken
	logger.Debug(token)

	// https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=ACCESS_TOKEN&js_code=CODE&grant_type=authorization_code
	reqParams := request.Parameters{
		"access_token": token,
		"grant_type":   qywxConf.GrantType,
		"js_code":      queryMap.Get("code"),
	}

	resp, err = request.GET(qywxConf.WechatURL, reqParams)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return ""
	}
	logger.Debug(resp)
	return string(resp)
}
