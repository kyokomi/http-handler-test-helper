package example_test

import (
	"net/http"
	"testing"

	"github.com/kyokomi/hhth"
)

func TestHogeHandler(t *testing.T) {
	hhtHelper := hhth.New(http.DefaultServeMux)

	resp := hhtHelper.Get("/hoge",
		hhth.TestCaseStatusCode(http.StatusOK),
		hhth.TestCaseContentType("text/plain; charset=utf-8"),
	)
	if resp.Error() != nil {
		t.Errorf("error %s", resp.Error())
	}
	if resp.String() != "hoge" {
		t.Errorf("error response body hoge != %s", resp.String())
	}
}

func TestJSONParse(t *testing.T) {
	hhtHelper := hhth.New(http.DefaultServeMux)

	var resp map[string]interface{}
	if err := hhtHelper.Get("/hoge.json",
		hhth.TestCaseStatusCode(http.StatusOK),
		hhth.TestCaseContentType("application/json; charset=UTF-8"),
	).JSON(&resp); err != nil {
		t.Errorf("error %s", err)
	}

	if resp["name"].(string) != "hoge" {
		t.Errorf("error json response name != %s", resp["name"])
	}

	if resp["age"].(float64) != 20 {
		t.Errorf("error json response age != %s", resp["age"])
	}
}