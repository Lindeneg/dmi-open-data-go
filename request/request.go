package request

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lindeneg/dmi-open-data-go/v2/constants"
)

type RequestConfig struct {
	Key  string
	Type constants.ApiType
}

type Fetcher interface {
	Fetch(url string, q Query) (body []byte, err error)
}

type Query map[string]interface{}

func New(k string, t constants.ApiType) *RequestConfig {
	return &RequestConfig{Key: k, Type: t}
}

func (c *RequestConfig) Fetch(path string, q Query) ([]byte, error) {
	var (
		err  error
		url  string
		body []byte
	)
	if url, err = constructUrl(*c, path, q); err != nil {
		return body, err
	}
	if body, err = getBody(url); err != nil {
		return body, err
	}
	return body, nil
}

func getBody(url string) ([]byte, error) {
	var (
		err  error
		body []byte
		resp *http.Response
	)

	if resp, err = http.Get(url); err != nil {
		return body, err
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return body, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, errors.New(string(body))
	}

	return body, nil
}
