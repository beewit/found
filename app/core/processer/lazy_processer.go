package processer

import (
	"github.com/beewit/found/app/common"
)

type LazyProcesser struct{}

func NewLazyProcesser() *LazyProcesser {
	return &LazyProcesser{}
}

func (this *LazyProcesser) Process(resp *common.Response, y *common.Yield) {
	y.AddItem(func() *common.Item {
		item := common.NewItem()
		item.Set("html", resp.Body)
		return item
	}())
}
