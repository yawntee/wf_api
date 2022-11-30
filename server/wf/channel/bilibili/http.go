package bilibili

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"wf_api/server/wf/internal"
)

const (
	CODE_SUCCESS = 0
)

var agentHeader = func(r *http.Request) {
	r.Header.Set("User-Agent", userAgent)
}

func (c *Channel) getPwdKey() (hash string, cipherKey *rsa.PublicKey, err error) {
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
		"cipher_type":         {"bili_login_rsa"},
		"version_code":        {versionCode()},
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
		Hash      string `json:"hash"`
		CipherKey string `json:"cipher_key"`
	}
	err = json.NewDecoder(internal.PostForm(originalDomain+"/api/external/issue/cipher/v3", signForm(body, false), agentHeader)).Decode(&resp)
	if err != nil {
		panic(err)
	}
	if resp.Code != CODE_SUCCESS {
		return "", nil, fmt.Errorf("密钥获取失败(%s)", resp.Message)
	}
	p, _ := pem.Decode([]byte(resp.CipherKey))
	key, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		panic(err)
	}
	return resp.Hash, key.(*rsa.PublicKey), nil
}

func (c *Channel) fetchGameUser(device *internal.Device, uid uint64, username, token string) (*internal.GameUser, error) {
	sid := strconv.FormatUint(uid, 10)
	body := url.Values{
		"game":       {internal.GlobalConfig.Game},
		"channelNo":  {c.GetChannelNo()},
		"merchantId": {merchantId},
		"userId":     {sid},
		"token":      {token},
	}
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Uid     uint64 `json:"uid"`
	}
	err := json.NewDecoder(internal.PostForm("https://paywfauth.leiting.com/terrace/phone_login_verify!biligameVerifySession.action", body, func(req *http.Request) {
		req.Header.Set("User-Agent", internal.UserAgent(device))
	})).Decode(&resp)
	if err != nil {
		panic(err)
	}
	if resp.Status != "success" {
		return nil, ErrLoginFailed
	}
	var Msg struct {
		Token string `json:"token"`
	}
	err = json.Unmarshal([]byte(resp.Message), &Msg)
	if err != nil {
		panic(err)
	}
	return &internal.GameUser{
		Uid:      uid,
		Sid:      sid,
		Username: username,
		Token:    Msg.Token,
	}, nil
}
