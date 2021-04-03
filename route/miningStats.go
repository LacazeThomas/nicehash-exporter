package route

import (
	"github.com/lacazethomas/nicehashExporter/config"
	"github.com/lacazethomas/nicehashExporter/models"
)

//GetBalance from Nicehash API
func GetMiningStats(cfg config.AppConfig) (m models.MiningStats, err error) {
	err = getToStructWithLogin(cfg, "/main/api/v2/mining/rigs2", &m)
	return
}
