package route

import "github.com/lacazethomas/nicehash-exporter/models"

var APIURL = "https://api2.nicehash.com/api/v2/time"

//GetTime from Nicehash API
func GetTime() (t models.Time, err error) {
	err = getToStruct(APIURL, &t)
	return
}
