package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpClient struct {
	BaseUrl string
}

type HttpClientInterface interface {
	Get(uri string, body interface{}) error
}

func NewHttpClient(baseUrl string) HttpClientInterface {

	return &HttpClient{
		BaseUrl: baseUrl,
	}
}

func (hc *HttpClient) Get(uri string, body interface{}) error {

	url := fmt.Sprintf("%s%s", hc.BaseUrl, uri)

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	bd := res.Body

	if err := json.NewDecoder(bd).Decode(body); err != nil {
		return err
	}

	return err
}
