package visit

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

var defaultclient *fasthttp.Client
var methodMap map[string]string

func init() {
	defaultclient = &fasthttp.Client{
		MaxConnsPerHost: 1000,
		ReadTimeout:     2 * time.Second,
		WriteTimeout:    2 * time.Second,
	}
	methodMap = make(map[string]string)
	methodMap["get"] = "GET"
	methodMap["GET"] = "GET"
	methodMap["post"] = "POST"
	methodMap["POST"] = "POST"
	methodMap["put"] = "PUT"
	methodMap["PUT"] = "PUT"
	methodMap["delete"] = "DELETE"
	methodMap["DELETE"] = "DELETE"

}

func CreateHttpClient(readTimeOut, writeTimeOut, maxConnsPerHost int) *fasthttp.Client {
	c := &fasthttp.Client{
		MaxConnsPerHost: maxConnsPerHost,
		ReadTimeout:     time.Duration(readTimeOut) * time.Millisecond,
		WriteTimeout:    time.Duration(writeTimeOut) * time.Millisecond,
	}
	return c
}

type ReqParmater struct {
	Method      string
	Url         string
	ContentType string
	Headers     map[string]string
	Body        []byte
	UserAgent   string
}

func DoRequest(r *ReqParmater, client *fasthttp.Client) ([]byte, error) {
	m := methodMap[r.Method]
	if m == "" {
		return []byte{}, errors.New("Unsupported method")
	}
	var c *fasthttp.Client
	if client == nil {
		c = defaultclient
	} else {
		c = client
	}
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.Header.SetMethod(m)
	req.SetRequestURI(r.Url)
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	if len(r.Body) > 0 {
		req.SetBody(r.Body)
	}
	if len(r.ContentType) > 0 {
		req.Header.SetContentType(r.ContentType)
	}
	if len(r.UserAgent) > 0 {
		req.Header.SetUserAgent(r.UserAgent)
	}
	err := c.Do(req, resp)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return []byte{}, errors.New(fmt.Sprintf("visit error, http code: %d", resp.StatusCode()))
	}
	resBody := resp.Body()
	res := make([]byte, len(resBody))
	copy(res, resBody)
	return res, nil
}
