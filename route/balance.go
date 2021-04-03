package route

import (
	"github.com/lacazethomas/nicehashExporter/config"
	"github.com/lacazethomas/nicehashExporter/models"
)

//GetBalance from Nicehash API
func GetBalance(cfg config.AppConfig) (b models.Balance, err error) {
	err = getToStructWithLogin(cfg, "/main/api/v2/accounting/account2/BTC", &b)
	return
}
