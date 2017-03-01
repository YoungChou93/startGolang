package page_analyzer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
	"strings"
)

type DefaultAnalyer struct {
}

func NewDefaultAnalyer() *DefaultAnalyer {
	return &DefaultAnalyer{}
}

func (this *DefaultAnalyer) Analyze(page *page.Page) {
	fmt.Println("开始分析")
	quety := page.GetHtmlParser()
	quety.Find("a").Each(func(i int, contentSelection *goquery.Selection) {
		a := contentSelection
		url, _ := a.Attr("href")
		if strings.Contains(url,".html")&& strings.Contains(url,"http://") {
			fmt.Println(url)
			page.AddTargetRequest(url)
		}
	})
	fmt.Println("分析结束")
}
func (this *DefaultAnalyer) AnalyzeFinsh() {


}