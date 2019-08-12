package login

import (
	"encoding/json"
	config "gin-demo/config"
	"gin-demo/defs"
	"gin-demo/logger"
	"gin-demo/utils"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

//Qywx user login
func Qywx(c *gin.Context) {
	qywxConf := config.Config().QyWechat
	queryMap := c.Request.URL.Query()

	//企业内部开发 获取token
	//https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ID&corpsecret=SECRECT
	reqTokenParam := url.Values{
		"corpid":     {qywxConf.Corpid},
		"corpsecret": {qywxConf.Secret},
	}

	urlGetToken, err := utils.EncodeURL(qywxConf.GetTokenURL, reqTokenParam)
	resp, err := utils.HTTPGet(urlGetToken)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return
	}
	getTokenRes := GetTokenRes{}
	json.Unmarshal([]byte(resp), &getTokenRes)

	token := getTokenRes.AccessToken
	logger.Debug(token)

	// https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=ACCESS_TOKEN&js_code=CODE&grant_type=authorization_code

	reqParams := url.Values{
		"access_token": {token},
		"grant_type":   {qywxConf.GrantType},
		"js_code":      {queryMap.Get("code")},
	}

	url, err := utils.EncodeURL(qywxConf.WechatURL, reqParams)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return
	}
	logger.Notice(url)

	resp, err = utils.HTTPGet(url)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return
	}

	logger.Debug(resp)
	c.String(http.StatusOK, resp)
}
