package main

import (
	"encoding/json"
	"github.com/beewit/found/global"
	"github.com/beewit/beekit/redis"
	"time"
	"github.com/sclevine/agouti"
	"net/http"
	"strings"
	"github.com/sclevine/agouti/api"
	"github.com/beewit/beekit/utils"
	"github.com/beewit/beekit/utils/convert"
)

const (
	SINA_SPIDER        = "SINA_SPIDER"
	SINA_SPIDER_DONE   = "SINA_SPIDER_DONE"
	SINA_SPIDER_FAILED = "SINA_SPIDER_FAILED"
	SINA_SPIDER_DATA   = "SINA_SPIDER_DATA"
)

func getSinaQueueCount() int64 {
	c, err := global.RD.GetSETCount(SINA_SPIDER)
	if err != nil {
		println("getSinaQueueCount，Error:", err.Error())
		return 0
	}
	return c
}
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
	go saveSinaData()
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
			time.Sleep(time.Second * 3)
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

			if getSinaQueueCount() <= 0 {
				/*获取粉丝和关注的用户href*/
				ClikFans()
				/*进入用户微博获取信息*/
				newSina := getSinaQueue()
				if newSina != "" {
					getUserInfo(newSina)
				}
			} else {
				newSina := getSinaQueue()
				if newSina != "" {
					getUserInfo(newSina)
				}
			}
		}

	}
}

func ClikFans() {
	eles, err := global.Page.Page.Find(".user_atten").All("li").Elements()
	if err != nil {
		return
	}
	if eles != nil && len(eles) > 1 {
		a, err := eles[0].GetElement(api.Selector{"css selector", "a"})
		if err != nil {
			global.Log.Info("获取关注节点失败", err.Error())
			return
		}
		href, _ := a.GetAttribute("href")

		a2, err := eles[1].GetElement(api.Selector{"css selector", "a"})
		if err != nil {
			global.Log.Info("获取关注节点失败", err.Error())
			return
		}
		href2, _ := a2.GetAttribute("href")
		if href != "" {
			getUserHref(href)
		}
		if href2 != "" {
			getUserHref(href2)
		}
	}
}

func ClikFrendsFans() {
	eles, err := global.Page.Page.Find(".PCD_counter").All("td.S_line1").Elements()
	if err != nil {
		println("获取微博关注/粉丝节点失败，", err.Error())
	} else {
		if eles != nil && len(eles) > 1 {
			a, err := eles[0].GetElement(api.Selector{"css selector", "a"})
			if err != nil {
				global.Log.Info("获取关注节点失败", err.Error())
				return
			}
			href, _ := a.GetAttribute("href")

			a2, err := eles[1].GetElement(api.Selector{"css selector", "a"})
			if err != nil {
				global.Log.Info("获取关注节点失败", err.Error())
				return
			}
			href2, _ := a2.GetAttribute("href")
			if href != "" {
				getUserHref(href)
			}
			if href2 != "" {
				getUserHref(href2)
			}
		}
	}
}

func getUserHref(href string) {
	if href != "" {
		println("待抓url：", href)
		global.Page.Navigate(href)
	} else {
		println("点击操作项")
	}
	eles, err := global.Page.Find(".follow_list").All("li.follow_item").Elements()
	if err != nil {
		eles, err = global.Page.Find(".member_ul").All("li.member_li").Elements()
		if err != nil {
			println("未找到微博用户列表标签")
		}
	}
	if err != nil {
		global.Log.Info("获取列表href错误", err.Error())
		return
	}
	if eles != nil {
		for i := range eles {
			a, err := eles[i].GetElement(api.Selector{"css selector", " .mod_pic a"})
			if err != nil {
				println("未找到列表中的微博用户节点标签")
				continue
			}
			//存入待抓队列
			href, err := a.GetAttribute("href")
			if err != nil {
				global.Log.Info("获取粉丝用户详情href失败", err.Error())
			} else {
				newHref := strings.Split(href, "?")
				addSinaQueue(newHref[0])
			}
		}
	}
	//下一页
	next := global.Page.Page.Find(".W_pages .next")
	cls, err := next.Attribute("class")
	if err != nil {
		global.Log.Info("获取下一页按钮样式失败", err.Error())
		return
	}
	if !strings.Contains(cls, "page_dis") {
		global.Page.RunScript("document.documentElement.scrollTop=document.body.clientHeight;", nil, nil)
		time.Sleep(time.Second * 2)
		next.Click()
		time.Sleep(time.Second * 5)
		html, _ := global.Page.HTML()
		if !strings.Contains(html, "由于系统限制，你无法查看") {
			getUserHref("")
		}
	}
}

