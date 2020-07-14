package rand

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

// Base64Encode 返回一个base64编码后的数据
func Base64Encode(raw string) string {
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

// Base64Decode 返回一个base64解码后的数据 解码失败将返回原数据
func Base64Decode(raw string) string {
	c, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return raw
	}
	return string(c)
}

// Md5Bits32 返回一个32位的md5数据
func Md5Bits32(raw string) string {
	h := md5.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1Bits40 返回一个40位的sha1数据
func Sha1Bits40(raw string) string {
	h := sha1.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256Bits64 返回一个64位的sha256数据
func Sha256Bits64(raw string) string {
	h := sha256.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

// RandomLetterLowercase 返回一个随机只有小写字母的给定长度的字符串
func RandomLetterLowercase(length uint16) string {
	const char = "abcdefghijklmnopqrstuvwxyz"
	var source = rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var s bytes.Buffer
	var i uint16
	for i = 0; i < length; i++ {
		s.WriteByte(char[r.Int63()%int64(len(char))])
	}
	return s.String()
}

// RandomLetterUppercase 返回一个随机只有大写字母的给定长度的字符串
func RandomLetterUppercase(length uint16) string {
	const char = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var source = rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var s bytes.Buffer
	var i uint16
	for i = 0; i < length; i++ {
		s.WriteByte(char[r.Int63()%int64(len(char))])
	}
	return s.String()
}

// RandomMixingLetterLowercaseAndNumbers 返回字母和数字混合小写
func RandomMixingLetterLowercaseAndNumbers(length uint16) string {
	const char = "abcdefghijklmnopqrstuvwxyz0123456789"
	var source = rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var s bytes.Buffer
	var i uint16
	for i = 0; i < length; i++ {
		s.WriteByte(char[r.Int63()%int64(len(char))])
	}
	return s.String()
}

// RandomMixingLettersAndNumbers 返回字母和数字混合大小写
func RandomMixingLettersAndNumbers(length uint16) string {
	const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var source = rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var s bytes.Buffer
	var i uint16
	for i = 0; i < length; i++ {
		s.WriteByte(char[r.Int63()%int64(len(char))])
	}
	return s.String()
}

// CustomRandomMixingLetters 返回自定义随机字符 第二个参数为自定义字符集(英文)
// eg. CustomRandomMixingLetters(16, "~`!@#$%^&*()_+|-=\\[]{};':\",./<>?")
func CustomRandomMixingLetters(length uint16, args string) string {
	var char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + args
	var source = rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var s bytes.Buffer
	var i uint16
	for i = 0; i < length; i++ {
		s.WriteByte(char[r.Int63()%int64(len(char))])
	}
	return s.String()
}
