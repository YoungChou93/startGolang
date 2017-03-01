package page_analyzer

import (
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
)

type PageAnalyzer interface {
	Analyze(page *page.Page)
	AnalyzeFinsh()
}

