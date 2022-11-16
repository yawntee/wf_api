package wf

import (
	"errors"
	"net/http"
	"strconv"
	"wf_api/util"
)

var (
	ErrNoLogin = errors.New("您还未登录")
)

type Client struct {
	*Device
	Channel  Channel
	gameUser *GameUser
	encrypt  struct {
		androidId  string
		oaid       string
		mac        string
		imei       string
		terminInfo string
		osVer      string
	}
	loaded     bool
	udid       int
	shortUdid  int
	loginToken string
	viewerId   int
}

func NewClient(channelType ChannelType) *Client {
	c := &Client{
		Device: NewDevice(),
	}
	var _channel Channel
	switch channelType {
	case LEITING:
		_channel = &LeitingChannel{
			Client: c,
		}
	}
	c.Channel = _channel
	//init
	c.encrypt.androidId = encodeHeader(string(DataCipher.Enc([]byte(c.AndroidId))))
	c.encrypt.oaid = encodeHeader(string(DataCipher.Enc([]byte(c.Oaid))))
	c.encrypt.mac = encodeHeader(string(DataCipher.Enc([]byte(c.Mac))))
	c.encrypt.imei = encodeHeader(string(DataCipher.Enc([]byte(c.Imei))))
	c.encrypt.terminInfo = encodeHeader(string(DataCipher.Enc([]byte(Config.DeviceName))))
	c.encrypt.osVer = encodeHeader(string(DataCipher.Enc([]byte(Config.DeviceName))))
	return c
}

func (c *Client) Login(usr, pwd string) error {
	gameUser, err := c.Channel.Login(usr, pwd)
	if err != nil {
		return err
	}
	c.gameUser = gameUser
	err = c.signUp()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) signReq(r *http.Request) {
	body, err := PeekBody(r)
	if err != nil {
		panic(err)
	}
	sign := c.gameUser.Sid
	if c.viewerId != 0 {
		sign += strconv.Itoa(c.viewerId)
	}
	sign += r.URL.Path
	sign += string(body)
	sign = util.Sha1([]byte(sign))
	setHeaders(r, map[string]string{
		"Accept-Encoding": "deflate, gzip",
		"Accept":          "text/xml, application/xml, application/xhtml+xml, text/html;q=0.9, text/plain;q=0.8, text/css, image/png, image/jpeg, image/gif;q=0.8, application/x-shockwave-flash, video/mp4;q=0.9, flv-application/octet-stream;q=0.8, video/x-flv;q=0.7, audio/mp4, application/futuresplash, */*;q=0.5",
		"User-Agent":      "Mozilla/5.0 (Android; U; zh-CN) AppleWebKit/533.19.4 (KHTML, like Gecko) AdobeAIR/33.1",
		"x-flash-version": "33,1,1,620",
		"Connection":      "Keep-Alive",
		"Referer":         "app:/worldflipper_android_release.swf",
		"Content-Type":    "application/x-www-form-urlencoded",
		"PARAM":           sign,
		"UDID":            c.gameUser.Sid,
		"RES_VER":         Config.ResVer,
		"LOGIN_TOKEN":     c.loginToken,
		"SHORT_UDID":      strconv.Itoa(c.shortUdid),
		"ANDROID_ID":      c.encrypt.androidId,
		"MEDIA":           c.Channel.GetMedia(),
		"CHANNEL":         c.Channel.GetChannelNo(),
		"OAID":            c.encrypt.oaid,
		"MAC":             c.encrypt.mac,
		"IMEI":            c.encrypt.imei,
		"DEVICE_NAME":     Config.DeviceName + " " + Config.OsVer,
		"APP_VER":         Config.VersionName,
		"DEVICE":          "2",
	})
}
