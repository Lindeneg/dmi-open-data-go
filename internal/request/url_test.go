package request

import (
	"fmt"
	"testing"
)

var (
	r RequestConfig = RequestConfig{Key: "test-key", Type: "climateData"}
)

func TestConstructUrlWithoutQuery(t *testing.T) {
	expected := "https://dmigw.govcloud.dk/v2/climateData/miles/davis?api-key=test-key"
	result, err := constructUrl(r, "miles/davis", Query{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != expected {
		t.Fatalf(fmt.Sprintf("result %s does not match the expected value %s", result, expected))
	}
}

func TestConstructUrlWithQuery(t *testing.T) {
	expected := "https://dmigw.govcloud.dk/v2/climateData/miles/davis?limit=100&api-key=test-key&some=query"
	result, err := constructUrl(r, "miles/davis", Query{"some": "query", "limit": 100})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != expected {
		t.Fatalf(fmt.Sprintf("result %s does not match the expected value %s", result, expected))
	}
}
