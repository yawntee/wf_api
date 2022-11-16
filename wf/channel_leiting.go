package wf

import (
	"encoding/json"
	"errors"
	"wf_api/util"
)

type LeitingChannel struct {
	*Client
}

func (c LeitingChannel) GetChannelNo() string {
	return "110001"
}

func (c LeitingChannel) GetMedia() string {
	return "M311463"
}

func (c *LeitingChannel) Login(usr, pwd string) (*GameUser, error) {
	header := map[string]string{
		"LtAid":        Config.Game,
		"LtKid":        Config.LtKid,
		"LtUid":        "0",
		"User-Agent":   UserAgent(c.Device),
		"Content-Type": "application/xxx-www-form-urlencoded",
	}
	body := util.ToJson(map[string]any{
		"game":        Config.Game,
		"channelNo":   c.GetChannelNo(),
		"os":          Config.Os,
		"mmid":        Config.Mmid,
		"checkAuth":   Config.CheckAuth,
		"media":       c.GetMedia(),
		"versionName": Config.VersionName,
		"versionCode": Config.VersionCode,
		"password":    string(PwdCipher.Enc([]byte(pwd))),
		"accompany":   Config.Accompany,
		"face":        Config.Face,
		"serial":      Serial(c.Device),
		"username":    usr,
	})
	var resp struct {
		Status  int    `json:"Status,omitempty"`
		Data    string `json:"Data,omitempty"`
		Message string `json:"Message,omitempty"`
	}
	err := json.NewDecoder(*Post(
		"https://loginwf.leiting.com/sdk/login.do",
		LoginCipher.Enc(body),
		HeaderBinder(header),
	)).Decode(&resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return util.FromJson(LoginCipher.Dec([]byte(resp.Data)), &GameUser{}), nil
}

func (c *LeitingChannel) SendOtp(phone string) error {
	//TODO implement me
	panic("implement me")
}

func (c *LeitingChannel) OtpLogin(phone, code string) (*GameUser, error) {
	//TODO implement me
	panic("implement me")
}
