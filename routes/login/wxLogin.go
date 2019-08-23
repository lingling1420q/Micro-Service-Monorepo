package login

import (
	config "gin-demo/config"
	"gin-demo/defs"
	"gin-demo/logger"
	"gin-demo/utils"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Wx(c *gin.Context) {
	wxConf := config.Config().Wechat
	queryMap := c.Request.URL.Query()
	reqParams := url.Values{
		"appid":      {wxConf.Appid},
		"secret":     {wxConf.Secret},
		"grant_type": {wxConf.GrantType},
		"js_code":    {queryMap.Get("code")},
	}

	url, err := utils.EncodeURL(wxConf.WechatURL, reqParams)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return
	}
	logger.Notice(url)

	resp, err := utils.HTTPGet(url)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return
	}
	logger.Debug(resp)
	c.String(http.StatusOK, resp)
}
