/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 20:18
*/

package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gupshub"
	"os"
	"testing"
)

func TestTemplates(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	if templates, err := wa.GetTemplateList(); err == nil {
		assert.NotNil(t, templates)
		PrintPrettyStruct(templates)
	} else {
		assert.Nil(t, templates)
		fmt.Printf("Err : %+v\n", err)
	}
}
