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
	AndroidId   string
	Oaid        string
	Mac         string
	Imei        string
	DeviceId    string
	DeviceBuild string
	DeviceToken string
}

func NewDevice() *Device {
	mac := [6]string{}
	for i := 0; i < 6; i++ {
		mac[i] = randHex(2)
	}
	deviceId := rand.Intn(256) +
		rand.Intn(256)*(1<<8) +
		rand.Intn(256)*(1<<16) +
		rand.Intn(256)*(1<<24) +
		rand.Intn(1024)*(1<<32)
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
	rand.Seed(time.Now().Unix())
	return randNum(len, 16)
}

func randNum(len int, radix int) string {
	rand.Seed(time.Now().Unix())
	str := make([]byte, len)
	for i := 0; i < len; i++ {
		str[i] = strconv.FormatInt(int64(rand.Intn(radix)), radix)[0]
	}
	return string(str)
}
