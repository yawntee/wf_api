package client

import (
	"github.com/pkg/errors"
	"net/http"
	"wf_api/server/wf/internal"
)

func (c *Client) SignUp() error {
	if c.inited {
		return nil
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
	if checkUpdate(c) {
		return errors.New("游戏资源更新中...")
	}
	return nil
}

func checkUpdate(c *Client) bool {
	body := map[string]any{
		"viewer_id":            c.viewerId,
		"client_asset_version": internal.GlobalConfig.ResVer,
	}
	var resp GameResp[internal.GameUpdateData]
	PostMsgpack(c, "https://shijtswygamegf.leiting.com//api/index.php/asset/get_path", body, &resp, func(req *http.Request) {
		c.SignReqWithViewerId(req)
		if internal.GlobalConfig.ResVer == "" {
			req.Header.Del("RES_VER")
		}
		req.Header.Set("ASSET_SIZE", "shortened")
	})
	if diff := resp.Data.Diff; diff != nil && len(diff) > 0 {
		internal.StartUpdateAssets(resp.Data)
		return true
	}
	return false
}

func (c *Client) LoadGameData() (*internal.GameUserInfo, error) {
	if c.apiCount > 0 {
		c.inited = false
	}
	err := c.SignUp()
	if err != nil {
		return nil, err
	}
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
