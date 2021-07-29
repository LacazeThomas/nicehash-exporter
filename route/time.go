package route

import (
	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
)

//GetTime from Nicehash API
func GetTime() (t models.Time, err error) {
	err = getToStruct(config.GetAppConfig(), "/api/v2/time", false, &t)
	return
}
