package downloader

import (
	"net/http"
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"fmt"
	"io/ioutil"
)

type HttpDownloader struct {

}

func NewHttpDownloader()*HttpDownloader{
     return &HttpDownloader{}
}


func (this *HttpDownloader)Download(request *crawler_request.CrawlerRequest){

	httprequest, _ := http.NewRequest("GET", request.GetUrl(), nil)

	httprequest.Header.Set("User-Agent", "Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)")
	client := http.DefaultClient
	res, e := client.Do(httprequest)
	if e != nil {
		fmt.Printf("Get请求%s返回错误:%s", request.GetUrl(), e)
		return
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		fmt.Println(resStr[:50])
	}

}
