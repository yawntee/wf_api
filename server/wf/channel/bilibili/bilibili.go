package bilibili

import (
	crand "crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/url"
	"strings"
	"wf_api/server/util"
	"wf_api/server/wf/internal"
)

var (
	ErrLoginFailed = errors.New("登录失败")
)

type Channel struct {
	Buvid      string `json:"buvid"`
	Udid       string `json:"udid"`
	Bdid       string `json:"bdid"`
	AccessKey  string `json:"accessKey"`
	captchaKey string
}

func NewChannel(device *internal.Device) *Channel {
	var c Channel
	if device != nil {
		c.Buvid = "XY" + salt(strings.ToUpper(util.Md5([]byte(strings.ToLower(device.Mac)))))
		udid := []byte(strings.ToLower(strings.ReplaceAll(device.Mac, ":", "")) + "|||")
		udid[0] = udid[0] ^ (byte(len(udid)) & 0xff)
		for i := 1; i < len(udid); i++ {
			udid[i] = (udid[i] ^ udid[i-1]) & 0xff
		}
		c.Udid = base64.StdEncoding.EncodeToString(udid)
		c.Bdid = uuid.NewString() + uuid.NewString()
		if len(c.Bdid) > 64 {
			c.Bdid = c.Bdid[:64]
		}
	}
	return &c
}

func (c *Channel) GetChannelNo() string {
	return "130061"
}

func (c *Channel) GetMedia() string {
	return "130061"
}

func (c *Channel) Login(device *internal.Device, usr, pwd string) (*internal.GameUser, error) {
	hash, cipherKey, err := c.getPwdKey()
	if err != nil {
		return nil, err
	}
	encPwd, err := rsa.EncryptPKCS1v15(crand.Reader, cipherKey, []byte(hash+pwd))
	if err != nil {
		panic(err)
	}
	body := url.Values{
		"cur_buvid":           {c.Buvid},
		"sdk_type":            {sdkType},
		"merchant_id":         {merchantId},
		"platform":            {platform},
		"apk_sign":            {apkSign},
		"platform_type":       {platform},
		"old_buvid":           {c.Buvid},
		"udid":                {c.Udid},
		"app_id":              {appId},
		"game_id":             {appId},
		"timestamp":           {timestamp()},
		"version_code":        {versionCode()},
		"bd_id":               {c.Bdid},
		"server_id":           {serverId},
		"version":             {version},
		"domain_switch_count": {domainSwitchCount},
		"country_code":        {countryCode},
		"app_ver":             {internal.GlobalConfig.VersionName},
		"domain":              {domain},
		"original_domain":     {originalDomain},
		"user_id":             {usr},
		"sdk_log_type":        {sdkLogType},
		"current_env":         {currentEnv},
		"sdk_ver":             {sdkVer},
		"pwd":                 {base64.StdEncoding.EncodeToString(encPwd)},
		"channel_id":          {channelId},
	}
	var resp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		AccessKey string `json:"access_key"`
		Uid       uint64 `json:"uid"`
	}
	err = json.NewDecoder(internal.PostForm(originalDomain+"/api/external/login/v3", signForm(body, false), agentHeader)).Decode(&resp)
	if err != nil {
		panic(err)
	}
	if resp.Code != CODE_SUCCESS {
		return nil, fmt.Errorf("%s", resp.Message)
	}
	c.AccessKey = resp.AccessKey
	user, err := c.fetchGameUser(device, resp.Uid, usr, resp.AccessKey)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Channel) SendOtp(device *internal.Device, phone string) error {
	body := url.Values{
		"cur_buvid":            {c.Buvid},
		"sdk_type":             {sdkType},
		"merchant_id":          {merchantId},
		"otp_type":             {"login"},
		"platform":             {platform},
		"apk_sign":             {apkSign},
		"platform_type":        {platform},
		"old_buvid":            {c.Buvid},
		"otp_channel_no":       {phone},
		"udid":                 {c.Udid},
		"app_id":               {appId},
		"game_id":              {appId},
		"timestamp":            {timestamp()},
		"otp_channel_category": {"tel"},
		"version_code":         {versionCode()},
		"bd_id":                {c.Bdid},
		"server_id":            {serverId},
		"version":              {version},
		"domain_switch_count":  {domainSwitchCount},
		"country_code":         {countryCode},
		"app_ver":              {internal.GlobalConfig.VersionName},
		"domain":               {domain},
		"original_domain":      {originalDomain},
		"sdk_log_type":         {sdkLogType},
		"current_env":          {currentEnv},
		"sdk_ver":              {sdkVer},
		"channel_id":           {channelId},
	}
	var resp struct {
		Code       int    `json:"code"`
		Message    string `json:"message"`
		CaptchaKey string `json:"captcha_key"`
	}
	err := json.NewDecoder(internal.PostForm(originalDomain+"/api/external/otp/send/v3", signForm(body, false), agentHeader)).Decode(&resp)
	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%s", resp.Message)
	}
	c.captchaKey = resp.CaptchaKey
	return nil
}

