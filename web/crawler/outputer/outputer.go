package outputer

import (
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
)

type Outputer interface {

	Output(page *page.Page)
}
