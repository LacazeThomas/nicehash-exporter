package route

import (
	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
)

//GetFeeRules from Nicehash API
func GetFeeRules(cfg config.AppConfig) (f models.Feerules, err error) {
	err = getToStruct(cfg, "/main/api/v2/public/service/fee/info", true, &f)
	return
}
