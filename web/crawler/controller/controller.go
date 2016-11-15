package controller

import (
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/downloader"
	"github.com/YoungZhou93/startGolang/web/crawler/scheduler"
)

type Controller struct {

	scheduler scheduler.Scheduler
	downloader downloader.Downloader
}

func NewController()*Controller{

	scheduler:=scheduler.NewListScheduler()
	downloader:=downloader.NewHttpDownloader()

	return Controller{scheduler,downloader}

}



func (this *Controller) SetUrl(url string, method string) {
	request := crawler_request.NewCrawlerRequest(url, method)
	this.SetRequest(request)

}

func (this *Controller) SetRequest(request *crawler_request.CrawlerRequest) {
	this.scheduler.PushBack(request)
}

func(this*Controller)Run(){
	for{
		request:=this.scheduler.PopFront()
		go this.downloader.Download(request)
	}
}

func(this*Controller)Stop(){


}