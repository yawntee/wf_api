package client

import (
	"wf_api/server/wf/internal"
)

func (c *Client) SignUp() {
	if c.inited {
		return
	}
	if c.GameUser == nil {
		panic(internal.ErrNoLogin)
	}
	c.apiCount = 0
	body := map[string]any{
		"androidId":  c.Encrypt.AndroidId,
		"deviceId":   c.DeviceId,
		"terminInfo": c.Encrypt.TerminInfo,
		"oaid":       c.Encrypt.Oaid,
		"media":      c.Channel.GetMedia(),
		"mac":        c.Encrypt.Mac,
		"channelNo":  c.Channel.GetChannelNo(),
		"osVer":      c.Encrypt.OsVer,
	}
	var resp GameResp[struct {
		LoginToken string `msgpack:"login_token" mapstructure:"login_token"`
		NewAccount int    `msgpack:"newAccount" mapstructure:"new_account"`
	}]
	PostMsgpack(c, "https://shijtswygamegf.leiting.com//api/index.php/tool/signup", body, &resp, c.SignReq)
	c.udid = resp.DataHeaders.ShortUdid
	c.shortudid = resp.DataHeaders.ShortUdid
	c.viewerId = resp.DataHeaders.ViewerId
	c.logintoken = resp.Data.LoginToken
	c.inited = true
}

func (c *Client) LoadGameData() (*internal.GameUserInfo, error) {
	if c.apiCount > 0 {
		c.inited = false
	}
	c.SignUp()
	body := map[string]any{
		"oaid":                 c.Encrypt.Oaid,
		"viewer_id":            c.viewerId,
		"graphics_device_name": "OpenGL (Baseline Extended)",
		"platform_os_version":  "Android " + internal.GlobalConfig.OsVer,
		"imei":                 c.Encrypt.Imei,
		"device_token":         c.DeviceToken,
		"mac":                  c.Encrypt.Mac,
		"keychain":             c.viewerId,
		"device_id":            c.DeviceId,
	}
	var resp GameResp[internal.GameUserInfo]
	PostMsgpack(c, "https://shijtswygamegf.leiting.com//api/index.php/load", body, &resp, c.SignReqWithViewerId)
	return &resp.Data, nil
}