func getUserInfo(href string) {
	global.Page.Page.Navigate(href)
	addSinaDoneQueue(href)

	m := map[string]interface{}{}
	m["url"] = href
	m["nickname"], _ = global.Page.Page.FindByClass("username").Text()
	m["head_img"], _ = global.Page.Page.FindByClass("photo").Attribute("src")
	m["desc"], _ = global.Page.Page.FindByClass("pf_intro").Text()
	eles, err := global.Page.Page.Find(".PCD_counter").All("strong").Elements()
	if err != nil {
		println("获取微博关注/粉丝/微博失败，", err.Error())
	} else {
		m["tweets"], _ = eles[0].GetText()
		m["follows"], _ = eles[1].GetText()
		m["fans"], _ = eles[2].GetText()
	}
	//获取微博说说
	m["weibos"] = getUserWeibo()
	//获取用户详细信息
	m = getUserInfoDetail(m)
	addSinaDataQueue(m)

	/*获取粉丝和关注的用户href*/
	ClikFrendsFans()

	newSina := getSinaQueue()
	if newSina != "" {
		getUserInfo(newSina)
	}
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func saveSinaData() {
	for {
		Try(func() {
			global.Log.Info("执行Sina数据保存服务")
			m := getSinaData()
			if m != nil {

				weibos := m["weibos"]
				delete(m, "weibos")

				sinaBaseId :=utils.ID()
				m["id"] = sinaBaseId
				m["ct_time"] = utils.CurrentTime()
				m["ut_time"] = utils.CurrentTime()
				if convert.ToString(m["desc"]) != "" {
					m["mydesc"] = m["desc"]
				}
				delete(m, "desc")
				_, err := global.DB.InsertMap("sina_user_base_info", m)
				if err != nil {
					global.Log.Error("sina_user_base_info 数据保存失败:", err.Error())
				} else {
					if weibos != nil {
						wb, err := convert.Obj2ListMap(weibos)
						if err != nil {
							global.Log.Error("weibo 数据转换[]map失败:", err.Error())
						} else {
							for i := 0; i < len(wb); i++ {
								wb[i]["id"] =utils.ID()
								wb[i]["sina_id"] = sinaBaseId
								if convert.ToString(wb[i]["like"]) != "" {
									wb[i]["likes"] = wb[i]["like"]
								}
								delete(wb[i], "like")
								if convert.ToString(wb[i]["like"]) != "" {
									wb[i]["source"] = wb[i]["from"]
								}
								delete(wb[i], "from")
							}
							_, err = global.DB.InsertMapList("weibo", wb)
							if err != nil {
								global.Log.Error("weibo 数据保存失败:", err.Error())
							}
						}
					}
				}
			}
			time.Sleep(time.Second * 10)
		}, func(e interface{}) {
			print(e)
		})
	}
}

func getUserWeibo() []map[string]interface{} {
	eles, err := global.Page.Page.Find(".WB_feed").All(".WB_feed_like").Elements()
	if err != nil {
		println("获取微博说说失败，", err.Error())
		return nil
	}
	if eles != nil {
		weiboMap := []map[string]interface{}{}
		for i := range eles {
			m := map[string]interface{}{}
			ele, _ := eles[i].GetElement(api.Selector{"css selector", ".WB_text"})
			m["content"] = nil
			if ele != nil {
				m["content"], _ = ele.GetText()
			}
			ele, _ = eles[i].GetElement(api.Selector{"css selector", ".WB_feed_expand"})
			m["transfer_content"] = nil
			if ele != nil {
				m["transfer_content"], _ = ele.GetText()
			}

			eles2, _ := eles[i].GetElements(api.Selector{"css selector", ".WB_from a"})
			m["say_time"] = nil
			m["source"] = nil
			if eles2 != nil && len(eles2) > 0 {
				m["say_time"], _ = eles2[0].GetText()
				if len(eles2) > 1 {
					m["source"], _ = eles2[1].GetText()
				}
			}
			//点赞
			eles2, _ = eles[i].GetElements(api.Selector{"css selector", "span[node-type=like_status] em"})
			m["likes"] = nil
			if ele != nil && len(eles2) > 1 {
				like, _ := eles2[1].GetText()
				if utils.IsValidNumber(like) {
					m["likes"] = like
				}
			}
			//评论
			eles2, _ = eles[i].GetElements(api.Selector{"css selector", "span[node-type=comment_btn_text] em"})
			m["comment"] = nil
			if ele != nil && len(eles2) > 1 {
				comment, _ := eles2[1].GetText()
				if utils.IsValidNumber(comment) {
					m["comment"] = comment
				}
			}
			//转载
			eles2, _ = eles[i].GetElements(api.Selector{"css selector", "span[node-type=forward_btn_text] em"})
			m["transfer"] = nil
			if ele != nil && len(eles2) > 1 {
				transfer, _ := eles2[1].GetText()
				if utils.IsValidNumber(transfer) {
					m["transfer"] = transfer
				}
			}
			weiboMap = append(weiboMap, m)
		}
		return weiboMap
	}
	return nil
}

func getUserInfoDetail(m map[string]interface{}) map[string]interface{} {
	err := global.Page.Page.FindByID("Pl_Core_UserInfo__6").FindByClass("WB_cardmore").Click()
	if err != nil {
		println("点击用户更多资料信息失败")
	} else {
		time.Sleep(time.Second * 3)
		//基本信息
		eles, err := global.Page.Page.Find(".WB_frame_c").All("li").Elements()
		if err != nil {
			println("获取微博用户基本信息失败")
		} else {
			if eles != nil {
				for i := range eles {
					tit, _ := eles[i].GetElement(api.Selector{"css selector", ".pt_title"})
					det, _ := eles[i].GetElement(api.Selector{"css selector", ".pt_detail"})
					title, _ := tit.GetText()
					detail, _ := det.GetText()
					if strings.Contains(title, "昵称") {
						m["nickname"] = detail
					} else if strings.Contains(title, "真实姓名") {
						m["name"] = detail
					} else if strings.Contains(title, "所在地") {
						m["place"] = detail
					} else if strings.Contains(title, "性别") {
						m["gender"] = detail
					} else if strings.Contains(title, "性取向") {
						m["sexorientation"] = detail
					} else if strings.Contains(title, "感情状况") {
						m["marriage"] = detail
					} else if strings.Contains(title, "生日") {
						m["birthday"] = detail
					} else if strings.Contains(title, "血型") {
						m["blood"] = detail
					} else if strings.Contains(title, "博客地址") {
						m["blog"] = detail
					} else if strings.Contains(title, "简介") {
						m["desc"] = detail
					} else if strings.Contains(title, "邮箱") {
						m["email"] = detail
					} else if strings.Contains(title, "QQ") {
						m["qq"] = detail
					} else if strings.Contains(title, "MSN") {
						m["msn"] = detail
					} else if strings.Contains(title, "标签") {
						m["mydesc"] = detail
					} else {
						m["more"] = convert.ToString(m["more"]) + "{" + title + detail + "},"
					}
				}
			}
		}
	}
	return m
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
