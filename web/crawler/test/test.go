package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"log"
)

func Spy(url string,number int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)")
	client := http.DefaultClient
	res, e := client.Do(req)
	if e != nil {
		fmt.Printf("Get请求%s返回错误:%s", url, e)
		return
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		fmt.Print(number)
		fmt.Println(":"+resStr[:50])
	}else{
		fmt.Print("fail:")
		fmt.Println(res.StatusCode)
	}
}

func main(){
	for i:=0;i<20;i++ {
		Spy("http://zhouyangalan.com.cn",i)
	}
}