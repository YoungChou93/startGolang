package main

import (
	"bufio"
	"fmt"
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/controller"
	"github.com/YoungZhou93/startGolang/web/crawler/page_analyzer"
	"io"
	"os"
	"strings"
	"strconv"
)

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func main() {
	crawler := controller.NewDefaultController()
	var requests []*crawler_request.CrawlerRequest
	fileName := "d:\\result.txt"
	f, _ := os.Open(fileName)
	buf := bufio.NewReader(f)
	i := 0
	for {

		line, _ := buf.ReadString('\r')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		strs := strings.Split(line, "#")

		request := crawler_request.NewCrawlerRequest(strs[1], "Get")
		request.SetLabel(strs[0])
		number,_:=strconv.Atoi(strs[0])
		if number<500 &&  number>= 400 {
			requests = append(requests, request)
		}

		if strings.Contains(line, "2273") {
			break
		}

		i++
	}

	crawler.SetRequests(requests)
	crawler.SetPageAnalyzer(page_analyzer.NewSaveAnalyer())
	crawler.Run()
}
