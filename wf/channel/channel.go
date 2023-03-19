package channel

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"math"
	"wf_api/wf/channel/bilibili"
	"wf_api/wf/channel/leiting"
	"wf_api/wf/internal"
)

type Channel interface {
	GetChannelNo() string
	GetMedia() string
	Login(device *internal.Device, usr, pwd string) (*internal.GameUser, error)
	SendOtp(device *internal.Device, phone string) error
	OtpLogin(device *internal.Device, phone, otp string) (*internal.GameUser, error)
	CheckLogin(device *internal.Device, user *internal.GameUser) error
}

var (
	ErrInvalidChannel = errors.New("无效的渠道")
)

type Id uint8

const (
	LEITING Id = iota
	BILIBILI
)

func (i Id) New(device *internal.Device) Channel {
	switch i {
	case LEITING:
		return leiting.NewChannel()
	case BILIBILI:
		return bilibili.NewChannel(device)
	default:
		panic(ErrInvalidChannel)
	}
}

func ParseChannel(id uint8) (Id, error) {
	if id < uint8(LEITING) || id > uint8(BILIBILI) {
		return math.MaxUint8, ErrInvalidChannel
	}
	return Id(id), nil
}

type Pack struct {
	Id      Id `json:"id"`
	Channel `json:"data"`
}

func (c *Pack) UnmarshalJSON(data []byte) error {
	var v struct {
		Id   Id  `json:"id"`
		Data any `json:"data"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	c.Id = v.Id
	c.Channel = c.Id.New(nil)
	err = mapstructure.Decode(v.Data, &c.Channel)
	if err != nil {
		return err
	}
	return nil
}
