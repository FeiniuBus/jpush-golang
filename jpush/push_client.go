package jpush

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const pushPath = "https://api.jpush.cn/v3/push"

// PushClient is
type PushClient struct {
	AppKey       string
	MasterSecret string
}

// SendPush is
func (client *PushClient) SendPush(payload *PushPayload) error {
	s, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err1 := client.sendPost(pushPath, string(s))
	if err1 != nil {
		return err1
	}

	return nil
}

func (client *PushClient) sendPost(uri string, body string) error {
	request, err := http.NewRequest("POST", uri, strings.NewReader(body))
	if err != nil {
		return err
	}

	request.SetBasicAuth(client.AppKey, client.MasterSecret)
	request.Header.Add("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	s := string(result)
	if resp.StatusCode != 200 {
		return errors.New(s)
	}

	return nil
}

// NewPushClient is
func NewPushClient(appkey, secret string) *PushClient {
	p := new(PushClient)
	p.AppKey = appkey
	p.MasterSecret = secret

	return p
}
