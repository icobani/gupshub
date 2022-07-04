/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 30.06.2022
   Time     : 14:36
*/

package test

import (
	"gupshub"
	"os"
	"testing"
)

func TestTemplateMessageRequest(t *testing.T) {
	toPhoneNumber := "905427662390"
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	_ = wa.OptIn(toPhoneNumber)
	templateMessageRequest := wa.NewTemplateMessageRequest()
	templateMessageRequest.Source = os.Getenv("WaSourcePhoneNumber")
	templateMessageRequest.Destination = toPhoneNumber
	templateMessageRequest.Template = &gupshub.TemplateData{
		ID:     "69b83f67-0372-4a5e-b748-0ecded143304",
		Params: []string{"https://download.icibot.app"},
	}
	//	templateMessageRequest.Message = &gupshub.TemplateMessage{
	//		Type: "location",
	//		Location: &gupshub.LocationPoint{
	//			Latitude:  "41.0687831",
	//			Longitude: "29.0110758",
	//		},
	//	}
	templateMessageRequest.Send()
}

//
//func TestCurl1(t *testing.T) {
//	client := &http.Client{}
//	var data = strings.NewReader(`source=902123567077&destination=905332117087&template={"id": "69b83f67-0372-4a5e-b748-0ecded143304","params": ["https://download.icibot.app"]}`)
//	req, err := http.NewRequest("POST", "http://api.gupshup.io/sm/api/v1/template/msg", data)
//	if err != nil {
//		log.Fatal(err)
//	}
//	req.Header.Set("apikey", "2lb8bzbdbvikrv9buvhhaxb8hj65be8a")
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//	bodyText, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s\n", bodyText)
//}
