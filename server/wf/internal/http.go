package internal

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
	"wf_api/server/wf/internal/context"
)

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

func Post(url string, body []byte, handler func(req *http.Request)) io.ReadCloser {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	if handler != nil {
		handler(req)
	}
	if GlobalConfig.Debug {
		request, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}
		DebugTitleMsg("<Request>", string(request))
	}
	context.HttpMutex.Lock()
	resp, err := http.DefaultClient.Do(req)
	time.Sleep(time.Second / 100)
	context.HttpMutex.Unlock()
	if err != nil {
		panic(err)
	}
	if GlobalConfig.Debug {
		response, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		DebugTitleMsg("<Response>", string(response))
	}
	return resp.Body
}

func PostForm(url string, body url.Values, handler func(req *http.Request)) io.ReadCloser {
	return Post(url, []byte(body.Encode()), func(req *http.Request) {
		if handler != nil {
			handler(req)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	})
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
