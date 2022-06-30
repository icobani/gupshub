/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 19:08
*/

package gupshub

import "fmt"

type ContactListResp struct {
	Status  string    `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Users   []Contact `json:"users,omitempty"`
}

type Contact struct {
	CountryCode          string `json:"countryCode"`
	LastMessageTimeStamp int64  `json:"lastMessageTimeStamp"`
	OptinTimeStamp       int    `json:"optinTimeStamp"`
	PhoneCode            string `json:"phoneCode"`
}

// Contacts Get all user opt-ins for an app
func (wa WhatsApp) Contacts() ([]Contact, error) {
	var resp ContactListResp
	if _, err := RClient.R().
		SetResult(&resp).
		SetHeader("apikey", wa.ApiKey).
		Get(fmt.Sprintf("https://api.gupshup.io/sm/api/v1/users/%s", wa.AppName)); err == nil {
		if resp.Status == "success" {
			return resp.Users, nil
		} else {
			return nil, fmt.Errorf("%s", resp.Message)
		}
	} else {
		return nil, err
	}
}

// OptIn Mark a user as opted-in
// Opt-in, müşterilerin bir markadan iletişim almak için açık onay vermesi.
/*
curl --request POST \
     --url https://api.gupshup.io/sm/api/v1/app/opt/in/ICITECH \
     --header 'Content-Type: application/x-www-form-urlencoded' \
     --header 'apikey: 2lb8bzbdbvikrv9buvhhaxb8hj65be8a' \
     --data user=905332117087
*/
func (wa WhatsApp) OptIn(phoneNumber string) error {
	if res, err := RClient.R().
		SetHeader("apikey", wa.ApiKey).
		SetFormData(map[string]string{
			"user": phoneNumber,
		}).
		Post(fmt.Sprintf("https://api.gupshup.io/sm/api/v1/app/opt/in/%s", wa.AppName)); err == nil {
		fmt.Println(res.StatusCode(), res.Status(), res.String(), res.Error())
		return nil
	} else {
		fmt.Println(err.Error())
		return err
	}
}

// OptOut Mark a user as opted-out
// Opt-out, müşterilerin bir markadan iletişim almak için açık onayı iptal etmesi.
func (wa WhatsApp) OptOut(phoneNumber string) error {
	if _, err := RClient.R().
		SetHeader("apikey", wa.ApiKey).
		SetFormData(map[string]string{
			"user": phoneNumber,
		}).
		Post(fmt.Sprintf("https://api.gupshup.io/sm/api/v1/app/opt/out/%s", wa.AppName)); err == nil {
		return nil
	} else {
		fmt.Println(err.Error())
		return err
	}
}
