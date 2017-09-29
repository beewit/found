package main

import (
	"encoding/json"
	"github.com/beewit/found/global"
	"github.com/beewit/beekit/redis"
	"time"
	"github.com/sclevine/agouti"
	"net/http"
)

const (
	SINA_SPIDER        = "SINA_SPIDER"
	SINA_SPIDER_DONE   = "SINA_SPIDER_DONE"
	SINA_SPIDER_FAILED = "SINA_SPIDER_FAILED"
	SINA_SPIDER_DATA   = "SINA_SPIDER_DATA"
)

func addSinaQueue(value string) {
	if !checkSinaDoneQueue(value) && !checkSinaFailedQueue(value) && !checkSinaQueue(value) {
		global.RD.SetSETString(SINA_SPIDER, value)
	}
}
func getSinaQueue() string {
	v, _ := global.RD.GetSETRandStringRm(SINA_SPIDER)
	return v
}
func checkSinaQueue(value string) bool {
	x, _ := global.RD.CheckSETString(SINA_SPIDER, value)
	return x > 0
}

func addSinaDoneQueue(value string) {
	global.RD.SetSETString(SINA_SPIDER_DONE, value)
}
func getSinaDoneQueue(value string) string {
	v, _ := global.RD.GetSETRandStringRm(SINA_SPIDER_DONE)
	return v
}
func checkSinaDoneQueue(value string) bool {
	x, _ := global.RD.CheckSETString(SINA_SPIDER_DONE, value)
	return x > 0
}

func addSinaFailedQueue(value string) {
	global.RD.SetSETString(SINA_SPIDER_FAILED, value)
}
func getSinaFailedQueue(value string) string {
	v, _ := global.RD.GetSETRandStringRm(SINA_SPIDER_FAILED)
	return v
}
func checkSinaFailedQueue(value string) bool {
	x, _ := global.RD.CheckSETString(SINA_SPIDER_FAILED, value)
	return x > 0
}

func addSinaDataQueue(value map[string]interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		global.Log.Info(err.Error())
		return
	}
	x, err := global.RD.SetSETString(SINA_SPIDER_DATA, string(b))
	if err != nil {
		global.Log.Info(err.Error())
	}
	global.Log.Info("添加数据到Redis结果%v", x)
}

func getSinaDataQueue() (string, error) {
	return global.RD.GetSETRandStringRm(SINA_SPIDER_DATA)
}

func getSinaData() map[string]interface{} {
	d, err := getSinaDataQueue()
	if err != nil {
		global.Log.Warning("无数据", err.Error())
		return nil
	}
	if d != "" {
		m := map[string]interface{}{}
		err = json.Unmarshal([]byte(d), &m)
		if err == nil {
			return m
		}
	}
	return nil
}

func main() {
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		"--start-maximized",
		"--disable-infobars",
		"--app=https://weibo.com/",
		"--webkit-text-size-adjust"}))
	driver.Start()
	var err error
	global.Page.Page, err = driver.NewPage()
	if err != nil {
		global.Log.Info(err.Error())
	} else {
		flog, ce := SetSinaCookieLogin("sinaWeibo" + global.SinaWeiboUName)
		if ce != nil {
			global.Log.Info(ce.Error())
			return
		}
		if flog {
			//global.Page.Navigate("https://weibo.com")
			//time.Sleep(time.Second * 3)
			name, ei := global.Page.Find(".nameBox").Text()
			if ei != nil {
				global.Log.Info(ei.Error())
			}
			if name != "" {
				global.Log.Info("头像", name)
			} else {
				flog = false
			}
		}
		if !flog {
			//global.Page.Navigate("http://weibo.com/login.php")
			//time.Sleep(time.Second * 3)
			global.Page.FindByID("loginname").Fill(global.SinaWeiboUName)
			time.Sleep(time.Second * 1)
			global.Page.Find("input[type=\"password\"]").Fill(global.SinaWeiboPwd)
			time.Sleep(time.Second * 2)
			global.Page.FindByXPath("//*[@id=\"pl_login_form\"]/div/div[3]/div[6]/a/span").Click()
			time.Sleep(time.Second * 2)
			c, err5 := global.Page.GetCookies()
			if err5 != nil {
				global.Log.Info("登陆失败", err5.Error())
				return
			}
			cookieJson, _ := json.Marshal(c)
			//global.Log.Info("cookie", string(cookieJson[:]))
			redis.Cache.SetString("sinaWeibo"+global.SinaWeiboUName, string(cookieJson[:]))

		}

	}
}

func SetSinaCookieLogin(key string) (bool, error) {
	rc := redis.Cache
	cookieRd, _ := rc.GetString(key)
	if cookieRd == "" {
		return false, nil
	}
	var cks = []*http.Cookie{}
	err := json.Unmarshal([]byte(cookieRd), &cks)
	if err != nil {
		return false, err
	}
	for i := range cks {
		cc := cks[i]
		global.Page.SetCookie(cc)
	}
	return true, nil
}
