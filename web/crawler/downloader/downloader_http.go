package downloader

import (
	"net/http"
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
	"io/ioutil"
	"fmt"
	"time"
	"math/rand"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

type HttpDownloader struct {

}

func NewHttpDownloader()*HttpDownloader{
     return &HttpDownloader{}
}

func (this *HttpDownloader)Download(request *crawler_request.CrawlerRequest)*page.Page{
	fmt.Println("开始下载")
	fmt.Println(request.Url)
	page:=page.NewPage(request)

	httpRequest, _ := http.NewRequest("GET", request.GetUrl(), nil)
	httpRequest.Header.Set("User-Agent",GetRandomUserAgent())
	client := http.DefaultClient
	res, e := client.Do(httpRequest)
	if e != nil {
		page.SetStatus(false,"error")
	}else {
		if res.StatusCode == 200 {
			body := res.Body
			defer body.Close()
			bodyByte, _ := ioutil.ReadAll(body)

			resStr := string(bodyByte)

			page.SetHeader(res.Header)
			page.SetCookies(res.Cookies())
			page.SetBodyStr(resStr)

			bodyReader := bytes.NewReader(bodyByte)
			var doc *goquery.Document
			doc, _ = goquery.NewDocumentFromReader(bodyReader);
			page.SetHtmlParser(doc)

		}
	}
	fmt.Println("下载结束")
	return page
}


var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}