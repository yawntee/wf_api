package bilibili

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"wf_api/wf/internal"
)

func (c *Channel) bdInfo(device *internal.Device) []byte {
	marshal, err := json.Marshal(map[string]any{
		"udid":          c.Udid,
		"bd_id":         c.Bdid,
		"cur_buvid":     c.Buvid,
		"old_buvid":     c.Buvid,
		"imei":          "",
		"mac":           device.Mac,
		"android_id":    device.AndroidId,
		"oaid":          device.Oaid,
		"model":         internal.GlobalConfig.DeviceName,
		"brand":         "Redmi",
		"dp":            "1080,1920,440",
		"net":           "4",
		"operators":     "中国移动",
		"supportedAbis": "[arm64-v8a : armeabi-v7a : armeabi]",
		"is_root":       "0",
		"pf_ver":        internal.GlobalConfig.OsVer,
		"platform_type": "Android",
		"ver":           internal.GlobalConfig.VersionName,
		"version_code":  versionCode(),
		"sdk_ver":       sdkVer,
		"app_id":        appId,
		"fts":           0,
		"first":         1,
		"files":         "/data/user/0/com.leiting.wf.bilibili/files",
		"pkg_name":      "com.leiting.wf.bilibili",
		"app_name":      "世界弹射物语",
		"finger_print":  "Redmi/Phoenix/Phoenix:12/${device.deviceBuild}/12.0.8.0:user/release-keys",
		"serial":        "unknown",
		"band":          "MPSS.HI.2.0.c7-00236-0510_2340_044fee7,MPSS.HI.2.0.c7-00236-0510_2340_044fee7",
		"cpu_count":     8,
		"cpu_model":     "AArch64 Processor rev 3 (aarch64,",
		"cpu_freq":      1804800,
		"cpu_verdor":    "Qualcomm",
		"sensor":        "{\"18\":\"pedometer  Non-wakeup\",\"19\":\"pedometer  Non-wakeup\",\"29\":\"stationary_detect\",\"17\":\"sns_smd  Wakeup\",\"22\":\"ccd_tilt  Wakeup\",\"11\":\"Rotation Vector  Non-wakeup\",\"3\":\"orientation  Non-wakeup\",\"33171027\":\"NonUi  Non-wakeup\",\"33171036\":\"pickup  Non-wakeup\",\"33171039\":\"oem13_light_smd  Non-wakeup\",\"33171032\":\"3d_signature  Non-wakeup\",\"30\":\"motion_detect\",\"33171007\":\"rohm_bu27030  Non-wakeup\",\"14\":\"ak0991x Magnetometer-Uncalibrated Non-wakeup\",\"2\":\"ak0991x Magnetometer Non-wakeup\",\"10\":\"linear_acceleration\",\"33171038\":\"knuckle  Non-wakeup\",\"9\":\"gravity  Non-wakeup\",\"15\":\"Game Rotation Vector  Non-wakeup\",\"4\":\"lsm6dso Gyroscope Non-wakeup\",\"33171029\":\"Aod  Non-wakeup\",\"27\":\"device_orient  Non-wakeup\",\"20\":\"sns_geomag_rv  Non-wakeup\",\"33171057\":\"rohm_bu27030_back  Non-wakeup\",\"5\":\"rohm_bu27030 Ambient Light Sensor Non-wakeup\",\"33171055\":\"rohm_bu27030_back  Non-wakeup\",\"35\":\"lsm6dso Accelerometer-Uncalibrated Non-wakeup\",\"16\":\"lsm6dso Gyroscope-Uncalibrated Non-wakeup\",\"33171030\":\"fod  Non-wakeup\",\"1\":\"lsm6dso Accelerometer Non-wakeup\",\"33171031\":\"Touch Sensor\",\"8\":\"Xiaomi Proximity\"}",
		"camcnt":        2,
		"campx":         "4640x2160",
		"camzoom":       "5.31",
		"bat_level":     "77",
		"bat_state":     "2",
		"ts":            timestamp(),
		"brightness":    238,
		"boot":          rand.Uint32(),
		"total_ram":     8589934592,
		"total_rom":     245760645480,
		"is_debug":      "1",
		"is_emu":        "000000",
		"time_zone":     "GMT+08:00 TimeZone id :Asia/Shanghai",
		"lang":          "ZH",
		"os":            "android",
		"axposed":       "",
		"maps":          "",
		"virtual":       "",
		"kernel_ver":    "4.19.112-perf-gdbf312947796",
	})
	if err != nil {
		panic(err)
	}
	return []byte(strings.ReplaceAll(string(marshal), "/", "\\/"))
}

func timestamp() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func versionCode() string {
	return strconv.Itoa(internal.GlobalConfig.VersionCode)
}
