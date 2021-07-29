package route

import (
	"bytes"
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

func postWithLogin(cfg config.AppConfig, endpoint string, body []byte) ([]byte, int, error) {

	req, err := http.NewRequest(
		"POST",
		"https://api2.nicehash.com/main/api/v2/mining/rigs/status2",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, 400, errors.Wrap(err, "error creating new request")
	}

	err = parsingHeader(req, endpoint, cfg, string(body))
	if err != nil {
		return nil, 400, errors.Wrap(err, "error parsing header")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 400, err
	}
	defer resp.Body.Close()

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 400, err
	}

	return f, resp.StatusCode, nil
}

func getToStruct(cfg config.AppConfig, endpoint string, loginHeader bool, target interface{}) (err error) {

	req, err := http.NewRequest(
		http.MethodGet,
		cfg.APIUrl+endpoint,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "error creating HTTP request")
	}

	if loginHeader {
		err = parsingHeader(req, endpoint, cfg, "")
		if err != nil {
			return errors.Wrap(err, "error parsing header")
		}
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

func getDigest(APISecret, APIKey string, time models.Time, nonce, XOrganizationId, method, endpoint, query, body string) string {

	message := APIKey + "\x00" + fmt.Sprint(time.ServerTime) + "\x00" + nonce + "\x00" + "\x00" + XOrganizationId + "\x00" + "\x00" + method + "\x00" + endpoint + "\x00" + query
	if method == http.MethodPost {
		message = message + "\x00" + body
	}

	h := hmac.New(sha256.New, []byte(APISecret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func parsingHeader(req *http.Request, endpoint string, cfg config.AppConfig, body string) error {

	time, err := GetTime()
	if err != nil {
		return errors.Wrap(err, "error getting timestamp")
	}
	req.Header.Add("X-Time", fmt.Sprint(time.ServerTime))

	nonce, err := utils.GenerateSecureToken(36)
	if err != nil {
		return errors.Wrap(err, "error creating nonce")
	}
	req.Header.Add("X-Nonce", nonce)

	RequestID, err := utils.GenerateSecureToken(36)
	if err != nil {
		return errors.Wrap(err, "error creating RequestID")
	}

	req.Header.Add("X-Request-Id", RequestID)
	req.Header.Add("X-User-Lang", "en")
	req.Header.Add("X-Organization-Id", cfg.XOrganizationId)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth", cfg.APIKey+":"+getDigest(cfg.APISecret, cfg.APIKey, time, nonce, cfg.XOrganizationId, req.Method, endpoint, req.URL.Query().Encode(), body))

	return nil
}
