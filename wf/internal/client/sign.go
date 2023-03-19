package client

import (
	"net/http"
	"strconv"
	"wf_api/util"
	"wf_api/wf/internal"
)

func (c *Client) SignReq(r *http.Request) {
	c._SignReq(r, false)
}

func (c *Client) SignReqWithViewerId(r *http.Request) {
	c._SignReq(r, true)
}

func (c *Client) _SignReq(r *http.Request, viewerId bool) {
	body, err := internal.PeekBody(r)
	if err != nil {
		panic(err)
	}
	sign := c.GameUser.Sid
	if viewerId {
		sign += strconv.Itoa(c.viewerId)
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
		"RES_VER":         internal.GlobalConfig.ResVer,
		"LOGIN_TOKEN":     c.logintoken,
		"SHORT_UDID":      strconv.Itoa(c.shortudid),
		"ANDROID_ID":      c.Encrypt.AndroidId,
		"MEDIA":           c.Channel.GetMedia(),
		"CHANNEL":         c.Channel.GetChannelNo(),
		"OAID":            c.Encrypt.Oaid,
		"MAC":             c.Encrypt.Mac,
		"IMEI":            c.Encrypt.Imei,
		"DEVICE_NAME":     internal.GlobalConfig.DeviceName + " " + internal.GlobalConfig.OsVer,
		"APP_VER":         internal.GlobalConfig.VersionName,
		"DEVICE":          "2",
	})
}
