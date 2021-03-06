/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 30.06.2022
   Time     : 13:57
*/

package gupshub

import (
	"encoding/json"
	"fmt"
)

type TemplateMessageRequest struct {
	wa          *WhatsApp
	Source      string           `json:"source"`
	Destination string           `json:"destination"`
	Template    *TemplateData    `json:"template,omitempty"`
	Message     *TemplateMessage `json:",omitempty"`
}

type TemplateData struct {
	ID     string   `json:"id"`
	Params []string `json:"params"`
}

func (t TemplateData) ToString() string {
	if b, err := json.Marshal(t); err != nil {
		return ""
	} else {
		return string(b)
	}
}

type TemplateMessage struct {
	Type     string         `json:"type,omitempty"`
	Image    *Link          `json:"image,omitempty"`
	Video    *Link          `json:"video,omitempty"`
	Document *Link          `json:"document,omitempty"`
	Location *LocationPoint `json:"location,omitempty"`
}

func (t TemplateMessage) ToString() string {
	if b, err := json.Marshal(t); err != nil {
		return ""
	} else {
		return string(b)
	}
}

type Link struct {
	Link     string `json:"link,omitempty"`
	FileName string `json:"filename,omitempty"`
}

type LocationPoint struct {
	Longitude string `json:"longitude,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
}

type TemplateMessageResponse struct {
	MessageId string `json:"messageId"`
	Status    string `json:"status"`
}

func (tmr TemplateMessageRequest) Send() (TemplateMessageResponse, error) {
	form := tmr.convertToMapStringString()
	fmt.Println(form)
	RClient.SetDebug(true)
	RClient.SetContentLength(true)

	var result TemplateMessageResponse

	if res, err := RClient.R().
		SetHeader("apikey", tmr.wa.ApiKey).
		SetHeader("Cache-Control", "no-cache").
		SetResult(result).
		SetFormData(form).
		Post("http://api.gupshup.io/sm/api/v1/template/msg"); err == nil {
		fmt.Println("ok", res.StatusCode(), res.Status(), res.String())
		fmt.Println(res.Request.FormData)
		return result, nil
	} else {
		fmt.Println("err : ", err)
		return result, err
	}
}

func (tmr TemplateMessageRequest) convertToMapStringString() map[string]string {
	form := map[string]string{}

	form["source"] = tmr.Source
	form["destination"] = tmr.Destination
	if tmr.Template != nil {
		//m := structs.Map(tmr.Template)
		form["template"] = tmr.Template.ToString()
		if tmr.Message != nil {
			form["message"] = tmr.Message.ToString()
		}
	}
	return form
}
