package outputer

import (
	"fmt"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
)

type OutputerPrint struct {
}

func NewOutputerPrint() *OutputerPrint {
	return &OutputerPrint{}
}

func (this *OutputerPrint) Output(page *page.Page) {
	fmt.Println("开始输出")
	fmt.Print("Url:")
	fmt.Println(page.GetRequest().GetUrl())
	fmt.Print("Header:")
	fmt.Println(page.GetHeader())
}
