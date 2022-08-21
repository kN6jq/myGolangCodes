package some_function

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func LoadFile(filename string) (lines []string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close() //nolint
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return
}

// 路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 读取文件的前n个字节
func ReadFileN(path string, n int) ([]byte, error) {
	data := make([]byte, n)

	if !IsFile(path) {
		wrappedErr := fmt.Errorf("File %s not found", path)
		return nil, wrappedErr
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		wrappedErr := fmt.Errorf("Open file %s error", path)
		return nil, wrappedErr
	}

	file.Read(data)

	return data, nil
}

// 写入字节到文件
func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0)
}

// 写入字符串到文件
func WriteToFile(wireteString, filename string) {
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(wireteString)
	fd.Write(buf)

}
