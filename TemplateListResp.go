/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 20:05
*/

package gupshub

import "fmt"

type TemplateListResp struct {
	Status    string        `json:"status,omitempty"`
	Message   string        `json:"reason,omitempty"`
	Templates []TemplateRec `json:"templates,omitempty"`
}

type TemplateRec struct {
	AppId          string `json:"appId"`
	Category       string `json:"category"`
	CreatedOn      int64  `json:"createdOn"`
	Data           string `json:"data"`
	ElementName    string `json:"elementName"`
	ExternalId     string `json:"externalId"`
	Id             string `json:"id"`
	LanguageCode   string `json:"languageCode"`
	LanguagePolicy string `json:"languagePolicy"`
	Master         bool   `json:"master"`
	Meta           string `json:"meta"`
	ModifiedOn     int64  `json:"modifiedOn"`
	Namespace      string `json:"namespace"`
	Status         string `json:"status"`
	TemplateType   string `json:"templateType"`
	Vertical       string `json:"vertical"`
}

// GetTemplateList Get all templates for an app
func (wa WhatsApp) GetTemplateList() ([]TemplateRec, error) {
	var resp TemplateListResp
	if _, err := RClient.R().
		SetResult(&resp).
		SetError(&resp).
		SetHeader("apikey", wa.ApiKey).
		Get(fmt.Sprintf("https://api.gupshup.io/sm/api/v1/template/list/%s", wa.AppName)); err == nil {
		if resp.Status == "success" {
			return resp.Templates, nil
		} else {
			return nil, fmt.Errorf("%s", resp.Message)
		}
	} else {
		return nil, err
	}
}
