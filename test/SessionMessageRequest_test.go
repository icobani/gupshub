/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 01.07.2022
   Time     : 16:08
*/

package test

import (
	"fmt"
	"gupshub"
	"os"
	"testing"
)

var toPhoneNumber = "905427662390"

func TestSessionMessageRequestTextMessage(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type: gupshub.MessageTypeText,
		Text: fmt.Sprintf("Merhaba %s, talebinizi aldık.", "Ali Sinan"),
	}
	smr.Send()
}

func TestSessionMessageRequestAudio(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type: gupshub.MessageTypeAudio,
		Url:  "https://www.buildquickbots.com/whatsapp/media/sample/audio/sample02.mp3",
	}
	smr.Send()
}

func TestSessionMessageRequestVideo(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type: gupshub.MessageTypeVideo,
		Url:  "https://www.buildquickbots.com/whatsapp/media/sample/video/sample01.mp4",
	}
	smr.Send()
}

func TestSessionMessageRequestFile(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:     gupshub.MessageTypeFile,
		Url:      "https://www.buildquickbots.com/whatsapp/media/sample/pdf/sample01.pdf",
		FileName: "Sample file",
	}
	smr.Send()
}

func TestSessionMessageRequestSticker(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type: gupshub.MessageTypeSticker,
		Url:  "http://www.buildquickbots.com/whatsapp/stickers/SampleSticker01.webp",
	}
	smr.Send()
}

func TestSessionMessageRequestLocation(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:      gupshub.MessageTypeLocation,
		Latitude:  "41.067500366576056",
		Longitude: "29.011861898838955",
		Name:      "ICITECH Teknoloji A.Ş.",
		Address:   "Esentepe Mahallesi Kasap sk. No:18/39 Esentepe Şişli/İstanbul",
	}
	smr.Send()
}

func TestSessionMessageRequestImage(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:        gupshub.MessageTypeImage,
		OriginalURL: "https://www.buildquickbots.com/whatsapp/media/sample/jpg/sample01.jpg",
		PreviewURL:  "https://www.buildquickbots.com/whatsapp/media/sample/jpg/sample01.jpg",
		Caption:     "Sample image",
	}
	smr.Send()
}

func TestSessionMessageRequestContact(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type: gupshub.MessageTypeContact,
		Contact: &gupshub.ContactCard{
			Name:     gupshub.NameData{"Ibrahm", "COBANI", "Ibrahim COBANI"},
			Birthday: "1975-03-02",
			Addresses: []gupshub.Address{
				{
					Type:        "WORK",
					Street:      "Esentepe Mahallesi Kasap sk. No:18/39 Esentepe Şişli",
					City:        "İstanbul",
					Country:     "Turkey",
					CountryCode: "tr",
					Zip:         "34394",
				},
			},
			Emails: []gupshub.Email{{Type: "WORK", Email: "ibrahim@imaconsult.com"}},
			Phones: []gupshub.Phone{{Type: "WORK", Phone: "+90 532 540 1194"}},
			Urls:   []gupshub.URL{{Type: "WORK", Url: "https://www.icitech.com.tr"}},
			Org:    gupshub.Organization{"ICITECH Teknoloji A.Ş.", "Development", "Manager"},
		},
	}
	smr.Send()
}

func TestSessionMessageRequestInteractiveListMessage(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:  gupshub.MessageInteractiveListMessage,
		Title: "Rezervasyon",
		Body: "Merhaba sevgili misafirimiz önümüzdeki sene için ön rezervasyon fırsatlarından yararlanmak için planladığınız " +
			"dönemi seçerek bize gönderin. Sizi ilgili dönem için ön rezervasyon fırsatları hakkında bilgilendirelim.",
		MsgID: "700",
		GlobalButtons: []gupshub.GlobalButton{
			{
				Type:  gupshub.MessageTypeText,
				Title: "Tatil Fırsatları",
			},
		},
		Items: []gupshub.ButtonItem{
			{
				Title:    "Sonbahar Dönemi",
				Subtitle: "Sonbaharda Fethiye Sizi Mutlu edecek",
				Options: []gupshub.ButtonOption{
					{
						Type:         gupshub.MessageTypeText,
						Title:        "Eylül",
						Description:  "Sonbahar kendini gösterirken göçmen kuşların göç mevsimi.",
						PostbackText: "700;rezervasyon;Eylül;2022",
					},
					{
						Type:         gupshub.MessageTypeText,
						Title:        "Ekim",
						Description:  "Yaprakların kıpkızıl olduğu yorgun mevsim.",
						PostbackText: "700;rezervasyon;Ekim;2022",
					},
					{
						Type:         gupshub.MessageTypeText,
						Title:        "Kasım",
						Description:  "Serin ama hala denizin tadını çıkartabilirsiniz.",
						PostbackText: "700;rezervasyon;Kasım;2022",
					},
				},
			},
		},
	}
	smr.Send()
}

