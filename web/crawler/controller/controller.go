package controller

import (
	"github.com/YoungZhou93/startGolang/web/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/startGolang/web/crawler/downloader"
	"github.com/YoungZhou93/startGolang/web/crawler/scheduler"
	"github.com/YoungZhou93/startGolang/web/crawler/page_analyzer"
	"github.com/YoungZhou93/startGolang/web/crawler/outputer"
	"github.com/YoungZhou93/startGolang/web/crawler/common/thread_manage"
	"time"
	"sync"
	"fmt"
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
	threadTotal uint

	//程序状态
	status int


}

func NewDefaultController()*Controller{
	return NewController(1)
}

func NewController(threadTotal uint)*Controller{

	scheduler:=scheduler.NewListScheduler()
	downloader:=downloader.NewHttpDownloader()
	pageAnalyzer:=page_analyzer.NewDefaultAnalyer()
	outputer:=outputer.NewOutputerPrint()
	threadManager:=thread_manage.NewChanThreadManager(threadTotal)
	status:=1

	return &Controller{scheduler,downloader,pageAnalyzer,outputer,threadManager,threadTotal,status}

}


func (this *Controller) SetUrl(url string) {
	request := crawler_request.NewCrawlerRequest(url, "GET")
	this.SetRequest(request)
}

func (this *Controller) SetUrls(urls []string){
	for _,url:= range urls{
		request := crawler_request.NewCrawlerRequest(url, "GET")
		this.SetRequest(request)
	}
}

func (this *Controller) SetRequest(request *crawler_request.CrawlerRequest) {
	this.scheduler.PushBack(request)
}

func (this *Controller) SetRequests(requests []*crawler_request.CrawlerRequest) {
	for _,request:= range requests{
		this.scheduler.PushBack(request)
	}
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
	for {
		fmt.Println(this.status)
		for this.status<1{
			time.Sleep(500 * time.Millisecond)
		}


		//从队列中取出一个请求
		request:=this.scheduler.PopFront()

		//time.Sleep()

		if request==nil && this.threadManager.Left()==this.threadManager.Total(){
			this.pageAnalyzer.AnalyzeFinsh()
			break
		}else if request==nil{
			time.Sleep(500 * time.Millisecond)
			continue
		}
		//可用线程数减一
                this.threadManager.Get()

		go func(request *crawler_request.CrawlerRequest) {
			defer this.threadManager.Free()
			page:=this.downloader.Download(request)
			this.pageAnalyzer.Analyze(page)
			this.SetRequests(page.GetTargetRequests())
			this.outputer.Output(page)
		}(request)

	}
}

func(this *Controller)Stop(){
	this.status=0
}

func(this *Controller)Restart(){
	this.status=1
}