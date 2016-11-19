package scheduler

import "github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"

type Scheduler interface {
	PushBack(request * crawler_request.CrawlerRequest)
	PopFront()(request * crawler_request.CrawlerRequest)
	Length()(int)
}
