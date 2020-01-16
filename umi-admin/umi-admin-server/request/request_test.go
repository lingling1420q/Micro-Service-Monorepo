package request

import (
	"testing"
)

func TestGET(t *testing.T) {
	baseURL := "http://localhost:9099"
	p := Parameters{}
	header := Parameters{"token": "f5e970d92bfce00875fc26a00a3b020dfabf8728"}
	userInfo, err := GET(baseURL, p, header)
	t.Log(userInfo)
	if err != nil {
		t.Error(err)
	}
}
