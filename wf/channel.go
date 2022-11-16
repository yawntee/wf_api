package wf

type Channel interface {
	GetChannelNo() string
	GetMedia() string
	Login(usr, pwd string) (*GameUser, error)
	SendOtp(phone string) error
	OtpLogin(phone, code string) (*GameUser, error)
}

type ChannelType int

const (
	LEITING ChannelType = iota
	BILIBILI
)
