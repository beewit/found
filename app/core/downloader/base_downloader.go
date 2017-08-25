package downloader

import (
	"github.com/beewit/found/app/common"
)

type BaseDownloader interface {
	Download(*common.Request, *common.Config) (*common.Response, error)
}
