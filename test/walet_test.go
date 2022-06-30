/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 20:31
WAAppName
WAApiKey
*/

package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gupshub"
	"os"
	"testing"
)

func TestWalletBalance(t *testing.T) {
	var wa gupshub.WhatsApp
	wa.ApiKey = os.Getenv("WAApiKey")
	if balance, err := wa.Balance(); err == nil {
		assert.NotNil(t, balance)
		fmt.Printf("Balance : %+v\n", balance)
	} else {
		assert.Zero(t, balance)
		fmt.Printf("Err : %+v\n", err.Error())
	}
}
