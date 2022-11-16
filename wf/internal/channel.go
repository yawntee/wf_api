package internal

type Channel interface {
	GetChannelNo() string
	GetMedia() string
	Login(usr, pwd string) (*GameUser, error)
	SendOtp(phone string) error
	OtpLogin(phone, code string) (*GameUser, error)
}
