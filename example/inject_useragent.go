package main

import (
	"regexp"

	"github.com/beewit/found/app/common"
	"github.com/beewit/found/app/core/engine"
	"github.com/beewit/found/app/core/pipeline"
	"github.com/beewit/found/app/plugin"
)

type MyProcesser struct{}

func NewMyProcesser() *MyProcesser {
	return &MyProcesser{}
}

func (this *MyProcesser) Process(resp *common.Response, y *common.Yield) {
	m := regexp.MustCompile(`(?s)<div id="ua_string">.*?</span>(.*?)</div>`).FindAllStringSubmatch(resp.Body, -1)
	for _, v := range m {
		item := common.NewItem()
		item.Set("user-agent", v[1])
		y.AddItem(item)
	}
}

func main() {
	engine.NewEngine("inject_useragent").
		SetStartUrl("http://my-user-agent.com/").
		SetProcesser(NewMyProcesser()).
		AddPlugin(plugin.NewUserAgentPlugin()).
		AddPipeline(pipeline.NewConsolePipeline()).
		SetConfig(common.NewConfig().SetHeaders(map[string]string{"User-Agent": "Dalvik/2.1.0 (Linux; U; Android 7.0; Mi-4c MIUI/7.3.16)"})).
		Start()
}
