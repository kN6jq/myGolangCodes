package some_function

import "time"

// 时间
func TimeMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// 时间格式化
// nowTime := time.Now().Format("20060102150405")
