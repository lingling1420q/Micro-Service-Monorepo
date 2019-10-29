package login

import (
	"context"
	"encoding/json"
	"fmt"
	"monaco/logger"
	"monaco/request"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func GetQR(c *gin.Context) {
	// 获取access token
	resp := getAccessToken()
	logger.Debug(resp)
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(resp), &m)
	if err != nil {
		logger.Error(err)
	}
	accessToken := m["access_token"].(string)
	// 获取二维码
	// POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
	// 构造请求参数
	// 坑！！！ 不要把access_token 放到body里面
	postMap := request.Parameters{
		"scene": "a=1",
	}
	//解析参数为JSON
	paramJSON, err := json.Marshal(postMap)
	if err != nil {
		return
	}
	logger.Notice("paramJSON: ", string(paramJSON))
	// 获取小程序二维码
	qrResp, err := request.POST("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="+accessToken, request.Parameters{}, paramJSON)
	if err != nil {
		logger.Error("get qr error")
	}
	logger.Debug("上传二维码")
	UploadQR(string(qrResp))
	c.String(http.StatusOK, string(qrResp))
}

func getAccessToken() string {
	// 获取accessToken
	// GET https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	// 构造参数
	reqTokenParam := request.Parameters{
		"grant_type": "client_credential",
		"appid":      "wx4bd3f0044643ec2c",
		"secret":     "ae33804c97317d0cfc66239027d106ca",
	}

	resp, _ := request.GET("https://api.weixin.qq.com/cgi-bin/token", reqTokenParam)
	fmt.Println(resp)
	return string(resp)
}

// 获取cos client
func getCosClient() *cos.Client {
	u, _ := url.Parse("https://erp-1257983906.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDe4EAlRu1xf7JF27z5vsa2FCuNccSnAgz",
			SecretKey: "DBuN4JI1XorfGj03qRo0xiR0QccR6Zpq",
		},
	})
	return client
}

// UploadQR 将二维码存上传到cos（腾讯云对象存储）
func UploadQR(qrResp string) {
	c := getCosClient()
	f := strings.NewReader(qrResp)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "image/jpeg",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			XCosACL: "public-read",
		},
	}
	_, err := c.Object.Put(context.Background(), "test/QR.png", f, opt)
	if err != nil {
		panic(err)
	}
}
