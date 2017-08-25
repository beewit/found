package main

import (
	"github.com/beewit/found/app/core/engine"
)

func main() {
	engine.NewQuickEngine("crawl_baidubaike_with_quickengine.json").Start()
}
