package agolet

import (
	"encoding/json"
	"fmt"
	"monaco/request"
)

func GetTocken() string {
	var tocken struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	p := request.Parameters{
		"corpid":     "ww419d29f97fe0161d",
		"corpsecret": "uvpuo_GAHxEhKalDJykR30XefH2Vo-9ZrQQ9uWx2TFw",
	}
	tockenBytes, _ := request.GET("https://qyapi.weixin.qq.com/cgi-bin/gettoken", p)

	if err := json.Unmarshal(tockenBytes, &tocken); err != nil {
		panic(err.Error())
	}
	return tocken.AccessToken
}

func PostMsg() {
	type textStruct struct {
		Content string `json:"content"`
	}
	type msgStruct struct {
		ToUser string `json:"touser"`
		// ToTag string `json:"totag"`
		MsgType string     `json:"msgtype"`
		AgentID int        `json:"agentid"`
		Text    textStruct `json:"text"`
	}
	text := textStruct{
		Content: "hello world !",
	}
	msg := &msgStruct{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: 1000002,
		Text:    text,
	}
	bytesJSONMsg, _ := json.Marshal(msg)
	tocken := GetTocken()
	p := request.Parameters{
		"access_token": tocken,
	}
	r, _ := request.POST("https://qyapi.weixin.qq.com/cgi-bin/message/send", p, bytesJSONMsg)
	fmt.Println(string(r))
}
