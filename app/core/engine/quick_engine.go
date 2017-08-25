package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/beewit/found/app/common"
	"github.com/beewit/found/app/core/extractor"
	"github.com/beewit/found/app/core/pipeline"
)

type QuickEngine struct {
	quickEngineConfig *QuickEngineConfig
	file              *os.File
	logFile           *os.File
	resetOutput       bool
}

func NewQuickEngine(quickEngineConfigPath string) *QuickEngine {
	qe := &QuickEngine{
		file:        nil,
		resetOutput: false,
	}
	qe.quickEngineConfig = NewQuickEngineConfig(quickEngineConfigPath)
	return qe
}

func (this *QuickEngine) GetQuickEngineConfig() *QuickEngineConfig {
	return this.quickEngineConfig
}

func (this *QuickEngine) GetEngine() *Engine {
	e := NewEngine(this.quickEngineConfig.TaskName).
		SetProcesser(NewQuickEngineProcesser(this.quickEngineConfig)).
		SetStartUrls(this.quickEngineConfig.StartUrls).
		SetConfig(this.quickEngineConfig.ToCommonConfig())

	if this.quickEngineConfig.LogFile != "" {
		var err error
		this.logFile, err = os.OpenFile(this.quickEngineConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(this.logFile)
		if err != nil {
			log.Fatal("cannot create log file: %s\n", this.quickEngineConfig.LogFile)
		}
	}

	if this.file != nil {
		return e.AddPipeline(pipeline.NewFilePipeline(this.file))
	} else if this.quickEngineConfig.OutputFile != "" {
		this.file, _ = os.Create(this.quickEngineConfig.OutputFile)
		return e.AddPipeline(pipeline.NewFilePipeline(this.file))
	}
	return e.AddPipeline(pipeline.NewConsolePipeline())

}

func (this *QuickEngine) SetOutputFile(file *os.File) *QuickEngine {
	this.file = file
	this.resetOutput = true
	return this
}
func (this *QuickEngine) Start() {
	if this.file != nil && !this.resetOutput {
		defer this.file.Close()
	}
	if this.logFile != nil {
		defer this.logFile.Close()
	}
	this.GetEngine().Start()
}

type QuickEngineConfig struct {
	TaskName   string   `json:"task_name"`
	StartUrls  []string `json:"start_urls"`
	Rules      []_Rule  `json:"rules"`
	OutputFile string   `json:"output_file"`
	Config     _Config  `json:"config"`
	LogFile    string   `json:"log_file"`
}

type _Rule struct {
	UrlMatch    string       `json:"url_match"`
	Succ        string       `json:"succ"`
	BaseUrl     string       `json:"base_url"`
	ItemRule    _ItemRule    `json:"item_rule"`
	RequestRule _RequestRule `json:"request_rule"`
	Merge       bool         `json:"merge"`
}

type _ItemRule struct {
	ScopeRule string            `json:"scope_rule"`
	KVRule    map[string]string `json:"kv_rule"`
	TrimFunc  string            `json:"trim_func"`
}

type _RequestRule _ItemRule

type _Config struct {
	Concurrency         int               `json:"concurrency"`
	PollingTime         string            `json:"polling_time"`
	WaitTime            string            `json:"wait_time"`
	DownloadTimeout     string            `json:"download_timeout"`
	ConnectionTimeout   string            `json:"connection_timeout"`
	MaxIdleConnsPerHost int               `json:"max_idle_conns_per_host"`
	MaxRetryTimes       int               `json:"max_retry_times"`
	MaxDepth            int               `json:"max_depth"`
	Logging             bool              `json:"logging"`
	Headers             map[string]string `json:"headers"`
	Succ                string            `json:"succ"`
}

func NewQuickEngineConfig(fileName string) *QuickEngineConfig {
	c := &QuickEngineConfig{}
	t := reflect.TypeOf(&c.Config)
	v := reflect.ValueOf(&c.Config)
	config := common.NewConfig()

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		value := v.Elem().FieldByName(field.Name)
		funcName := "Get" + field.Name
		switch value.Type().Kind() {
		case reflect.Int:
			defaultValue := reflect.ValueOf(config).MethodByName(funcName).Call([]reflect.Value{})
			value.SetInt(defaultValue[0].Int())
		case reflect.String:
			defaultValue := reflect.ValueOf(config).MethodByName(funcName).Call([]reflect.Value{})
			value.SetString(fmt.Sprintf("%s", defaultValue[0].Interface()))
		case reflect.Bool:
			defaultValue := reflect.ValueOf(config).MethodByName(funcName).Call([]reflect.Value{})
			value.SetBool(defaultValue[0].Bool())
		}
	}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, c)
	return c
}

