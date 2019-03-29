package rest

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	headers := http.Header{}
	url := url.URL{}
	url.Scheme = "http"
	url.Host = "192.168.37.150:8080"
	url.Path = "/api/v1/namespaces/"
	client := NewHttpClient(&url, headers)

	resp, err := client.Get()
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	resp, err = client.Post(body, map[string]string{})
	if err != nil {
		t.Fatal(err.Error())
	}
	body, err = ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	t.Fatal("COM")
}

