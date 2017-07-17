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
	aToken, err := wechatSdk.GetOauthAccessToken(code)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(aToken)
}

func TestGetOauthUserInfo(t *testing.T) {
	openid := ""
	accessToken := ""
	wechatSdk, err := wechatsdk.GetWechatSdk(APP_ID, APP_SECRET, APP_TOKEN)
	if err != nil {
		t.Fatal(err)
	}
	uInfo, err := wechatSdk.GetOauthUserInfo(openid, accessToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uInfo)
}
