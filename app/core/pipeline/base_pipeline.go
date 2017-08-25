package pipeline

import (
	"github.com/beewit/found/app/common"
)

type BasePipeline interface {
	Pipe([]*common.Item, bool)
}
