package request

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lindeneg/dmi-open-data-go/constants"
)

type Config struct {
	Key  string
	Type constants.ApiType
}

type Request struct {
	Res *Response
	Config
}

type Query map[string]interface{}

func NewRequest(k string, a constants.ApiType) (*Request, *Response) {
	res := &Response{}
	req := &Request{
		Res: res,
		Config: Config{
			Key:  k,
			Type: a,
		},
	}
	return req, res
}

func (r *Request) Get(p string, q Query) {
	rs := r.Res
	url, err := constructUrl(*r, p, q)
	if err != nil {
		rs.Err = err
	} else {
		res, err := http.Get(url)
		rs.Code = res.StatusCode
		if err != nil {
			rs.Err = err
		} else {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				rs.Err = err
			} else {
				if rs.Code == 200 {
					rs.Body = body
				} else {
					rs.Err = errors.New(string(body))
				}
			}
		}
	}
}
