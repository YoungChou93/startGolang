package downloader

import "github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"

type Downloader interface {
	Download(request *crawler_request.CrawlerRequest)
}