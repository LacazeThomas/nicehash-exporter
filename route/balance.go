package route

import (
	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
)

//GetBalance from Nicehash API
func GetBalance(cfg config.AppConfig) (b models.Balance, err error) {
	err = getToStruct(cfg, "/main/api/v2/accounting/account2/BTC", true, &b)
	return
}
