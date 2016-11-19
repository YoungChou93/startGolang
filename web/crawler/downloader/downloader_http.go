package downloader

import (
	"net/http"
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
	"io/ioutil"
)

type HttpDownloader struct {

}

func NewHttpDownloader()*HttpDownloader{
     return &HttpDownloader{}
}

func (this *HttpDownloader)Download(request *crawler_request.CrawlerRequest)*page.Page{

	page:=page.NewPage(request)

	httpRequest, _ := http.NewRequest("GET", request.GetUrl(), nil)
	httpRequest.Header.Set("User-Agent", "Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)")
	client := http.DefaultClient
	res, e := client.Do(httpRequest)
	if e != nil {
		page.SetStatus(false,"error")
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		page.SetHeader(res.Header)
		page.SetCookies(res.Cookies())
		page.SetBodyStr(resStr)
	}
	return page
}
