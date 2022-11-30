package bilibili

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"golang.org/x/exp/maps"
	"net/url"
	"sort"
	"strings"
	"wf_api/server/util"
)

const (
	appKey = "9618d7d201064f80a01329dfbc4da9e4"
	rsaKey = "\x30\x81\x9f\x30\x0d\x06\x09\x2a\x86\x48\x86\xf7\x0d\x01\x01\x01\x05\x00\x03\x81\x8d\x00\x30\x81\x89\x02\x81\x81\x00\x9b\xd2\x51\x71\x0f\x53\x89\x5f\xeb\x03\xbc\x6d\x67\x9f\xe4\x06\x8b\xde\x70\x21\x13\xfb\xd5\x8d\x2b\xda\x3b\x38\xec\xd8\x13\x99\x89\xcc\x28\xab\x70\x55\xb4\x5d\x72\x9b\xdc\xfc\xc6\x1e\x39\x03\xf5\x39\x40\x94\x71\x00\x81\x25\xa2\x96\xa5\x3b\x6c\xb8\x87\x85\x3e\x05\xd8\x94\x8a\xbb\x96\x8a\x0a\xab\x12\xf9\xc7\x30\x76\xaa\xe1\x00\x8a\x74\x55\x45\x35\x9a\x3e\x90\xf9\xbf\x44\x5e\xe2\x7e\x1b\xe9\x26\x66\x12\xbe\x01\x8d\x58\x36\x42\x15\x6b\x54\x4f\x20\x1d\x69\x05\x42\x54\x59\xc1\xfe\x5f\xe9\x3a\xeb\x91\xd3\xfe\x05\x02\x03\x01\x00\x01"
)

func salt(str string) string {
	var builder strings.Builder
	builder.Write([]byte{str[2], str[12], str[22]})
	builder.WriteString(str)
	return builder.String()
}

func sign(str string) string {
	return util.Md5([]byte(str + appKey))
}

func signForm(form url.Values, feign bool) url.Values {
	keys := maps.Keys(form)
	sort.Strings(keys)
	var builder strings.Builder
	for _, k := range keys {
		builder.WriteString(form.Get(k))
	}
	sign := sign(builder.String())
	form.Set("sign", sign)
	if feign {
		key, err := x509.ParsePKIXPublicKey([]byte(rsaKey))
		if err != nil {
			panic(err)
		}
		enc, err := rsa.EncryptPKCS1v15(rand.Reader, key.(*rsa.PublicKey), []byte(sign))
		if err != nil {
			panic(err)
		}
		form.Set("feign_sign", base64.StdEncoding.EncodeToString(enc))
	}
	return form
}
