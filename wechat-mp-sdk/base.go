package wechat_mp_sdk

import (
	"github.com/yinhui87/go-component/language"
	"strings"
)

func (this *MpSdk) CheckSignature(signature, timestamp, nonce string) (pass bool) {
	arrInfo := []string{this.Token, timestamp, nonce}
	arrInfo = language.ArraySort(arrInfo).([]string)
	arrInfoStr := strings.Join(arrInfo, "")
	curSignature := this.sha1(arrInfoStr)
	return curSignature == signature
}
