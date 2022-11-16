package module

import (
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"wf_api/wf/internal"
	"wf_api/wf/internal/context"
)

type LoginModule struct {
	*context.ClientContext
}

func NewLoginModule(clientContext *context.ClientContext) *LoginModule {
	return &LoginModule{ClientContext: clientContext}
}

func (c *LoginModule) SignUp() error {
	if c.GameUser == nil {
		return internal.ErrNoLogin
	}
	c.Loaded = false
	body := internal.WrapMsgpack(map[string]any{
		"androidId":  c.Encrypt.AndroidId,
		"deviceId":   c.DeviceId,
		"terminInfo": c.Encrypt.TerminInfo,
		"oaid":       c.Encrypt.Oaid,
		"media":      c.Channel.GetMedia(),
		"mac":        c.Encrypt.Mac,
		"channelNo":  c.Channel.GetChannelNo(),
		"osVer":      c.Encrypt.OsVer,
	})
	var resp internal.GameResp[struct {
		LoginToken string `msgpack:"login_token"`
		NewAccount int    `msgpack:"newAccount"`
	}]
	err := internal.UnwrapMsgpack(internal.Post("https://shijtswygamegf.leiting.com//api/index.php/tool/signup", body, c.SignReq), &resp)
	if err != nil {
		return errors.WithMessage(err, "载入游戏数据失败")
	}
	slog.Debug(internal.FuncName(), resp)
	if code := resp.DataHeaders.ResultCode; code != internal.SUCCESS {
		return errors.New(code.Msg())
	}
	c.Udid = resp.DataHeaders.ShortUdid
	c.ShortUdid = resp.DataHeaders.ShortUdid
	c.ViewerId = resp.DataHeaders.ViewerId
	c.LoginToken = resp.Data.LoginToken
	return nil
}

func (c *LoginModule) Load() error {
	if c.Loaded {
		err := c.SignUp()
		if err != nil {
			return err
		}
	}
	body := internal.WrapMsgpack(map[string]any{
		"oaid":                 c.Encrypt.Oaid,
		"viewer_id":            c.ViewerId,
		"graphics_device_name": "OpenGL (Baseline Extended)",
		"platform_os_version":  "Android " + internal.Config.OsVer,
		"imei":                 c.Encrypt.Imei,
		"device_token":         c.DeviceToken,
		"mac":                  c.Encrypt.Mac,
		"keychain":             c.ViewerId,
		"device_id":            c.DeviceId,
	})
	var resp internal.GameResp[internal.GameUserInfo]
	err := internal.UnwrapMsgpack(internal.Post("https://shijtswygamegf.leiting.com//api/index.php/load", body, c.SignReq), &resp)
	if err != nil {
		return err
	}
	slog.Debug(internal.FuncName(), resp)
	if code := resp.DataHeaders.ResultCode; code != internal.SUCCESS {
		return errors.New(code.Msg())
	}
	c.Loaded = true
	return nil
}
