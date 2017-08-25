package main

import (
	"github.com/beewit/found/app/core/engine"
)

func main() {
	engine.
		NewQuickEngine("spider.json").
		GetEngine().
		//AddPlugin(plugin.NewProxyPlugin()).
		Start()
}
