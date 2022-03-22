package request

import (
	"errors"
	"fmt"

	"github.com/lindeneg/dmi-open-data-go/internal/constants"
)

func constructUrl(r RequestConfig, p string, q Query) (string, error) {
	base := fmt.Sprintf("%s/v2/%s/%s", constants.BaseUrl, r.Type, p)
	if r.Key == "" {
		return "", errors.New("no API key specified")
	}
	q["api-key"] = r.Key
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
