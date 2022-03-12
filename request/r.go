package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/lindeneg/dmi-open-data-go/v2/constants"
)

type Config struct {
	Key  string
	Type constants.ApiType
}

type Fetcher interface {
	Fetch(u string) (rb responseBody, err error)
}

type responseBody struct {
	body []byte
	res  *http.Response
}

type Result struct {
	res  []byte
	done chan struct{}
	mux  sync.Mutex
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
		return []byte{}, errors.New(string(body))
	}

	return body, nil
}

func H() {
	fmt.Println(fmt.Sprintf("%s/v2/climateData/%s", constants.BaseUrl, "collections/stationValue/items"))
	b, e := getBody(fmt.Sprintf("%s/v2/climateData/%s", constants.BaseUrl, "collections/stationValue/items"))

	fmt.Println(e)
	fmt.Println(b)
}
