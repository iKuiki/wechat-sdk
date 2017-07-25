package wechat_app_sdk

import (
	"errors"
	"github.com/astaxie/beego/httplib"
)

func (this *AppSdk) Code2Session(code string) (session AppSessionKey, err error) {
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	req.Param("appid", this.AppId)
	req.Param("secret", this.AppSecret)
	req.Param("js_code", code)
	req.Param("keygrant_type", "authorization_code")
	err = req.ToJSON(&session)

	if err != nil {
		return AppSessionKey{}, errors.New("request error: " + err.Error())
	}
	if session.WechatError != nil && session.Errcode != 0 {
		return AppSessionKey{}, errors.New(session.Errmsg)
	}
	return session, nil
}
