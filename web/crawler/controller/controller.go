package controller

import (
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/downloader"
	"github.com/YoungZhou93/startGolang/web/crawler/scheduler"
	"github.com/YoungZhou93/startGolang/web/crawler/page_analyzer"
	"github.com/YoungZhou93/startGolang/web/crawler/outputer"
	"github.com/YoungZhou93/startGolang/web/crawler/common/thread_manage"
)

type Controller struct {

	//请求队列组件
	scheduler scheduler.Scheduler
	//下载组件
	downloader downloader.Downloader
	//页面分析组件
	pageAnalyzer page_analyzer.PageAnalyzer
	//输出组件
	outputer outputer.Outputer

	//线程管理组件
	threadManager thread_manage.ThreadManager
	threadnum uint
}

func NewDefaultController()*Controller{
	return NewController(1)
}

func NewController(threadnum uint)*Controller{
	scheduler:=scheduler.NewListScheduler()
	downloader:=downloader.NewHttpDownloader()
	pageAnalyzer:=page_analyzer.NewDefaultAnalyer()
	outputer:=outputer.NewOutputerPrint()
	threadManager:=thread_manage.NewChanThreadManager(threadnum)

	return &Controller{scheduler,downloader,pageAnalyzer,outputer,threadManager,threadnum}

}


func (this *Controller) SetUrl(url string, method string) {
	request := crawler_request.NewCrawlerRequest(url, method)
	this.SetRequest(request)
}

func (this *Controller) SetRequest(request *crawler_request.CrawlerRequest) {
	this.scheduler.PushBack(request)
}

func (this *Controller)SetThreadManager(threadManager thread_manage.ThreadManager)*Controller{
	this.threadManager=threadManager
	return this
}

func (this *Controller)SetScheduler(scheduler scheduler.Scheduler)*Controller{
	this.scheduler=scheduler
	return this
}

func (this *Controller)SetDownloader(downloader downloader.Downloader)*Controller{
	this.downloader=downloader
	return this
}

func (this *Controller)SetPageAnalyzer( pageAnalyzer page_analyzer.PageAnalyzer)*Controller{
	this.pageAnalyzer=pageAnalyzer
	return this
}

func (this *Controller)SetOutputer(outputer outputer.Outputer)*Controller{
	this.outputer=outputer
	return this
}

func(this *Controller)Run(){
	for this.scheduler.Length()>0{
		request:=this.scheduler.PopFront()
		page:=this.downloader.Download(request)
		this.outputer.Output(page)
	}
}

func(this *Controller)Stop(){


}