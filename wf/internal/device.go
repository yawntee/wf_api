package internal

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	deviceImei = "000000000000000"
)

type Device struct {
	AndroidId   string `json:"android_id"`
	Oaid        string `json:"oaid"`
	Mac         string `json:"mac"`
	Imei        string `json:"imei"`
	DeviceId    string `json:"device_id"`
	DeviceBuild string `json:"device_build"`
	DeviceToken string `json:"device_token"`
}

var rander = rand.New(rand.NewSource(time.Now().Unix()))

func NewDevice() *Device {
	mac := [6]string{}
	for i := 0; i < 6; i++ {
		mac[i] = randHex(2)
	}
	deviceId := rander.Intn(256) +
		rander.Intn(256)*(1<<8) +
		rander.Intn(256)*(1<<16) +
		rander.Intn(256)*(1<<24) +
		rander.Intn(1024)*(1<<32)
	return &Device{
		AndroidId:   randHex(16),
		Oaid:        randHex(16),
		Mac:         strings.ToUpper(strings.Join(mac[:], ":")),
		Imei:        deviceImei,
		DeviceId:    strconv.Itoa(deviceId),
		DeviceBuild: strings.ToUpper(randNum(6, 27)),
		DeviceToken: randHex(32),
	}
}

func randHex(len int) string {
	return randNum(len, 16)
}

func randNum(len int, radix int) string {
	str := make([]byte, len)
	for i := 0; i < len; i++ {
		str[i] = strconv.FormatInt(int64(rander.Intn(radix)), radix)[0]
	}
	return string(str)
}
