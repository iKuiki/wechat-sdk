package wechat_mp_sdk

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

type MpSdk struct {
	AppId     string
	AppSecret string
	Token     string
}

func GetMpSdk(appId, appSecret, appToken string) (wechatSdk *MpSdk, err error) {
	if appId == "" {
		return nil, errors.New("appId empty")
	}
	return &MpSdk{
		AppId:     appId,
		AppSecret: appSecret,
		Token:     appToken,
	}, nil
}

func (this *MpSdk) sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
