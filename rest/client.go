package rest

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

type IHttpClient interface {
	Get() (resp *http.Response, err error)
	Post(body []byte, headers map[string]string) (resp *http.Response, err error)
	Put(body []byte, headers map[string]string) (resp *http.Response, err error)
	Delete() (resp *http.Response, err error)
	Patch() (resp *http.Response, err error)


	dial(method string) (resp *http.Response, err error)
	SetHeader(key string, values ...string)
	SetUrl(url *url.URL)
	GetPath() string
}

type HttpClient struct {
	url *url.URL

	headers http.Header

	namespace string
	resourceType string
	resourceName string

	timeout time.Duration

	err error
	body io.Reader

	ctx context.Context

	Client *http.Client
}

func (c *HttpClient) getUrl() string {
	// 拼接地址

	return c.url.String()
}

func (c *HttpClient) GetPath() string {
	return c.url.Path
}

func (c *HttpClient) Get() (resp *http.Response, err error) {
	return c.dial("GET")
}

func (c *HttpClient) Post(body []byte, headers map[string]string) (resp *http.Response, err error) {
	for k, v := range headers {
		c.headers.Del(k)
		c.SetHeader(k, v)
	}

	c.body = bytes.NewReader(body)

	return c.dial("POST")
}

func (c *HttpClient) Delete() (resp *http.Response, err error) {

	return c.dial("DELETE")
}

func (c *HttpClient) Put(body []byte, headers map[string]string) (resp *http.Response, err error) {
	for k, v := range headers {
		c.headers.Del(k)
		c.SetHeader(k, v)
	}

	c.body = bytes.NewReader(body)

	return c.dial("PUT")
}

func (c *HttpClient) Patch() (resp *http.Response, err error) {
	return c.dial("PATCH")
}

func (c *HttpClient) dial(method string) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, c.getUrl(), c.body)
	if err != nil {
		return nil, err
	}
	// 设置header头部
	req.Header = c.headers
	// 用json数据格式
	//req.Header.Add("Content-Type", "application/json")

	return c.Client.Do(req)
}

func (c *HttpClient) SetHeader(key string, values ...string) {
	if c.headers == nil {
		c.headers = http.Header{}
	}
	c.headers.Del(key)
	for _, header := range values {
		c.headers.Add(key, header)
	}
}

func (c *HttpClient) SetUrl(url *url.URL) {
	c.url = url
}

func NewHttpClient(url *url.URL, headers http.Header) IHttpClient {
	return &HttpClient{
		url:     url,
		headers: headers,
		Client:  &http.Client{},
		body:    bytes.NewReader([]byte{}),
	}
}