package client

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"strings"
	"wf_api/server/wf/internal"
)

type ResultCode int

var (
	SUCCESS   ResultCode = 1
	UPDATE    ResultCode = 214
	OVERSPEED ResultCode = 204
	LOGINED   ResultCode = 223
)

func (c ResultCode) Msg() string {
	switch c {
	case SUCCESS:
		return "成功"
	case UPDATE:
		return "有新的资源"
	case OVERSPEED:
		return "操作速度过快，请稍后再试"
	case LOGINED:
		return "您的账号已在其它地方登录"
	default:
		return fmt.Sprintf("<code %d>", c)
	}
}

type GameResp[T any] struct {
	DataHeaders struct {
		DeviceId    int        `msgpack:"device_id"`
		ForceUpdate bool       `msgpack:"force_update"`
		ResultCode  ResultCode `msgpack:"result_code"`
		Servertime  int        `msgpack:"servertime"`
		ShortUdid   int        `msgpack:"short_udid"`
		ViewerId    int        `msgpack:"viewer_id"`
	} `msgpack:"data_headers"`
	Data T `msgpack:"data"`
}

func PostMsgpack[T any](c *Client, url string, body any, target *GameResp[T], handler func(req *http.Request)) {
	if !strings.HasSuffix(url, "/signup") && !c.inited {
		err := c.SignUp()
		if err != nil {
			panic(err)
		}
	}
	retry := 0
start:
	retry++
	wrappedBody := internal.WrapMsgpack(body)
	req, err := http.NewRequest("POST", url, bytes.NewReader(wrappedBody))
	if err != nil {
		panic(err)
	}
	if handler != nil {
		handler(req)
	}
	if internal.GlobalConfig.Debug {
		request, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}
		internal.DebugTitleMsg("<Request>", string(request))
		fmt.Printf("%+v\n", body)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if internal.GlobalConfig.Debug {
		request, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		internal.DebugTitleMsg("<Response>", string(request))
	}
	var cache GameResp[any]
	err = internal.UnwrapMsgpack(&resp.Body, &cache)
	if err != nil {
		panic(err)
	}
	if internal.GlobalConfig.Debug {
		fmt.Printf("%+v\n", cache)
	}
	if code := cache.DataHeaders.ResultCode; code != SUCCESS {
		switch code {
		case UPDATE:
			var data internal.GameUpdateData
			err := mapstructure.Decode(cache.Data, &data)
			if err != nil {
				panic(err)
			}
			internal.StartUpdateAssets(data)
		case OVERSPEED:
			panic(OVERSPEED.Msg())
		case LOGINED:
			err := c.SignUp()
			if err != nil {
				panic(err)
			}
			if retry >= 3 {
				panic(LOGINED.Msg())
			}
			goto start
		default:
			fmt.Println("RequestRawBody", string(wrappedBody))
			panic(errors.New(code.Msg()))
		}
	}
	target.DataHeaders = cache.DataHeaders
	err = mapstructure.Decode(cache.Data, &target.Data)
	if err != nil {
		panic(err)
	}
}
