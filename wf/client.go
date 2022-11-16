package main

import (
	"net/http"
	"strconv"
	"wf_api/util"
	"wf_api/wf/internal"
	"wf_api/wf/internal/channel"
	"wf_api/wf/internal/context"
	"wf_api/wf/internal/module"
)

type ChannelType int

const (
	LEITING ChannelType = iota
	BILIBILI
)

type Client struct {
	*context.ClientContext
	*module.LoginModule
	*module.ShopModule
}

func NewClient(channelType ChannelType) *Client {
	cx := context.NewClientContext()
	c := &Client{
		ClientContext: cx,
		LoginModule:   module.NewLoginModule(cx),
		ShopModule:    module.NewShopModule(cx),
	}
	var _channel internal.Channel
	switch channelType {
	case LEITING:
		_channel = &channel.LeitingChannel{
			ClientContext: c.ClientContext,
		}
	}
	c.Channel = _channel
	//init
	c.Encrypt.AndroidId = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.AndroidId))))
	c.Encrypt.Oaid = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.Oaid))))
	c.Encrypt.Mac = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.Mac))))
	c.Encrypt.Imei = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.Imei))))
	c.Encrypt.TerminInfo = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(internal.Config.DeviceName))))
	c.Encrypt.OsVer = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(internal.Config.DeviceName))))
	return c
}

func (c *Client) Login(usr, pwd string) error {
	gameUser, err := c.Channel.Login(usr, pwd)
	if err != nil {
		return err
	}
	c.GameUser = gameUser
	err = c.SignUp()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) signReq(r *http.Request) {
	body, err := internal.PeekBody(r)
	if err != nil {
		panic(err)
	}
	sign := c.GameUser.Sid
	if c.ViewerId != 0 {
		sign += strconv.Itoa(c.ViewerId)
	}
	sign += r.URL.Path
	sign += string(body)
	sign = util.Sha1([]byte(sign))
	internal.SetHeaders(r, map[string]string{
		"Accept-Encoding": "deflate, gzip",
		"Accept":          "text/xml, application/xml, application/xhtml+xml, text/html;q=0.9, text/plain;q=0.8, text/css, image/png, image/jpeg, image/gif;q=0.8, application/x-shockwave-flash, video/mp4;q=0.9, flv-application/octet-stream;q=0.8, video/x-flv;q=0.7, audio/mp4, application/futuresplash, */*;q=0.5",
		"User-Agent":      "Mozilla/5.0 (Android; U; zh-CN) AppleWebKit/533.19.4 (KHTML, like Gecko) AdobeAIR/33.1",
		"x-flash-version": "33,1,1,620",
		"Connection":      "Keep-Alive",
		"Referer":         "app:/worldflipper_android_release.swf",
		"Content-Type":    "application/x-www-form-urlencoded",
		"PARAM":           sign,
		"UDID":            c.GameUser.Sid,
		"RES_VER":         internal.Config.ResVer,
		"LOGIN_TOKEN":     c.LoginToken,
		"SHORT_UDID":      strconv.Itoa(c.ShortUdid),
		"ANDROID_ID":      c.Encrypt.AndroidId,
		"MEDIA":           c.Channel.GetMedia(),
		"CHANNEL":         c.Channel.GetChannelNo(),
		"OAID":            c.Encrypt.Oaid,
		"MAC":             c.Encrypt.Mac,
		"IMEI":            c.Encrypt.Imei,
		"DEVICE_NAME":     internal.Config.DeviceName + " " + internal.Config.OsVer,
		"APP_VER":         internal.Config.VersionName,
		"DEVICE":          "2",
	})
}
