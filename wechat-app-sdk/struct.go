package wechat_app_sdk

import (
	"github.com/yinhui87/wechat-sdk"
)

type AppSessionKey struct {
	*wechatsdk.WechatError
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

type UserInfo struct {
	AvatarURL string `json:"avatarUrl"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    uint8  `json:"gender"`
	Language  string `json:"language"`
	NickName  string `json:"nickName"`
	OpenID    string `json:"openId"`
	Province  string `json:"province"`
	UnionID   string `json:"unionId"`
	Watermark struct {
		Appid     string `json:"appid"`
		Timestamp uint32 `json:"timestamp"`
	} `json:"watermark"`
}
