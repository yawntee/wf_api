package wf

import "fmt"

func UserAgent(device *Device) string {
	return fmt.Sprintf(
		"Dalvik/2.1.0 (Linux; U; Android %s; %s Build/%s)",
		Config.OsVer,
		Config.DeviceName,
		device.DeviceBuild,
	)
}

func Serial(device *Device) string {
	return fmt.Sprintf(
		"%s|%s|%s|%s",
		device.Mac,
		device.Imei,
		"unknown",
		device.AndroidId,
	)
}
