package test

import (
	"github.com/yinhui87/wechat-sdk/wechat-app-sdk"
	"testing"
)

const (
	APP_ID     = ""
	APP_SECRET = ""
)

var (
	sdk *wechat_app_sdk.AppSdk
)

func init() {
	var err error
	sdk, err = wechat_app_sdk.GetAppSdk(APP_ID, APP_SECRET)
	if err != nil {
		panic(err)
	}
}
func TestCode2Session(t *testing.T) {
	code := ""
	session, err := sdk.Code2Session(code)
	if err != nil {
		t.Fatalf("Code2Session error: %#v\n", err)
	}
	t.Log(session)
}

func TestDecryptUserInfo(t *testing.T) {
	sessionKey := ""
	encryptData := ""
	iv := ""
	info, err := sdk.DecryptUserInfo(sessionKey, encryptData, iv)
	if err != nil {
		t.Fatalf("Decrypt User Info error: %#v\n", err)
	}
	t.Logf("%#v", info)
}
