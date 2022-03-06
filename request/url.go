package request

import (
	"errors"
	"fmt"

	"github.com/lindeneg/dmi-open-data-go/constants"
)

func constructUrl(r Request, p string, q Query) (string, error) {
	base := fmt.Sprintf("%s/v2/%s/%s", constants.BaseUrl, r.Config.Type, p)
	if r.Config.Key == "" {
		return "", errors.New("no API key specified")
	}
	q["api-key"] = r.Config.Key
	return appendQuery(base, q), nil
}

func appendQuery(b string, q Query) string {
	r := "?"
	i := 0
	l := len(q) - 1
	for key, value := range q {
		if value != nil {
			r += fmt.Sprintf("%s=%v", key, value)
			if i != l {
				r += "&"
			}
		}
		i++
	}
	return b + r
}
