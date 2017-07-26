package wechat_app_sdk

import (
	"errors"
)

type AppSdk struct {
	AppId     string
	AppSecret string
}

func GetAppSdk(appId, appSecret string) (wechatSdk *AppSdk, err error) {
	if appId == "" {
		return nil, errors.New("appId empty")
	}
	return &AppSdk{
		AppId:     appId,
		AppSecret: appSecret,
	}, nil
}
