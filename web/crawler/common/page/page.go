
package page

import (
    "net/http"
    "github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "github.com/hu17889/go_spider/core/common/mlog"
)

type Page struct {

    success  bool
    errormsg string

    request *crawler_request.CrawlerRequest
    body string

    docParser *goquery.Document

    header  http.Header
    cookies []*http.Cookie

    nextRequests []*crawler_request.CrawlerRequest
}

func NewPage(request *crawler_request.CrawlerRequest)*Page{
    return &Page{request:request}

}

func (this *Page) SetHeader(header http.Header) {
    this.header = header
}

func (this *Page) GetHeader() http.Header {
    return this.header
}

func (this *Page) SetCookies(cookies []*http.Cookie) {
    this.cookies = cookies
}

func (this *Page) GetCookies() []*http.Cookie {
    return this.cookies
}

func (this *Page) IsSuccess() bool {
    return !this.success
}

func (this *Page) Errormsg() string {
    return this.errormsg
}

func (this *Page) SetStatus(success bool, errormsg string) {
    this.success = success
    this.errormsg = errormsg
}

func (this *Page) SetRequest(request *crawler_request.CrawlerRequest) *Page {
    this.request = request
    return this
}

func (this *Page) GetRequest() *crawler_request.CrawlerRequest {
    return this.request
}


func (this *Page) AddTargetRequest(url string) *Page {
    this.nextRequests = append(this.nextRequests, crawler_request.NewCrawlerRequest(url,"Get"))
    return this
}

func (this *Page) GetTargetRequests() []*crawler_request.CrawlerRequest {
    return this.nextRequests
}

func (this *Page) AddTargetRequests(urls []string, respType string) *Page {
    for _, url := range urls {
        this.AddTargetRequest(url)
    }
    return this
}

func (this *Page) SetBodyStr(body string) *Page {
    this.body = body
    return this
}


func (this *Page) GetBodyStr() string {
    return this.body
}

func (this *Page) SetHtmlParser(doc *goquery.Document) *Page {
    this.docParser = doc
    return this
}

func (this *Page) GetHtmlParser() *goquery.Document {
    return this.docParser
}


func (this *Page) ResetHtmlParser() *goquery.Document {
    r := strings.NewReader(this.body)
    var err error
    this.docParser, err = goquery.NewDocumentFromReader(r)
    if err != nil {
        mlog.LogInst().LogError(err.Error())
        panic(err.Error())
    }
    return this.docParser
}


