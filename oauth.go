package wechatsdk

import (
	"errors"
	"github.com/astaxie/beego/httplib"
)

func (this *WechatSdk) GetAccessToken(code string) (accessToken WechatOauthToken, err error) {
	req := httplib.Get("https://api.weixin.qq.com/sns/oauth2/access_token")
	req.Param("appid", this.AppId)
	req.Param("secret", this.AppSecret)
	req.Param("code", code)
	req.Param("grant_type", "authorization_code")
	err = req.ToJSON(&accessToken)
	if err != nil {
		return WechatOauthToken{}, errors.New("request error: " + err.Error())
	}
	if accessToken.Errcode != 0 {
		return WechatOauthToken{}, errors.New(accessToken.Errmsg)
	}
	return accessToken, nil
}
