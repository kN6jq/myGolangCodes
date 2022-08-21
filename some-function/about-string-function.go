package some_function

import (
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
