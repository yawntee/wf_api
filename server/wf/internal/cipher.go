package internal

import (
	"crypto/aes"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"wf_api/server/util"
)

type cipherType int

const (
	Text = iota
	Base64
	Hex
)

type cipher struct {
	util.Cipher
	EncType cipherType
	DecType cipherType
}

func _in(data []byte, _type cipherType) []byte {
	switch _type {
	case Text:
		return data
	case Hex:
		bytes, err := hex.DecodeString(string(data))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return bytes
	case Base64:
		bytes, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return bytes
	default:
		fmt.Println("无效的加密数据类型")
		return nil
	}
}

func _out(data []byte, _type cipherType) []byte {
	switch _type {
	case Text:
		return data
	case Hex:
		return []byte(hex.EncodeToString(data))
	case Base64:
		return []byte(base64.StdEncoding.EncodeToString(data))
	default:
		fmt.Println("无效的加密数据类型")
		return nil
	}
}

func (c cipher) Enc(data []byte) []byte {
	return _out(c.Cipher.Enc(_in(data, c.EncType)), c.DecType)
}

func (c cipher) Dec(data []byte) []byte {
	return _out(c.Cipher.Dec(_in(data, c.DecType)), c.EncType)
}

var (
	DataCipher = cipher{
		Cipher:  *util.NewCipher(aes.NewCipher, []byte("#LeitingAESKey#!"), nil),
		EncType: Text,
		DecType: Base64,
	}
	LoginCipher = cipher{
		Cipher:  *util.NewCipher(aes.NewCipher, []byte("#LeitingAESKey#!"), []byte("LeitingAESIVKEY!")),
		EncType: Text,
		DecType: Base64,
	}
	PwdCipher = cipher{
		Cipher:  *util.NewCipher(des.NewCipher, []byte("leiting\000"), nil),
		EncType: Text,
		DecType: Hex,
	}
	//ConfigCipher = cipher{
	//	Cipher:  *util.NewCipher(des.NewCipher, []byte("548711fd"), nil),
	//	EncType: Text,
	//	DecType: Hex,
	//}
)
