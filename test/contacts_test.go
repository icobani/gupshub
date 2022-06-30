/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 19:10
*/

package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gupshub"
	"os"
	"testing"
)

func TestContacts(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	if users, err := wa.Contacts(); err == nil {
		assert.NotNil(t, users)
		PrintPrettyStruct(users)
	} else {
		assert.Nil(t, users)
		fmt.Printf("Err : %+v\n", err)
	}
}

func TestOptIn(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	if err := wa.OptIn("905337691921"); err == nil {
		assert.Nil(t, err)
	} else {
		fmt.Println(err.Error())
		assert.NotNil(t, err)
	}
}

func TestOptOut(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.AppName = os.Getenv("WAAppName")
	wa.ApiKey = os.Getenv("WAApiKey")
	if err := wa.OptOut("905322551289"); err == nil {
		assert.Nil(t, err)
	} else {
		fmt.Println(err.Error())
		assert.NotNil(t, err)
	}
}

func PrintPrettyStruct(value interface{}) {
	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", string(valueJSON))
}
