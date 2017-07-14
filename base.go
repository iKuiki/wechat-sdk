package wechatsdk

import (
	"github.com/yinhui87/go-component/language"
	"strings"
)

func (this *WechatSdk) CheckSignature(signature, timestamp, nonce string) (pass bool) {
	arrInfo := []string{this.Token, timestamp, nonce}
	arrInfo = language.ArraySort(arrInfo).([]string)
	arrInfoStr := strings.Join(arrInfo, "")
	curSignature := this.sha1(arrInfoStr)
	return curSignature == signature
}
