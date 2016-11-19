package main

import "github.com/YoungZhou93/startGolang/web/crawler/controller"

func main() {

	crawler := controller.NewDefaultController()
	crawler.SetUrl("http://www.baidu.com", "Get")
	crawler.Run()

}

