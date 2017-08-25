package scheduler

import (
	"github.com/beewit/found/app/common"
)

type BaseScheduler interface {
	Push(*common.Request)
	Poll() *common.Request
	Count() int
}