func (this *QuickEngineConfig) stringToDuration(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

func (this *QuickEngineConfig) ToCommonConfig() *common.Config {
	e := common.NewConfig().SetConcurrency(this.Config.Concurrency).
		SetPollingTime(this.stringToDuration(this.Config.PollingTime)).
		SetWaitTime(this.stringToDuration(this.Config.WaitTime)).
		SetDownloadTimeout(this.stringToDuration(this.Config.DownloadTimeout)).
		SetConnectionTimeout(this.stringToDuration(this.Config.ConnectionTimeout)).
		SetMaxIdleConnsPerHost(this.Config.MaxIdleConnsPerHost).
		SetMaxRetryTimes(this.Config.MaxRetryTimes).
		SetMaxDepth(this.Config.MaxDepth).
		SetLogging(this.Config.Logging).
		SetHeaders(this.Config.Headers).
		SetSucc(this.Config.Succ)
	return e
}

type QuickEngineProcesser struct {
	config *QuickEngineConfig
	depth  int
}

func NewQuickEngineProcesser(config *QuickEngineConfig) *QuickEngineProcesser {
	return &QuickEngineProcesser{config: config}
}

func (this *QuickEngineProcesser) processItems(resp *common.Response, y *common.Yield, rule _Rule) {
	var TrimFunc extractor.TrimFunc
	switch rule.ItemRule.TrimFunc {
	case "trim_html_tags":
		TrimFunc = extractor.TrimHtmlTags
	case "trim_blank":
		TrimFunc = extractor.TrimBlank
	}

	items := extractor.NewExtractor().
		SetScopeRule(rule.ItemRule.ScopeRule).
		SetRules(rule.ItemRule.KVRule).
		SetTrimFunc(TrimFunc).
		Extract(resp)
	for _, item := range items {
		y.AddItem(item)
	}
}

func (this *QuickEngineProcesser) processRequests(resp *common.Response, y *common.Yield, rule _Rule) {
	var TrimFunc extractor.TrimFunc
	switch rule.RequestRule.TrimFunc {
	case "trim_html_tags":
		TrimFunc = extractor.TrimHtmlTags
	case "trim_blank":
		TrimFunc = extractor.TrimBlank
	}

	items := extractor.NewExtractor().
		SetScopeRule(rule.RequestRule.ScopeRule).
		SetRules(rule.RequestRule.KVRule).
		SetTrimFunc(TrimFunc).
		Extract(resp)
	for _, item := range items {
		for _, url := range item.GetAll() {
			if strings.HasPrefix(url, "http://") {
				y.AddRequest(common.NewRequest(url))
			} else {
				y.AddRequest(common.NewRequest(rule.BaseUrl + url))
			}
		}
	}
}

func (this *QuickEngineProcesser) Process(resp *common.Response, y *common.Yield) {
	common.Try(func() {
		for _, rule := range this.config.Rules {
			if regexp.MustCompile(rule.UrlMatch).MatchString(resp.Url) {
				if rule.Succ != "" && !strings.Contains(resp.Body, rule.Succ) {
					log.Printf("cannot find succ string:%s", rule.Succ)
					break
				}
				if rule.ItemRule.ScopeRule != "" {
					this.processItems(resp, y, rule)
				}
				if rule.RequestRule.ScopeRule != "" {
					this.processRequests(resp, y, rule)
				}
				y.SetMerge(rule.Merge)
				break //Only use the first match
			}
		}
	}, func(e interface{}) {
		log.Printf("pannic %s\n", e)
	})
}
