/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 01.07.2022
   Time     : 14:56
*/

package gupshub

import (
	"encoding/json"
	"fmt"
)

type SessionMessageRequest struct {
	wa          *WhatsApp
	Source      string              `json:"source"`
	Destination string              `json:"destination"`
	SourceName  string              `json:"src.name"`
	Channel     string              `json:"channel"`
	Message     *SessionMessageData `json:"message,omitempty"`
}

func (s SessionMessageRequest) ToString() string {
	if b, err := json.Marshal(s); err != nil {
		return ""
	} else {
		return string(b)
	}
}

type SessionMessageData struct {
	Type          string         `json:"type,omitempty"`
	Text          string         `json:"text,omitempty"`
	Url           string         `json:"url,omitempty"`
	FileName      string         `json:"file_name,omitempty"`
	Longitude     string         `json:"longitude,omitempty"`
	Latitude      string         `json:"latitude,omitempty"`
	Name          string         `json:"name,omitempty"`
	Address       string         `json:"address,omitempty"`
	OriginalURL   string         `json:"originalUrl,omitempty"`
	PreviewURL    string         `json:"previewUrl,omitempty"`
	Caption       string         `json:"caption,omitempty"`
	Contact       *ContactCard   `json:"contact,omitempty"`
	Title         string         `json:"title,omitempty"`
	Body          string         `json:"body,omitempty"`
	MsgID         string         `json:"msgid,omitempty"`
	GlobalButtons []GlobalButton `json:"globalButtons,omitempty"`
	Items         []ButtonItem   `json:"items,omitempty"`
	Content       Content        `json:"content"`
	Options       []Option       `json:"options"`
}

type Content struct {
	Type     string `json:"type,omitempty"`
	Header   string `json:"header,omitempty"`
	Text     string `json:"text,omitempty"`
	Url      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

type Option struct {
	Type         string `json:"type,omitempty"`
	Title        string `json:"title,omitempty"`
	PostbackText string `json:"postbackText,omitempty"`
}

type GlobalButton struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
}

type ButtonOption struct {
	Type         string `json:"type,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	PostbackText string `json:"postbackText,omitempty"`
}

type ButtonItem struct {
	Title    string         `json:"title,omitempty"`
	Subtitle string         `json:"subtitle,omitempty"`
	Options  []ButtonOption `json:"options,omitempty"`
}

type Address struct {
	Type        string `json:"type,omitempty"`
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	State       string `json:"state,omitempty"`
	Zip         string `json:"zip,omitempty"`
}

type Email struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`
}

type NameData struct {
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	FormattedName string `json:"formattedName,omitempty"`
}

type Organization struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type Phone struct {
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty"`
	WaId  string `json:"wa_id,omitempty"`
}

type URL struct {
	Url  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type ContactCard struct {
	Addresses []Address    `json:"addresses,omitempty"`
	Birthday  string       `json:"birthday,omitempty"`
	Emails    []Email      `json:"emails,omitempty"`
	Name      NameData     `json:"name,omitempty"`
	Org       Organization `json:"org,omitempty"`
	Phones    []Phone      `json:"phones,omitempty"`
	Urls      []URL        `json:"urls,omitempty"`
}

func (smd SessionMessageData) ToString() string {
	if b, err := json.Marshal(smd); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (s *SessionMessageRequest) Send() {
	fmt.Println(s.ToString())
	RClient.SetDebug(true)
	RClient.SetContentLength(true)

	if res, err := RClient.R().
		SetHeader("apikey", s.wa.ApiKey).
		SetHeader("Cache-Control", "no-cache").
		SetFormData(s.convertToMapStringString()).
		Post("https://api.gupshup.io/sm/api/v1/msg"); err == nil {
		fmt.Println("ok", res.StatusCode(), res.Status(), res.String())
		fmt.Println(res.Request.FormData)
	} else {
		fmt.Println("err : ", err)
	}
}

func (s *SessionMessageRequest) convertToMapStringString() map[string]string {
	form := map[string]string{}
	form["source"] = s.Source
	form["destination"] = s.Destination
	form["src.name"] = s.SourceName
	form["channel"] = s.Channel
	if s.Message != nil {
		form["message"] = s.Message.ToString()
	}

	return form
}
