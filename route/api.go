package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
)

func HandleMining(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetAppConfig()

	values := models.Status2{Action: r.URL.Query().Get("action"), RigID: r.URL.Query().Get("rigId")}
	json_body, err := json.Marshal(values)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, statusCode, err := postWithLogin(cfg, "/main/api/v2/mining/rigs/status2", json_body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(result)
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetAppConfig()

	rigId := r.URL.Query().Get("rigId")

	rigDetail := models.RigDetails{}
	var result string

	err := getToStruct(cfg, "/main/api/v2/mining/rig2/"+rigId, true, &rigDetail)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)

	if rigDetail.Minerstatus == "MINING" {
		result = "1"
	} else {
		result = "0"
	}

	w.Write([]byte(result))
}
