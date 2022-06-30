/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Contact     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 29.06.2022
   Time     : 20:29
*/

package gupshub

import "fmt"

type WalletBalanceResp struct {
	Status  string  `json:"status,omitempty"`
	Message string  `json:"message,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}

// Balance Check the wallet balance for an account
func (wa WhatsApp) Balance() (float64, error) {
	var resp WalletBalanceResp
	if _, err := RClient.R().
		SetResult(&resp).
		SetError(&resp).
		SetHeader("apikey", wa.ApiKey).
		Get("https://api.gupshup.io/sm/api/v2/wallet/balance"); err == nil {
		if resp.Status == "success" {
			return resp.Balance, nil
		} else {
			return 0, fmt.Errorf("%s", resp.Message)
		}
	} else {
		return 0, err
	}
}
