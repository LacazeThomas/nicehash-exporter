package route

import (
	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
)

//GetBalance from Nicehash API
func GetMiningStats(cfg config.AppConfig) (m models.MiningStats, err error) {
	err = getToStructWithLogin(cfg, "/main/api/v2/mining/rigs2", &m)
	return
}
