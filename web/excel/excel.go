package main

import (
	"github.com/tealeg/xlsx"
	"github.com/YoungZhou93/startGolang/web/crawler/controller"
	"strings"
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/page_analyzer"
	"strconv"
)

func main() {
	crawler := controller.NewDefaultController()
	var requests []*crawler_request.CrawlerRequest
	excelFileName := "d:\\数据.xlsx"
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i!=0 {
				cell := row.Cells[1]
				value, _ := cell.String()
				if strings.ContainsAny(value, "[") {
					str := strings.Split(value, "[")
					value = str[0]
				}
				request := crawler_request.NewCrawlerRequest("http://zhongyisousuo.com/s?q=" + value,"Get")
				request.SetLabel(strconv.Itoa(i))
				requests = append(requests, request)
			}
		}
	}

	crawler.SetRequests(requests)
	crawler.SetPageAnalyzer(page_analyzer.NewGeturlAnalyer())
	crawler.Run()
}
