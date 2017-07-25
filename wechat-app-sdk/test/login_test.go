package test

import (
	"github.com/yinhui87/wechat-sdk/wechat-app-sdk"
	"testing"
)

const (
	APP_ID     = ""
	APP_SECRET = ""
)

func TestCode2Session(t *testing.T) {
	code := ""
	appSdk, err := wechat_app_sdk.GetAppSdk(APP_ID, APP_SECRET)
	if err != nil {
		t.Fatalf("Get sdk error: %#v", err)
	}
	session, err := appSdk.Code2Session(code)
	if err != nil {
		t.Fatalf("Code2Session error: %#v", err)
	}
	t.Log(session)
}
