package main

import (
	"github.com/beewit/beekit/utils"
	"github.com/beewit/found/app"
	"github.com/beewit/found/global"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
	"strings"
	"time"
	"github.com/beewit/beekit/utils/convert"
	"fmt"
)

func main() {
	var err error
	global.Driver = agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		"--start-maximized",
		"--disable-infobars",
		"--app=http://www.baidu.com",
		"--webkit-text-size-adjust"}))
	var startUrl = make([]string, 10)
	for i := 9; i >= 0; i-- {
		startUrl[i] = fmt.Sprintf("http://www.weixinqun.com/group?p=%d", i)
	}
	global.Driver.Start()

	err = app.Start(
		&app.Spider{
			Table:    "wechat_group",
			Name:     "weixinqun",
			Domain:   "http://www.weixinqun.com",
			StartUrl: startUrl,
			Sleep:    2,
		},
		func(spider *app.Spider) {
			thisUrl, _ := global.WechatGroupPage.URL()
			if thisUrl != "" {
				url := spider.TrimSpace(global.WechatGroupPage.URL())
				if thisUrl == "http://www.weixinqun.com/group" || strings.Contains(thisUrl, "http://www.weixinqun.com/group?p=") {
					//列表页
					html, err := global.WechatGroupPage.HTML()
					if err != nil {
						println(html)
					}
					eles, err := global.WechatGroupPage.Find("#tab_head").All("li").Elements()
					if err != nil {
						global.Log.Error(err.Error())
						return
					}
					if eles != nil {
						for i := range eles {
							e, _ := eles[i].GetElement(api.Selector{"css selector", ".goods_name a"})
							if e != nil {
								href, err := e.GetAttribute("href")
								if err != nil {
									global.Log.Error(err.Error())
									continue
								}
								if href != "" {
									flog := spider.AddQueue(href)
									global.Log.Info("添加链接：" + convert.ToString(flog))
								}
							}
						}
					}
				} else if strings.Contains(thisUrl, "http://www.weixinqun.com/group?id=") {
					//详情页面
					m := map[string]interface{}{}
					m["id"] = utils.ID()
					m["url"] = url
					m["name"] = spider.TrimSpace(global.WechatGroupPage.FindByClass("des_info_text").Text())
					m["intro"] = spider.TrimSpace(global.WechatGroupPage.FirstByClass("des_info_text2").Text())
					m["type"] = spider.GetSelector(global.WechatGroupPage.Find(".other-info"), "li").GetFirstFind("a").GetFirstText()
					reTime := spider.GetSelector(global.WechatGroupPage.Find(".other-info"), "li").Eq(1).GetText()
					if reTime != "" {
						reTime = strings.Replace(reTime, "时间：", "", -1)
						reTime = strings.TrimSpace(reTime)
					}
					m["re_time"] = reTime
					m["area"] = spider.GetSelector(global.WechatGroupPage.Find(".other-info"), "li").Eq(2).GetFirstFind("a").GetText()
					m["code"] = spider.GetSelector(global.WechatGroupPage.Find(".other-info"), "li").Eq(3).GetFirstFind("a").GetText()

					m["owner_wechat"] = spider.GetSelectorAll(global.WechatGroupPage, ".des_info_text2").Eq(1).GetText()
					hot := 0
					hotStr := spider.TrimSpace(global.WechatGroupPage.FirstByClass("Pink").Text())
					if utils.IsValidNumber(hotStr) {
						hot = convert.MustInt(hotStr)
					}
					m["hot"] = hot
					m["ct_time"] = utils.CurrentTime()
					m["ut_time"] = m["ct_time"]
					m["domain"] = spider.Domain
					if convert.ToString(m["name"]) != "" {
						spider.AddDataQueue(m)
						spider.AddDoneQueue(convert.ToString(m["url"]))
					} else {
						spider.AddFailedQueue(convert.ToString(m["url"]))
					}
				}
			}
		},
		func(spider *app.Spider) {
			for {
				global.Log.Info("【%s】数据保存程序启动", spider.Name)

				spider.SaveData()

				time.Sleep(time.Second * 10)
			}
		})
	if err != nil {
		global.Log.Error(err.Error())
	}

}
