package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	lock sync.Mutex
)

type XrayResAll struct {
	VerboseName string        `json:"verbose_name"`
	Parent      string        `json:"parent"`
	Domain      string        `json:"domain"`
	Cname       []string      `json:"cname"`
	IP          []IP          `json:"ip"`
	Web         []interface{} `json:"web"`
	Extra       []Extra       `json:"extra"`
}
type IP struct {
	IP      string `json:"ip"`
	Asn     string `json:"asn"`
	Country string `json:"country"`
}
type Extra struct {
	Source string `json:"source"`
	Detail string `json:"detail"`
}

func main() {
	Read2()

}

func Read2() {
	//获得一个file
	f, err := os.Open("dxy.cn.json")
	if err != nil {
		fmt.Println("read fail")
	}

	//把file读取到缓冲区中
	defer f.Close()
	var xrayResult XrayResAll
	buf := bufio.NewReader(f)

	for {
		//从file读取到buf中
		lines, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		// 过滤存在[  ]  等字符
		if len(lines) < 10 {

		} else {
			line := strings.TrimRight(lines, ",\n")
			//fmt.Println(line)
			err = json.Unmarshal([]byte(line), &xrayResult)
			if err != nil {
				panic(err)
			}
			if len(xrayResult.IP) != 0 {
				fmt.Printf("Source: %s, Subdomain: %s\n ", xrayResult.VerboseName, xrayResult.Domain)
			}
		}
	}
}
