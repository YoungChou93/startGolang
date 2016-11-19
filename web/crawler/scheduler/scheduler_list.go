package scheduler

import (
	"container/list"
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"sync"
)

type ListScheduler struct {
	queue *list.List
	locker *sync.Mutex
}

func NewListScheduler() *ListScheduler {
	queue := list.New()
	locker := new(sync.Mutex)
	return &ListScheduler{queue: queue,locker:locker}
}

func (this *ListScheduler) PushBack(request * crawler_request.CrawlerRequest) {
	this.locker.Lock()
	this.queue.PushBack(request)
	this.locker.Unlock()

}

func (this *ListScheduler) PopFront() *crawler_request.CrawlerRequest {
	this.locker.Lock()
	if this.queue.Len() <= 0 {
		this.locker.Unlock()
		return nil
	}
	request := this.queue.Front().Value.(*crawler_request.CrawlerRequest)
	this.queue.Remove(this.queue.Front())
	this.locker.Unlock()
	return request
}

func (this *ListScheduler) Length() int {
	this.locker.Lock()
	length:=this.queue.Len()
	this.locker.Unlock()
	return length
}
