package some_function

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

// 获取字符串的md5 hash值
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// 字符串转MD5
func NewMd5(str ...string) string {
	h := md5.New()
	for _, v := range str {
		h.Write([]byte(v))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// 获取字节的md5 hash值
func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return hex.EncodeToString(ret[:])
}

// 获取字符串md5
func MD5(s string) string {
	return MD5Bytes([]byte(s))
}

// 获取文件md5
func MD5File(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return MD5Bytes(data), nil
}
