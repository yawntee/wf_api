package internal

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"net/http/httputil"
)

type ResultCode int

var (
	SUCCESS ResultCode = 1
)

var httpClient = http.Client{}

func SetHeaders(r *http.Request, headers map[string]string) {
	for k, v := range headers {
		if v != "" {
			r.Header.Set(k, v)
		}
	}
}

func HeaderBinder(headers map[string]string) func(r *http.Request) {
	return func(r *http.Request) {
		SetHeaders(r, headers)
	}
}

func Post(url string, body []byte, handler func(req *http.Request)) *io.ReadCloser {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	handler(req)
	if Config.Debug {
		request, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}
		slog.Debug("<Request>", string(request))
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	if Config.Debug {
		request, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		slog.Debug("<Response>", string(request))
	}
	return &resp.Body
}

func PeekBody(r *http.Request) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}
	body := buf.Bytes()
	r.Body = io.NopCloser(bytes.NewReader(body))
	return body, err
}

func (c ResultCode) Msg() string {
	switch c {
	case SUCCESS:
		return "success"
	default:
		return fmt.Sprintf("<code %d>", c)
	}
}
