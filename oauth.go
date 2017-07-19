package wechatsdk

import (
	"errors"
	"github.com/astaxie/beego/httplib"
)

func (this *WechatSdk) GetOauthAccessToken(code string) (accessToken WechatOauthToken, err error) {
	req := httplib.Get("https://api.weixin.qq.com/sns/oauth2/access_token")
	req.Param("appid", this.AppId)
	req.Param("secret", this.AppSecret)
	req.Param("code", code)
	req.Param("grant_type", "authorization_code")
	err = req.ToJSON(&accessToken)
	if err != nil {
		return WechatOauthToken{}, errors.New("request error: " + err.Error())
	}
	if accessToken.WechatError != nil && accessToken.Errcode != 0 {
		return WechatOauthToken{}, errors.New(accessToken.Errmsg)
	}
	return accessToken, nil
}

func (this *WechatSdk) GetOauthUserInfo(openid, accessToken string) (userInfo WechatOauthUserInfo, err error) {
	req := httplib.Get("https://api.weixin.qq.com/sns/userinfo")
	req.Param("access_token", accessToken)
	req.Param("openid", openid)
	req.Param("lang", "zh_CN")
	err = req.ToJSON(&userInfo)
	if err != nil {
		return WechatOauthUserInfo{}, errors.New("request error: " + err.Error())
	}
	if userInfo.WechatError != nil && userInfo.Errcode != 0 {
		return WechatOauthUserInfo{}, errors.New(userInfo.Errmsg)
	}
	return userInfo, nil
}
