package Concurrent

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(ChannelStart())
	fmt.Println(NomalStart())
	WaitGroup()
}

func WaitGroup() {

	startime := time.Now()
	wg.Add(1)
	for i := 0; i < 7; i++ {
		go func(i int) {
			defer wg.Done()
			Spider1(i, nil)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(startime)
	fmt.Println(elapsed)
}

func ChannelStart() time.Duration {
	startime := time.Now()
	ch := make(chan bool)
	for i := 0; i < 7; i++ {
		go Spider1(i, ch)
	}
	for i := 0; i < 7; i++ {
		<-ch
	}
	elapsed := time.Since(startime)
	return elapsed
}

func NomalStart() time.Duration {
	startime := time.Now()
	for i := 0; i < 7; i++ {
		Spider(i)
	}
	elapsed := time.Since(startime)
	return elapsed
}

func Spider(page int) {
	// 1. 发送信息
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://www.shenmezhidedu.com/category/shudan/page/"+strconv.Itoa(page), nil)
	if err != nil {
		fmt.Println("req error ", err)
	}
	req.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("Accept-Encoding", " gzip, deflate, br")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Cookie", " ll=\"108090\"; bid=Y4v5uOQFXDM; __utma=30149280.718263598.1661682708.1661682708.1661682708.1; __utmc=30149280; __utmz=30149280.1661682708.1.1.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmt=1; __utmb=30149280.1.10.1661682708; ap_v=0,6.0; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1661682715%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=3f38537a9e744537.1661682715.1.1661682715.1661682715.; _pk_ses.100001.4cf6=*; __utma=223695111.1778217804.1661682715.1661682715.1661682715.1; __utmb=223695111.0.10.1661682715; __utmc=223695111; __utmz=223695111.1661682715.1.1.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=D3F80C43991A6C6EC95CACA624C6C5E08|8a4fd6f71a2ea7978f13e69da411a6dc")
	req.Header.Add("Sec-Fetch-Dest", " document")
	req.Header.Add("Sec-Fetch-Mode", " navigate")
	req.Header.Add("Sec-Fetch-Site", " none")
	req.Header.Add("Sec-Fetch-User", " ?1")
	req.Header.Add("Upgrade-Insecure-Requests", " 1")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", " ")
	req.Header.Add("sec-ch-ua-mobile", " ?0")
	req.Header.Add("sec-ch-ua-platform", "Windows")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp error ", err)
	}
	defer resp.Body.Close()
	// 2. 解析网页
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("docDetail error ", err)
	}
	//#content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 >
	// #content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 >
	// #content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 > div:nth-child(2) > div > div.list-content > div.list-body > a
	// #content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 > div:nth-child(1) > div > div.list-content > div.list-footer > div > a.d-none.d-md-inline-block.text-xs
	docDetail.Find("#content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3").Each(
		func(i int, s *goquery.Selection) {
			for j := 1; j < 25; j++ {
				title := s.Find("div:nth-child(" + strconv.Itoa(j) + ") > div > div.list-content > div.list-body > a").Text()
				author := s.Find("div:nth-child(" + strconv.Itoa(j) + ") > div > div.list-content > div.list-footer > div > a.d-none.d-md-inline-block.text-xs").Text()
				like := s.Find("div:nth-child(" + strconv.Itoa(j) + ") > div > div.list-content > div.list-footer > div > div.text-muted > a:nth-child(2) > small").Text()

				fmt.Println("title: ", strings.TrimSpace(strings.Trim(title, "\n")))
				fmt.Println("author: ", author)
				fmt.Println("like: ", like)
			}
		})

	// 4. 保存信息
}
func Spider1(page int, ch chan bool) {
	// 1. 发送信息
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://www.shenmezhidedu.com/category/shudan/page/"+strconv.Itoa(page), nil)
	if err != nil {
		fmt.Println("req error ", err)
	}
	req.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("Accept-Encoding", " gzip, deflate, br")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Cookie", " ll=\"108090\"; bid=Y4v5uOQFXDM; __utma=30149280.718263598.1661682708.1661682708.1661682708.1; __utmc=30149280; __utmz=30149280.1661682708.1.1.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmt=1; __utmb=30149280.1.10.1661682708; ap_v=0,6.0; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1661682715%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=3f38537a9e744537.1661682715.1.1661682715.1661682715.; _pk_ses.100001.4cf6=*; __utma=223695111.1778217804.1661682715.1661682715.1661682715.1; __utmb=223695111.0.10.1661682715; __utmc=223695111; __utmz=223695111.1661682715.1.1.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=D3F80C43991A6C6EC95CACA624C6C5E08|8a4fd6f71a2ea7978f13e69da411a6dc")
	req.Header.Add("Sec-Fetch-Dest", " document")
	req.Header.Add("Sec-Fetch-Mode", " navigate")
	req.Header.Add("Sec-Fetch-Site", " none")
	req.Header.Add("Sec-Fetch-User", " ?1")
	req.Header.Add("Upgrade-Insecure-Requests", " 1")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", " ")
	req.Header.Add("sec-ch-ua-mobile", " ?0")
	req.Header.Add("sec-ch-ua-platform", "Windows")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp error ", err)
	}
	defer resp.Body.Close()
	// 2. 解析网页
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("docDetail error ", err)
	}
	//#content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 >
	// #content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 >
	// #content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 > div:nth-child(2) > div > div.list-content > div.list-body > a
	// #content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3 > div:nth-child(1) > div > div.list-content > div.list-footer > div > a.d-none.d-md-inline-block.text-xs
	docDetail.Find("#content > div > div > div > div.list-archive.list-grouped.row.my-n2.my-md-n3").Each(
		func(i int, s *goquery.Selection) {
			for j := 1; j < 25; j++ {
				title := s.Find("div:nth-child(" + strconv.Itoa(j) + ") > div > div.list-content > div.list-body > a").Text()
				author := s.Find("div:nth-child(" + strconv.Itoa(j) + ") > div > div.list-content > div.list-footer > div > a.d-none.d-md-inline-block.text-xs").Text()
				like := s.Find("div:nth-child(" + strconv.Itoa(j) + ") > div > div.list-content > div.list-footer > div > div.text-muted > a:nth-child(2) > small").Text()

				fmt.Println("title: ", strings.TrimSpace(strings.Trim(title, "\n")))
				fmt.Println("author: ", author)
				fmt.Println("like: ", like)
			}
		})
	if ch != nil {
		ch <- true
	}
	// 4. 保存信息
}
