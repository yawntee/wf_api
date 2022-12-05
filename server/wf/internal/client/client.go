package client

import (
	"wf_api/server/wf/channel"
	"wf_api/server/wf/internal"
)

type Client struct {
	*internal.Device `json:"device"`
	Channel          *channel.Pack      `json:"channel"`
	GameUser         *internal.GameUser `json:"game_user"`
	Encrypt          struct {
		AndroidId  string `json:"android_id,omitempty"`
		Oaid       string `json:"oaid,omitempty"`
		Mac        string `json:"mac,omitempty"`
		Imei       string `json:"imei,omitempty"`
		TerminInfo string `json:"termin_info,omitempty"`
		OsVer      string `json:"os_ver,omitempty"`
	} `json:"encrypt"`
	inited     bool
	udid       int
	shortudid  int
	logintoken string
	viewerId   int
	apiCount   int
}

func NewClient(id channel.Id) *Client {
	device := internal.NewDevice()
	c := &Client{
		Channel: &channel.Pack{
			Id:      id,
			Channel: id.New(device),
		},
		Device: device,
	}
	//loginGame
	c.Encrypt.AndroidId = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.AndroidId))))
	c.Encrypt.Oaid = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.Oaid))))
	c.Encrypt.Mac = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.Mac))))
	c.Encrypt.Imei = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(c.Imei))))
	c.Encrypt.TerminInfo = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(internal.GlobalConfig.DeviceName))))
	c.Encrypt.OsVer = internal.EncodeHeader(string(internal.DataCipher.Enc([]byte(internal.GlobalConfig.DeviceName))))
	return c
}

func (c *Client) Login(usr, pwd string) error {
	gameUser, err := c.Channel.Channel.Login(c.Device, usr, pwd)
	if err != nil {
		return err
	}
	c.GameUser = gameUser
	return nil
}

func (c *Client) SendOtp(phone string) error {
	err := c.Channel.SendOtp(c.Device, phone)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) OtpLogin(phone, otp string) error {
	gameUser, err := c.Channel.Channel.OtpLogin(c.Device, phone, otp)
	if err != nil {
		return err
	}
	c.GameUser = gameUser
	return nil
}

func (c *Client) CheckLogin() error {
	err := c.Channel.CheckLogin(c.Device, c.GameUser)
	if err != nil {
		return err
	}
	return nil
}
