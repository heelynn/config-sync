package client

import (
	"config-sync/pkg/zlog"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

type HttpClient struct {
	Hosts   []string
	Method  HttpMethod
	Url     string
	Headers map[string]string
	Params  map[string]string
	Body    string
}

func NewHttpClientHosts(hosts []string, method HttpMethod, url string) *HttpClient {
	return &HttpClient{
		Hosts:  hosts,
		Method: method,
		Url:    url,
	}
}

func NewHttpClient(host string, method HttpMethod, url string) *HttpClient {
	return &HttpClient{
		Hosts:  []string{host},
		Method: method,
		Url:    url,
	}
}

func (c *HttpClient) AddHeader(key, value string) {
	if c.Headers == nil {
		c.Headers = make(map[string]string)
	}
	c.Headers[key] = value
}

func (c *HttpClient) AddParam(key, value string) {
	if c.Params == nil {
		c.Params = make(map[string]string)
	}
	c.Params[key] = value
}

func (c *HttpClient) SetBody(body string) {
	c.Body = body
}

// DoGetResponse 发送请求并获取响应
func (c *HttpClient) doGetResponseHostIndex(hostIndex int) (*http.Response, error) {
	if c.Hosts == nil || len(c.Hosts) == 0 {
		return nil, fmt.Errorf("hosts is empty")
	}
	baseURL := c.Hosts[hostIndex] + c.Url
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	params := url.Values{}
	if c.Params != nil && len(c.Params) > 0 {
		for k, v := range c.Params {
			params.Set(k, v)
		}
	}
	// 将查询参数添加到URL
	urlWithParams := baseURL + "?" + params.Encode()
	req, err := http.NewRequest(string(c.Method), urlWithParams, strings.NewReader(c.Body))
	if err != nil {
		return nil, err
	}
	if c.Headers != nil && len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	zlog.Logger.Infof("HTTP request url: %s, method: %s ,response status: [%d]", urlWithParams, req.Method, resp.StatusCode)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Do 发送请求并获取响应, 并返回状态码和响应体
func (c *HttpClient) doHostIndex(hostIndex int) (statusCode int, body []byte, err error) {
	resp, err := c.doGetResponseHostIndex(hostIndex)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}

// DoGetResponse 发送请求并获取响应
func (c *HttpClient) DoGetResponse() (*http.Response, error) {
	return c.doGetResponseHostIndex(0)
}

// DoGetResponse 发送请求并获取响应
func (c *HttpClient) Do() (statusCode int, body []byte, err error) {
	return c.doHostIndex(0)
}

// DoInstances 发送请求并获取响应, 节点选择失败时, 尝试下一个
func (c *HttpClient) DoInstances() (stausCode int, body []byte, err error) {
	if c.Hosts == nil || len(c.Hosts) == 0 {
		return 0, nil, fmt.Errorf("hosts is empty")
	}
	for i, _ := range c.Hosts {
		code, bytes, err := c.doHostIndex(i)
		if err != nil {
			// 集群选择失败，尝试下一个集群
			zlog.Logger.Error(err)
			continue
		}
		return code, bytes, nil
	}
	return 0, nil, fmt.Errorf("all host failed")
}

// DoGetResponseInstances 发送请求并获取响应, 节点选择失败时, 尝试下一个
func (c *HttpClient) DoGetResponseInstances() (*http.Response, error) {
	if c.Hosts == nil || len(c.Hosts) == 0 {
		return nil, fmt.Errorf("hosts is empty")
	}
	for i, _ := range c.Hosts {
		resp, err := c.doGetResponseHostIndex(i)
		if err != nil {
			// 集群选择失败，尝试下一个集群
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("all host failed")
}
