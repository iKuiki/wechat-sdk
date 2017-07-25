package wechat_app_sdk

import (
	"github.com/yinhui87/wechat-sdk"
)

type AppSessionKey struct {
	*wechatsdk.WechatError
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
}
