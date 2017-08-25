package processer

import (
	"github.com/beewit/found/app/common"
)

type BaseProcesser interface {
	Process(*common.Response, *common.Yield)
}
