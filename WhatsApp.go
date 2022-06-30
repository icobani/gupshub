/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 18:51
*/

package gupshub

import (
	"github.com/go-resty/resty/v2"
	"time"
)

var (
	RClient = resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(3 * time.Second).
		SetTimeout(1 * time.Minute).
		SetContentLength(true)
)

type WhatsApp struct {
	AppName string `json:"app_name"`
	ApiKey  string `json:"api_key"`
}

func (wa *WhatsApp) NewTemplateMessageRequest() TemplateMessageRequest {
	return TemplateMessageRequest{
		wa: wa,
	}
}
