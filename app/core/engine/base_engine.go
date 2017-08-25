package engine

import (
	"github.com/beewit/found/app/common"
	"github.com/beewit/found/app/core/downloader"
	"github.com/beewit/found/app/core/pipeline"
	"github.com/beewit/found/app/core/processer"
	"github.com/beewit/found/app/core/scheduler"
	"github.com/beewit/found/app/plugin"
)

type BaseEngine interface {
	Start()
	SetStartUrl(string) *BaseEngine
	SetStartUrls([]string) *BaseEngine
	SetScheduler(scheduler.BaseScheduler) *BaseEngine
	SetDownloader(downloader.BaseDownloader) *BaseEngine
	SetProcesser(processer.BaseProcesser) *BaseEngine
	AddPipeline(pipeline.BasePipeline) *BaseEngine
	AddPlugin(plugin.BasePlugin) *BaseEngine
	SetConfig(*common.Config) *BaseEngine
}
