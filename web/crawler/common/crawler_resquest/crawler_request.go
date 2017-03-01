package crawler_request

type CrawlerRequest struct {

	//请求的url
	Url string

	//请求方法(post/get)
	Method string

	//Post提交数据
	//PostData string

	//标签
	Label string

}

func NewCrawlerRequest(url string,method string)*CrawlerRequest{
	return &CrawlerRequest{url,method,""}
}

func (this *CrawlerRequest)GetUrl()string{
	return this.Url
}

func (this *CrawlerRequest)SetLabel(label string){
	 this.Label=label
}

func (this *CrawlerRequest)GetLabel()string{
	return this.Label
}