package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/beewit/found/app/common"
	"github.com/beewit/found/app/core/engine"
	"github.com/beewit/found/app/core/extractor"
	"github.com/beewit/found/app/core/pipeline"
)

type MyProcesser struct{}

func NewMyProcesser() *MyProcesser {
	return &MyProcesser{}
}

func (this *MyProcesser) Process(resp *common.Response, y *common.Yield) {
	items := extractor.NewExtractor().
		SetScopeRule(`(?s)<dt class="basicInfo-item name">.*?</dd>`).
		SetRules(map[string]string{
			"key":   `(?s)name">(.*?)</dt>`,
			"value": `(?s)value">(.*?)</dd>`,
		}).
		SetTrimFunc(extractor.TrimHtmlTags).
		Extract(resp)

	for _, item := range items {
		y.AddItem(item)
	}
}

func getUrlsFromFile(fileName string) []string {
	var urls = []string{}
	file, _ := os.Open(fileName)
	r := bufio.NewReader(file)
	for i := 0; i < 10; i++ {
		line, _ := r.ReadString('\n')
		urls = append(urls, strings.TrimSpace(line))
	}
	return urls
}

func main() {
	engine.NewEngine("crawl_baidubaike_with_extractor").
		AddPipeline(pipeline.NewConsolePipeline()).
		SetProcesser(NewMyProcesser()).
		SetStartUrls(getUrlsFromFile("test.url")).
		Start()
}
