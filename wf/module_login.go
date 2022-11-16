package wf

import (
	"github.com/pkg/errors"
)

func (c *Client) signUp() error {
	if c.gameUser == nil {
		return ErrNoLogin
	}
	c.loaded = false
	body := wrapMsgpack(map[string]any{
		"androidId":  c.encrypt.androidId,
		"deviceId":   c.DeviceId,
		"terminInfo": c.encrypt.terminInfo,
		"oaid":       c.encrypt.oaid,
		"media":      c.Channel.GetMedia(),
		"mac":        c.encrypt.mac,
		"channelNo":  c.Channel.GetChannelNo(),
		"osVer":      c.encrypt.osVer,
	})
	var resp GameResp[struct {
		LoginToken string `msgpack:"login_token"`
		NewAccount int    `msgpack:"newAccount"`
	}]
	err := unwrapMsgpack(Post("https://shijtswygamegf.leiting.com//api/index.php/tool/signup", body, c.signReq), &resp)
	if err != nil {
		return errors.WithMessage(err, "载入游戏数据失败")
	}
	if Config.Debug {
		DebugMsgf("Signup", "%+v\n", resp)
	}
	if code := resp.DataHeaders.ResultCode; code != SUCCESS {
		return errors.New(code.Msg())
	}
	c.udid = resp.DataHeaders.ShortUdid
	c.shortUdid = resp.DataHeaders.ShortUdid
	c.viewerId = resp.DataHeaders.ViewerId
	c.loginToken = resp.Data.LoginToken
	return nil
}

func (c *Client) Load() error {
	if c.loaded {
		err := c.signUp()
		if err != nil {
			return err
		}
	}
	body := wrapMsgpack(map[string]any{
		"oaid":                 c.encrypt.oaid,
		"viewer_id":            c.viewerId,
		"graphics_device_name": "OpenGL (Baseline Extended)",
		"platform_os_version":  "Android " + Config.OsVer,
		"imei":                 c.encrypt.imei,
		"device_token":         c.DeviceToken,
		"mac":                  c.encrypt.mac,
		"keychain":             c.viewerId,
		"device_id":            c.DeviceId,
	})
	var resp GameResp[GameUserInfo]
	err := unwrapMsgpack(Post("https://shijtswygamegf.leiting.com//api/index.php/load", body, c.signReq), &resp)
	if err != nil {
		return err
	}
	if Config.Debug {
		DebugMsgf("Load", "%+v\n", resp)
	}
	if code := resp.DataHeaders.ResultCode; code != SUCCESS {
		return errors.New(code.Msg())
	}
	c.loaded = true
	return nil
}
