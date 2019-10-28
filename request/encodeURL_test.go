package request

import (
	"testing"
)

func TestEncodeURL(t *testing.T) {
	p := Parameters{
		"hello": "world",
	}
	URL, _ := EncodeURL("wwww.google.com", p)
	t.Log(URL)
}