func TestSessionMessageRequestQuickReplyText(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:  gupshub.MessageInteractiveQuickReplyMessage,
		MsgID: "700",
		Content: gupshub.Content{
			Type:    gupshub.MessageTypeText,
			Header:  "Clup Tuana Spa & Wellness",
			Caption: "Lütfen seçiminizi yapınız.",
			Text:    "Tuana Spa'da bulunan birbirinden farklı masajlar ile kış yorgunluğunu üzerinizden atabilir, cilt bakımı ve süt banyosu ile yenilenip, sauna ve hamama girerek gönlünüzce rahatlayabilirsiniz. Restoranlarımızdaki lezzetli yemeklerin ardından formunuzu korumak için fitness salonuna uğramayı da ihmal etmeyin.",
		},
		Options: []gupshub.Option{
			{
				Type:         gupshub.MessageTypeText,
				Title:        "70€ Günlük",
				PostbackText: "700;spa;daily",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "300€ haftalık",
				PostbackText: "700;spa;weekly",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "İlgilenmiyorum",
				PostbackText: "700;spa;not_interested",
			},
		},
	}
	smr.Send()
}

func TestSessionMessageRequestQuickReplyImage(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:  gupshub.MessageInteractiveQuickReplyMessage,
		MsgID: "700",
		Content: gupshub.Content{
			Type:    gupshub.MessageTypeImage,
			Url:     "https://b1development.s3.eu-central-1.amazonaws.com/icibotV2/700/Items/26080_3.jpeg",
			Caption: "Lütfen seçiminizi yapınız.",
			Text:    "Çocukları eğlenceli aktiviteler bekliyor.\n\nUzun zamandır bitiremediğiniz en sevdiğiniz kitabı okuyabilir, kumsalda hafif meltemde şekerleme yapabilir veya denizde serinleyebilir, çocuğunuz becerilerini keşfedip yaşıtlarıyla eğlenirken siz de kendinizi şımartabilirsiniz. Tüm bu aktivitelerle birlikte akşamları daha önce hiç duymadığınız bir şarkıyı dinlemeye veya yepyeni bir fıkra ile gülmeye hazır olmalısınız. Yine de çocuğunuzun sizi nasıl şaşırtacağını asla bilemezsiniz.",
		},
		Options: []gupshub.Option{
			{
				Type:         gupshub.MessageTypeText,
				Title:        "10€ Günlük",
				PostbackText: "700;spa;daily",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "700€ haftalık",
				PostbackText: "700;spa;weekly",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "İlgilenmiyorum",
				PostbackText: "700;spa;not_interested",
			},
		},
	}
	smr.Send()
}

func TestSessionMessageRequestQuickReplyVideo(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:  gupshub.MessageInteractiveQuickReplyMessage,
		MsgID: "700",
		Content: gupshub.Content{
			Type:    gupshub.MessageTypeVideo,
			Url:     "https://media.istockphoto.com/videos/girl-twirling-with-dandelion-seeds-video-id493319011",
			Caption: "Lütfen seçiminizi yapınız.",
			Text:    "Çocukları eğlenceli aktiviteler bekliyor.\n\nUzun zamandır bitiremediğiniz en sevdiğiniz kitabı okuyabilir, kumsalda hafif meltemde şekerleme yapabilir veya denizde serinleyebilir, çocuğunuz becerilerini keşfedip yaşıtlarıyla eğlenirken siz de kendinizi şımartabilirsiniz. Tüm bu aktivitelerle birlikte akşamları daha önce hiç duymadığınız bir şarkıyı dinlemeye veya yepyeni bir fıkra ile gülmeye hazır olmalısınız. Yine de çocuğunuzun sizi nasıl şaşırtacağını asla bilemezsiniz.",
		},
		Options: []gupshub.Option{
			{
				Type:         gupshub.MessageTypeText,
				Title:        "10€ Günlük",
				PostbackText: "700;spa;daily",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "700€ haftalık",
				PostbackText: "700;spa;weekly",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "İlgilenmiyorum",
				PostbackText: "700;spa;not_interested",
			},
		},
	}
	smr.Send()
}

func TestSessionMessageRequestQuickReplyFile(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	smr := wa.NewSessionMessageRequest()
	smr.Source = os.Getenv("WaSourcePhoneNumber")
	smr.Destination = toPhoneNumber
	smr.Message = &gupshub.SessionMessageData{
		Type:  gupshub.MessageInteractiveQuickReplyMessage,
		MsgID: "700",
		Content: gupshub.Content{
			Type:     gupshub.MessageTypeFile,
			Url:      "https://b1development.s3.eu-central-1.amazonaws.com/icibotV2/273/GASTRO%20JAZZ%20PDF.pdf",
			FileName: "GASTRO JAZZ Menu.pdf",
			Caption:  "Hemen yerinizi ayırtın.",
			Text:     "Menüye Gözat.",
		},
		Options: []gupshub.Option{
			{
				Type:         gupshub.MessageTypeText,
				Title:        "20€ Bugün",
				PostbackText: "700;spa;daily",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "20€ yarın",
				PostbackText: "700;spa;weekly",
			},
			{
				Type:         gupshub.MessageTypeText,
				Title:        "35€ Haftasonu",
				PostbackText: "700;spa;not_interested",
			},
		},
	}
	smr.Send()
}
