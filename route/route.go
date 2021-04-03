package route

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
	"github.com/lacazethomas/nicehash-exporter/utils"

	"github.com/pkg/errors"
)

func getToStructWithLogin(cfg config.AppConfig, endpoint string, target interface{}) (err error) {

	req, err := parsingHeader(endpoint, cfg)
	if err != nil {
		return errors.Wrap(err, "error parsing header")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error executing HTTP request")
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("error sending HTTP request code error %d", res.StatusCode))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling body to target struct")
	}

	return
}

func getToStruct(APIUrl string, target interface{}) (err error) {

	req, err := http.NewRequest(
		http.MethodGet,
		APIUrl,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "error creating HTTP request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error executing HTTP request")
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("error sending HTTP request code error %d", res.StatusCode))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling body to target struct")
	}

	return
}

func getDigest(APISecret, APIKey string, time models.Time, nonce, XOrganizationId, method, endpoint string) string {

	message := APIKey + "\x00" + fmt.Sprint(time.ServerTime) + "\x00" + nonce + "\x00" + "\x00" + XOrganizationId + "\x00" + "\x00" + method + "\x00" + endpoint + "\x00"
	h := hmac.New(sha256.New, []byte(APISecret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func parsingHeader(endpoint string, cfg config.AppConfig) (*http.Request, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		cfg.APIUrl+endpoint,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "error creating HTTP request")
	}

	time, err := GetTime()
	if err != nil {
		return nil, errors.Wrap(err, "error getting timestamp")
	}
	req.Header.Add("X-Time", fmt.Sprint(time.ServerTime))

	nonce, err := utils.GenerateSecureToken(36)
	if err != nil {
		return nil, errors.Wrap(err, "error creating nonce")
	}
	req.Header.Add("X-Nonce", nonce)

	RequestID, err := utils.GenerateSecureToken(36)
	if err != nil {
		return nil, errors.Wrap(err, "error creating RequestID")
	}
	req.Header.Add("X-Request-Id", RequestID)

	req.Header.Add("X-User-Lang", "en")
	req.Header.Add("X-Organization-Id", cfg.XOrganizationId)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth", cfg.APIKey+":"+getDigest(cfg.APISecret, cfg.APIKey, time, nonce, cfg.XOrganizationId, "GET", endpoint))

	return req, nil
}