func (c *Channel) OtpLogin(device *internal.Device, phone, code string) (*internal.GameUser, error) {
	body := url.Values{
		"cur_buvid":           {c.Buvid},
		"sdk_type":            {sdkType},
		"merchant_id":         {merchantId},
		"platform":            {platform},
		"apk_sign":            {apkSign},
		"platform_type":       {platform},
		"old_buvid":           {c.Buvid},
		"captcha_key":         {c.captchaKey},
		"udid":                {c.Udid},
		"app_id":              {appId},
		"game_id":             {appId},
		"timestamp":           {timestamp()},
		"version_code":        {versionCode()},
		"mobile":              {phone},
		"otp":                 {code},
		"bd_id":               {c.Bdid},
		"server_id":           {serverId},
		"version":             {version},
		"domain_switch_count": {domainSwitchCount},
		"country_code":        {countryCode},
		"app_ver":             {internal.GlobalConfig.VersionName},
		"domain":              {domain},
		"original_domain":     {originalDomain},
		"sdk_log_type":        {sdkLogType},
		"current_env":         {currentEnv},
		"sdk_ver":             {sdkVer},
		"channel_id":          {channelId},
	}
	var resp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		AccessKey string `json:"access_key"`
		Uid       uint64 `json:"uid"`
	}
	err := json.NewDecoder(internal.PostForm(originalDomain+"/api/external/login/otp/v3", signForm(body, true), agentHeader)).Decode(&resp)
	if err != nil {
		panic(err)
	}
	if resp.Code != CODE_SUCCESS {
		return nil, fmt.Errorf("%s", resp.Message)
	}
	c.AccessKey = resp.AccessKey
	user, err := c.fetchGameUser(device, resp.Uid, phone, resp.AccessKey)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Channel) CheckLogin(device *internal.Device, user *internal.GameUser) error {
	bdInfo := encBdInfo(c.bdInfo(device))
	body := url.Values{
		"cur_buvid":           {c.Buvid},
		"sdk_type":            {sdkType},
		"merchant_id":         {merchantId},
		"platform":            {platform},
		"apk_sign":            {apkSign},
		"platform_type":       {platform},
		"old_buvid":           {c.Buvid},
		"udid":                {c.Udid},
		"app_id":              {appId},
		"game_id":             {appId},
		"timestamp":           {timestamp()},
		"version_code":        {versionCode()},
		"bd_id":               {c.Bdid},
		"server_id":           {serverId},
		"version":             {version},
		"domain_switch_count": {domainSwitchCount},
		"app_ver":             {internal.GlobalConfig.VersionName},
		"domain":              {domain},
		"access_key":          {c.AccessKey},
		"bd_info":             {base64.StdEncoding.EncodeToString(bdInfo)},
		"original_domain":     {originalDomain},
		"sdk_log_type":        {sdkLogType},
		"current_env":         {currentEnv},
		"sdk_ver":             {sdkVer},
		"channel_id":          {channelId},
	}
	var resp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		AccessKey string `json:"access_key"`
		Uid       uint64 `json:"uid"`
	}
	err := json.NewDecoder(internal.PostForm(originalDomain+"/api/external/user.token.oauth.login/v3", signForm(body, false), agentHeader)).Decode(&resp)
	if err != nil {
		panic(err)
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%s", resp.Message)
	}
	c.AccessKey = resp.AccessKey
	newUser, err := c.fetchGameUser(device, resp.Uid, user.Username, resp.AccessKey)
	if err != nil {
		return err
	}
	*user = *newUser
	return nil
}
