package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log(Config.DB)
	t.Log(Config.DB.Mysql)
	t.Log(Config.DB.Redis)
	t.Log(Config.Server)
	t.Log(Config.Server.Debug)
	t.Log(Config.Server.Port)
	t.Log(Config.Server.LogPath)
	t.Log(Config.Wechat)
	t.Log(Config.Wechat.WechatURL)
	t.Log(Config.Wechat.Appid)
	t.Log(Config.Wechat.Secret)
	t.Log(Config.Wechat.GrantType)
	t.Log(Config.QyWechat)
}
