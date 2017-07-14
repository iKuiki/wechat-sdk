package wechatsdk

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

type WechatSdk struct {
	AppId     string
	AppSecret string
	Token     string
}

func GetWechatSdk(appId, appSecret, appToken string) (wechatSdk *WechatSdk, err error) {
	if appId == "" {
		return nil, errors.New("appId empty")
	}
	return &WechatSdk{
		AppId:     appId,
		AppSecret: appSecret,
		Token:     appToken,
	}, nil
}

func (this *WechatSdk) sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
