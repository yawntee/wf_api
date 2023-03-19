package leiting

import (
	"encoding/json"
	"github.com/pkg/errors"
	"wf_api/util"
	"wf_api/wf/internal"
)

type stdResp struct {
	Status  int    `json:"Status,omitempty"`
	Data    string `json:"data,omitempty"`
	Message string `json:"Message,omitempty"`
}

type Channel struct {
}

func NewChannel() *Channel {
	return &Channel{}
}

func (c *Channel) GetChannelNo() string {
	return "110001"
}

func (c *Channel) GetMedia() string {
	return "M311463"
}

func (c *Channel) Login(device *internal.Device, usr, pwd string) (*internal.GameUser, error) {
	header := map[string]string{
		"LtAid":        internal.GlobalConfig.Game,
		"LtKid":        internal.GlobalConfig.LtKid,
		"LtUid":        "0",
		"User-Agent":   internal.UserAgent(device),
		"Content-Type": "application/xxx-www-form-urlencoded",
	}
	body := util.ToJson(map[string]any{
		"game":        internal.GlobalConfig.Game,
		"channelNo":   c.GetChannelNo(),
		"os":          internal.GlobalConfig.Os,
		"mmid":        internal.GlobalConfig.Mmid,
		"checkAuth":   internal.GlobalConfig.CheckAuth,
		"media":       c.GetMedia(),
		"versionName": internal.GlobalConfig.VersionName,
		"versionCode": internal.GlobalConfig.VersionCode,
		"password":    string(internal.PwdCipher.Enc([]byte(pwd))),
		"accompany":   internal.GlobalConfig.Accompany,
		"face":        internal.GlobalConfig.Face,
		"serial":      internal.Serial(device),
		"username":    usr,
	})
	var resp stdResp
	err := json.NewDecoder(internal.Post(
		"https://loginwf.leiting.com/sdk/login.do",
		internal.LoginCipher.Enc(body),
		internal.HeaderBinder(header),
	)).Decode(&resp)
	if err != nil {
		return nil, err
	}
	internal.DebugMsg(resp)
	if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return util.FromJson(internal.LoginCipher.Dec([]byte(resp.Data)), &internal.GameUser{}), nil
}

func (c *Channel) SendOtp(device *internal.Device, phone string) error {
	header := map[string]string{
		"LtAid":        internal.GlobalConfig.Game,
		"LtKid":        "v3.0.8",
		"LtUid":        "0",
		"Content-Type": "application/json",
		"User-Agent":   internal.UserAgent(device),
	}
	body := util.ToJson(map[string]any{
		"phone": phone,
		"type":  "LOGIN_MT",
	})
	var resp stdResp
	err := json.NewDecoder(internal.Post(
		"https://member.leiting.com/aes/message/send_phone_code",
		internal.LoginCipher.Enc(body),
		internal.HeaderBinder(header),
	)).Decode(&resp)
	if err != nil {
		return err
	}
	internal.DebugMsg(resp)
	if resp.Status != 0 {
		return errors.New(resp.Message)
	}
	return nil
}

func (c *Channel) OtpLogin(device *internal.Device, phone, code string) (*internal.GameUser, error) {
	header := map[string]string{
		"LtAid":        internal.GlobalConfig.Game,
		"LtKid":        internal.GlobalConfig.LtKid,
		"LtUid":        "0",
		"User-Agent":   internal.UserAgent(device),
		"Content-Type": "application/xxx-www-form-urlencoded",
	}
	body := util.ToJson(map[string]any{
		"game":        internal.GlobalConfig.Game,
		"code":        code,
		"channelNo":   c.GetChannelNo(),
		"os":          internal.GlobalConfig.Os,
		"mmid":        internal.GlobalConfig.Mmid,
		"checkAuth":   internal.GlobalConfig.CheckAuth,
		"media":       c.GetMedia(),
		"type":        "mt",
		"versionName": internal.GlobalConfig.VersionName,
		"versionCode": internal.GlobalConfig.VersionCode,
		"accompany":   internal.GlobalConfig.Accompany,
		"face":        internal.GlobalConfig.Face,
		"phone":       phone,
		"serial":      internal.Serial(device),
	})
	var resp stdResp
	err := json.NewDecoder(internal.Post(
		"https://loginwf.leiting.com/sdk/code_login.do",
		internal.LoginCipher.Enc(body),
		internal.HeaderBinder(header),
	)).Decode(&resp)
	if err != nil {
		return nil, err
	}
	internal.DebugMsg(resp)
	if resp.Status != 0 {
		return nil, errors.New(string(internal.LoginCipher.Dec([]byte(resp.Data))))
	}
	return util.FromJson(internal.LoginCipher.Dec([]byte(resp.Data)), &internal.GameUser{}), nil
}

func (c *Channel) CheckLogin(device *internal.Device, user *internal.GameUser) error {
	header := map[string]string{
		"LtAid":        internal.GlobalConfig.Game,
		"LtKid":        internal.GlobalConfig.LtKid,
		"LtUid":        "0",
		"User-Agent":   internal.UserAgent(device),
		"Content-Type": "application/xxx-www-form-urlencoded",
	}
	body := util.ToJson(map[string]any{
		"game":        internal.GlobalConfig.Game,
		"channelNo":   c.GetChannelNo(),
		"os":          internal.GlobalConfig.Os,
		"mmid":        internal.GlobalConfig.Mmid,
		"checkAuth":   internal.GlobalConfig.CheckAuth,
		"media":       c.GetMedia(),
		"versionName": internal.GlobalConfig.VersionName,
		"versionCode": internal.GlobalConfig.VersionCode,
		"sid":         user.Sid,
		"token":       user.Token,
		"accompany":   internal.GlobalConfig.Accompany,
		"face":        internal.GlobalConfig.Face,
		"serial":      internal.Serial(device),
		"username":    user.Username,
	})
	var resp stdResp
	err := json.NewDecoder(internal.Post(
		"https://loginwf.leiting.com/sdk/check_login.do",
		internal.LoginCipher.Enc(body),
		internal.HeaderBinder(header),
	)).Decode(&resp)
	if err != nil {
		return err
	}
	internal.DebugMsg(resp)
	if resp.Status != 0 {
		return errors.New(resp.Message)
	}
	newUser := util.FromJson(internal.LoginCipher.Dec([]byte(resp.Data)), &internal.GameUser{})
	user.Sid = newUser.Sid
	user.Username = newUser.Username
	return nil
}
