package downloader

import (
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
)

type Downloader interface {
	Download(request *crawler_request.CrawlerRequest)*page.Page
}