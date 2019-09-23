package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	//	"log"
	"net/http"
	//"unicode"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"

	//	"golang.org/x/text/unicode"

	"golang.org/x/net/html/charset"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	/*
		resp, err := http.Get(url)

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	*/
	client := &http.Client{}

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	///	reqest.Header.Add("Accept", "text/html")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	//	reqest.Header.Add("Accept-Encoding", "gzip, deflate")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	//转码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(
	r io.Reader) encoding.Encoding {
	bytes, _ := bufio.NewReader(r).Peek(1024)

	/*if err != nil {
		log.Printf("Fetcher error :%v", err)
		return unicode.ASCII_Hex_Digit
	}
	*/
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
