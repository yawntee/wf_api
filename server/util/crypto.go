package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

var (
	_sha1 = sha1.New()
	_md5  = md5.New()
)

func Sha1(data []byte) string {
	return digest(_sha1, data)
}

func Md5(data []byte) string {
	return digest(_md5, data)
}

func digest(algorithm hash.Hash, data []byte) string {
	algorithm.Reset()
	algorithm.Write(data)
	return hex.EncodeToString(algorithm.Sum(nil))
}

type Cipher struct {
	Enc func([]byte) []byte
	Dec func([]byte) []byte
}

func NewCipher(init func([]byte) (cipher.Block, error), key []byte, iv []byte) *Cipher {
	var cip Cipher
	blk, err := init(key)
	if err != nil {
		panic(err)
	}
	if iv != nil {
		enc := cipher.NewCBCEncrypter(blk, iv)
		cip.Enc = func(data []byte) []byte {
			content := PKCS5Padding(data, blk.BlockSize())
			buf := make([]byte, len(content))
			enc.CryptBlocks(buf, content)
			return buf
		}
		dec := cipher.NewCBCDecrypter(blk, iv)
		cip.Dec = func(data []byte) []byte {
			buf := make([]byte, len(data))
			dec.CryptBlocks(buf, data)
			return PKCS5Trimming(buf)
		}
	} else {
		cip.Enc = func(data []byte) []byte {
			content := PKCS5Padding(data, blk.BlockSize())
			buf := make([]byte, len(content))
			for start := 0; start < len(content); start += blk.BlockSize() {
				end := start + blk.BlockSize()
				blk.Encrypt(buf[start:end], content[start:end])
			}
			return buf
		}
		cip.Dec = func(data []byte) []byte {
			buf := make([]byte, len(data))
			for start := 0; start < len(data); start += blk.BlockSize() {
				end := start + blk.BlockSize()
				blk.Decrypt(buf[start:end], data[start:end])
			}
			return PKCS5Trimming(buf)
		}
	}
	return &cip
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
