package scheduler

import (
	"container/list"

	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
)

type ListScheduler struct {
	queue *list.List
}

func NewListScheduler() *ListScheduler {
	queue := list.New()
	return &ListScheduler{queue: queue}
}

func (this *ListScheduler) PushBack(request * crawler_request.CrawlerRequest) {
	this.queue.PushBack(request)

}

func (this *ListScheduler) PopFront() *crawler_request.CrawlerRequest {
	if this.queue.Len() <= 0 {
		return nil
	}
	request := this.queue.Front().Value.(crawler_request.CrawlerRequest)
	this.queue.Remove(this.queue.Front())
	return request
}

func (this *ListScheduler) Length() int {
	return this.queue.Len()
}
