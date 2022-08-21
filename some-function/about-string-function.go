package some_function

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

// 列表转字符串
func SliceSToString(slices [][]string) (result string) {
	b, err := json.Marshal(slices)
	if err != nil {
		return
	}
	result = string(b)
	return
}

// 获取字符串的md5 hash值
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
