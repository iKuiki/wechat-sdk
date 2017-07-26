package wechat_app_sdk

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/httplib"
	"github.com/xlstudio/wxbizdatacrypt"
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

func (this *AppSdk) DecryptUserInfo(sessionKey, encryptedData, iv string) (userInfo UserInfo, err error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: this.AppId, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptedData, iv, true)
	if err != nil {
		return UserInfo{}, errors.New("Decrypt error: " + err.Error())
	} else {
		str := result.(string)
		err = json.Unmarshal([]byte(str), &userInfo)
		if err != nil {
			return UserInfo{}, errors.New("json.Unmarshal error: " + err.Error())
		}
		return userInfo, nil
	}
}
