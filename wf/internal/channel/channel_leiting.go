package channel

import (
	"encoding/json"
	"errors"
	"golang.org/x/exp/slog"
	"wf_api/util"
	"wf_api/wf/internal"
	"wf_api/wf/internal/context"
)

type LeitingChannel struct {
	*context.ClientContext
}

func (c LeitingChannel) GetChannelNo() string {
	return "110001"
}

func (c LeitingChannel) GetMedia() string {
	return "M311463"
}

func (c *LeitingChannel) Login(usr, pwd string) (*internal.GameUser, error) {
	header := map[string]string{
		"LtAid":        internal.Config.Game,
		"LtKid":        internal.Config.LtKid,
		"LtUid":        "0",
		"User-Agent":   internal.UserAgent(c.Device),
		"Content-Type": "application/xxx-www-form-urlencoded",
	}
	body := util.ToJson(map[string]any{
		"game":        internal.Config.Game,
		"channelNo":   c.GetChannelNo(),
		"os":          internal.Config.Os,
		"mmid":        internal.Config.Mmid,
		"checkAuth":   internal.Config.CheckAuth,
		"media":       c.GetMedia(),
		"versionName": internal.Config.VersionName,
		"versionCode": internal.Config.VersionCode,
		"password":    string(internal.PwdCipher.Enc([]byte(pwd))),
		"accompany":   internal.Config.Accompany,
		"face":        internal.Config.Face,
		"serial":      internal.Serial(c.Device),
		"username":    usr,
	})
	var resp struct {
		Status  int    `json:"Status,omitempty"`
		Data    string `json:"Data,omitempty"`
		Message string `json:"Message,omitempty"`
	}
	err := json.NewDecoder(*internal.Post(
		"https://loginwf.leiting.com/sdk/login.do",
		internal.LoginCipher.Enc(body),
		internal.HeaderBinder(header),
	)).Decode(&resp)
	if err != nil {
		return nil, err
	}
	slog.Debug(internal.FuncName(), resp)
	if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return util.FromJson(internal.LoginCipher.Dec([]byte(resp.Data)), &internal.GameUser{}), nil
}

func (c *LeitingChannel) SendOtp(phone string) error {
	//TODO implement me
	panic("implement me")
}

func (c *LeitingChannel) OtpLogin(phone, code string) (*internal.GameUser, error) {
	//TODO implement me
	panic("implement me")
}
