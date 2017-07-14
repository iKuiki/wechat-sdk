package test

import (
	"github.com/yinhui87/wechat-sdk"
	"testing"
)

const (
	APP_ID     = ""
	APP_SECRET = ""
	APP_TOKEN  = ""
)

func TestGetOauthAccessToken(t *testing.T) {
	code := ""
	wechatSdk, err := wechatsdk.GetWechatSdk(APP_ID, APP_SECRET, APP_TOKEN)
	if err != nil {
		t.Fatal(err)
	}
	aToken, err := wechatSdk.GetAccessToken(code)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(aToken)
}
