package crawler_request

type CrawlerRequest struct {

	//请求的url
	Url string

	//请求方法(post/get)
	Method string

	//Post提交数据
	//PostData string

}

func NewCrawlerRequest(url string,method string)*CrawlerRequest{
	return &CrawlerRequest{url,method}
}

func (this *CrawlerRequest)GetUrl()string{
	return this.Url
}