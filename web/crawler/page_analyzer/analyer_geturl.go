package page_analyzer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
	"strings"
	"os"
)

type GeturlAnalyer struct {
	urls []string
}

func NewGeturlAnalyer() *GeturlAnalyer {
	return &GeturlAnalyer{}
}

func (this *GeturlAnalyer) Analyze(page *page.Page) {
	fmt.Println("开始分析")
	number :=page.GetRequest().Label
	fmt.Println(number)
	quety := page.GetHtmlParser()
	bodystr := page.GetBodyStr()
	if strings.Contains(bodystr, "result-url") {
		resultselection := quety.Find(".result-url")
		if resultselection != nil {
			url :=""
			for i :=0 ; i<resultselection.Length() ; i++ {
				node := resultselection.Get(i)
				url = goquery.NewDocumentFromNode(node).Find("a").Text()
				if strings.ContainsAny(url,"jb39.com"){
					break
				}
			}

			//page.AddTargetRequest(band)
			fmt.Print("得到")
			fmt.Println(url)
			this.urls=append(this.urls,number+"#"+url)
		}
	}
	fmt.Println("分析结束")
}

func (this *GeturlAnalyer) AnalyzeFinsh() {
	fileName := "d:\\result.txt"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	for _,str := range this.urls {
		dstFile.WriteString(str + "\r\n")
	}

}