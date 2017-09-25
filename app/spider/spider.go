package spider

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beewit/beekit/utils"
	"github.com/beewit/found/global"
	"github.com/sclevine/agouti"
)

var (
	urls []string
)

type rule struct {
	Name           string            `json:"name"`           //名称
	Mode           string            `json:"mode"`           //模式，Chrome/HttpClient	默认：HttpClient
	StartPage      []string          `json:"startPage"`      //开始页
	ConnTimeout    int               `json:"connTimeout"`    //连接状态超时 WSARecv tcp: i/o timeout 默认：3s
	TryTimes       int               `json:"tryTimes"`       //尝试下载的最大次数						默认：1 次
	RetryPause     int               `json:"retryPause"`     //下载失败后，下次尝试下载的等待时间
	Reloadable     bool              `json:"reloadable"`     //是否允许重复该链接下载
	DelayTime      byte              `json:"delayTime"`      //两次采集间隔时间
	HelperURL      string            `json:"helperUrl"`      //分页地址
	DetailURL      string            `json:"detailUrl"`      //详情地址
	Fields         map[string]string `json:"fields"`         //要保存的字段列表
	Redirect       bool              `json:"redirect"`       //是否允许重定向
	RemovalReg     string            `json:"removalReg"`     //队列校验的重复标记 		 	   默认：无
	RemovalRules   string            `json:"removalRules"`   //重复标记检测规则 M:月，d：天，h：小时，m：分钟，s：秒 1M 1d 1h 1m 1s	 默认：永久唯一校验
	Thread         int               `json:"thread"`         //线程数量                      默认：1
	Header         []string          `json:"header"`         //HTTP Head 			    	 默认：无
	Cookies        []string          `json:"cookies"`        //HTTP Cookies 		    	 默认：无
	QueueURLRules  map[string]string `json:"queueURLRules"`  //进入任务队列的URL规则 regex | equal | start | end | contains 默认：regex
	StorageService string            `json:"storageService"` //存储服务，默认
	//Priority      int               `json:"priority"`      //指定调度优先级，默认为0（最小优先级为0）  ** 提到爬虫属性中便于排序
}

type Spider struct {
	rule       rule //具体任务
	page       *agouti.Page
	chanQueue  chan string
	chanDone   chan string
	chanFailed chan string
}

func NewSpider(task map[string]interface{}, page *agouti.Page) *Spider {

	return &Spider{
		page: page,
		//待抓
		chanQueue: make(chan string, 300),
		//已抓
		chanDone: make(chan string, 300),
		//错误
		chanFailed: make(chan string, 300),
	}
}

func (s *Spider) Run() {

}

func (s *Spider) Fetch(url string) {
	defer s.Destroy()
	s.page.Navigate(url)
	body, err := s.page.HTML()
	if err != nil {
		global.Log.Error(err.Error())
	}
	s.chanQueue <- body
}

func (s *Spider) Extractor() {
	for {
		//todo 以下是解析逻辑
		html := <-s.chanQueue
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			if utils.IsUrl(href) {
				title := strings.TrimSpace(s.Text())
				if len(title) > 3 {
					log.Printf("%s - %s\n", title, href)
				}
			}
		})
	}
}

func (s *Spider) AddUrlQueue(chanUrl chan string) {

}

func (s *Spider) PostSpiderTaskResult(chanMaps chan map[string]string) {

}

func (s *Spider) SaveQueueToFile() {
	//把
}

func (s *Spider) Destroy() {
	close(s.chanQueue)
	close(s.chanDone)
	close(s.chanFailed)
}
