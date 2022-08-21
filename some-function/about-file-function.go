package some_function

import (
	"os"
)

// 对文件内容进行追加
func WirteFileAppend(fileName string, newSubDomain []string) {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	for _, st := range newSubDomain {
		buf := []byte(st + "\n")
		fd.Write(buf)
	}
	fd.Close()
}
