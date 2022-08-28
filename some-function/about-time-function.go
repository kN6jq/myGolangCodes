package some_function

import (
	"fmt"
	"time"
)

var gotime string = "2006-01-02 15:04:05"

// 时间
func TimeMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// 时间格式化
// nowTime := time.Now().Format("20060102150405")

func time1() {
	// 时间计算
	startime := time.Now()
	elapsed := time.Since(startime)
	fmt.Println(elapsed)
}
