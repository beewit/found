package main

import (
	"github.com/beewit/found/app/core/engine"
	"github.com/beewit/found/app/core/pipeline"
)

func main() {
	url := "http://m.baidu.com"
	engine.NewEngine("crawl_baidu_and_print_it").AddPipeline(pipeline.NewConsolePipeline()).SetStartUrl(url).Start()
}
