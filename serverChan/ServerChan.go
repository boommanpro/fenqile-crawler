package serverChan

import (
	"encoding/json"
	"errors"
	"fenqile-crawler/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ResponseMessage struct {
	ErrorCode    int    `json:"errno"`
	ErrorMessage string `json:"errmsg"`
}

func SendMessage(config config.ServerChan, text, desp string) error {

	client := http.Client{}
	if config.Proxy {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://127.0.0.1:8888")
		}
		transport := &http.Transport{Proxy: proxy}
		client.Transport = transport
	}
	requestUrl := config.Server + config.SecretKey + ".send"
	data := url.Values{
		"text": {text},
		"desp": {desp},
	}
	request, _ := http.NewRequest(http.MethodPost, requestUrl, strings.NewReader(data.Encode()))

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	if config.Debug {
		bytes, _ := httputil.DumpRequest(request, true)
		fmt.Printf("ServerChan Request :\n%s\n\n", bytes)
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if config.Debug {
		bytes, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("%s", bytes)
	}
	var responseMsg ResponseMessage
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err

	}

	err = json.Unmarshal(bytes, &responseMsg)
	if err != nil {
		return err
	}
	if responseMsg.ErrorCode != 0 {
		return errors.New(responseMsg.ErrorMessage)
	}
	return nil
}
