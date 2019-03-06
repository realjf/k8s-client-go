package rest

import (
	"net/http"
	"net/url"
	"testing"
	"io/ioutil"
)

func TestNewHttpClient(t *testing.T) {
	headers := http.Header{}
	url := url.URL{}
	url.Scheme = "http"
	url.Host = "www.baidu.com"
	client := NewHttpClient(&url, headers)

	resp, err := client.Get()
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	t.Fatal("COM")
}

