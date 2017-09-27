package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/beewit/beekit/redis"
	"github.com/beewit/found/global"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
	"regexp"
	"github.com/beewit/beekit/utils"
	"github.com/beewit/beekit/utils/convert"
)

const (
	QQ_SPIDER        = "QQ_SPIDER"
	QQ_SPIDER_DONE   = "QQ_SPIDER_DONE"
	QQ_SPIDER_FAILED = "QQ_SPIDER_FAILED"
	QQ_SPIDER_DATA   = "QQ_SPIDER_DATA"
)

func addQueue(value string) {
	if !checkDoneQueue(value) && !checkFailedQueue(value) && !checkQueue(value) {
		global.RD.SetSETString(QQ_SPIDER, value)
	}
}
func getQueue() string {
	v, _ := global.RD.GetSETRandStringRm(QQ_SPIDER)
	return v
}
func checkQueue(value string) bool {
	x, _ := global.RD.CheckSETString(QQ_SPIDER, value)
	return x > 0
}

func addDoneQueue(value string) {
	global.RD.SetSETString(QQ_SPIDER_DONE, value)
}
func getDoneQueue(value string) string {
	v, _ := global.RD.GetSETRandStringRm(QQ_SPIDER_DONE)
	return v
}
func checkDoneQueue(value string) bool {
	x, _ := global.RD.CheckSETString(QQ_SPIDER_DONE, value)
	return x > 0
}

func addFailedQueue(value string) {
	global.RD.SetSETString(QQ_SPIDER_FAILED, value)
}
func getFailedQueue(value string) string {
	v, _ := global.RD.GetSETRandStringRm(QQ_SPIDER_FAILED)
	return v
}
func checkFailedQueue(value string) bool {
	x, _ := global.RD.CheckSETString(QQ_SPIDER_FAILED, value)
	return x > 0
}

func addDataQueue(value map[string]interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		println(err.Error())
		return
	}
	x, err := global.RD.SetSETString(QQ_SPIDER_DATA, string(b))
	if err != nil {
		println(err.Error())
	}
	println(x)
}

func getDataQueue() (string, error) {
	return global.RD.GetSETRandStringRm(QQ_SPIDER_DATA)
}

func getData() map[string]interface{} {
	d, err := getDataQueue()
	if err != nil {
		println(err.Error())
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

func saveData() {
	for {
		global.Log.Info("执行数据保存服务")
		m := getData()
		if m != nil {
			iw, _ := utils.NewIdWorker(1)

			qqBase := map[string]interface{}{}
			qqBase_id, _ := iw.NextId()
			qqBase["id"] = qqBase_id
			qqBase["qq"] = m["qq"]
			qqBase["nickname"] = m["nice_name"]
			qqBase["head_img"] = m["head_img"]
			qqBase["url"] = m["url"]

			if m["info"] != nil {
				info, err := convert.Obj2Map(m["info"])
				if err != nil {
					global.Log.Error("info 转换失败:", err.Error())
				} else {
					qqBase["gender"] = info["sex"]
					qqBase["age"] = info["age"]
					qqBase["constell"] = info["astro"]
					qqBase["birthday"] = info["birthday"]
					qqBase["place"] = info["live_address"]
					qqBase["marriage"] = info["marriage"]
					qqBase["blood_type"] = info["blood"]
					qqBase["hometown"] = info["hometown_address"]
					qqBase["profession"] = info["career"]
					qqBase["company"] = info["company"]
					qqBase["company_address"] = info["company_caddress"]
					qqBase["detail_address"] = info["caddress"]
				}
			}

			qqBase["ct_time"] = utils.CurrentTime()
			qqBase["ut_time"] = utils.CurrentTime()
			_, err := global.DB.InsertMap("qq_user_base_info", qqBase)
			if err != nil {
				global.Log.Error("qq_user_base_info 数据保存失败:", err.Error())
			} else {
				//QQ动态
				if m["tell"] != nil {
					sayMapList := []map[string]interface{}{}
					tell, err := convert.Obj2ListMap(m["tell"])
					if err != nil {
						global.Log.Error("tell QQ 动态 转换失败:", err.Error())
					} else {
						for i := 0; i < len(tell); i++ {
							qqSay_id, _ := iw.NextId()
							sayMap := map[string]interface{}{}
							comments := ""
							if tell[i]["tell_comments"] != nil {
								b, err := json.Marshal(tell[i]["tell_comments"])
								if err != nil {
									global.Log.Error("tell_comments 评论 Json转换失败:", err.Error())
								} else {
									comments = string(b)
								}
							}
							sayMap["id"] = qqSay_id
							sayMap["qq_base_id"] = qqBase_id
							sayMap["say"] = tell[i]["tell_text"]
							sayMap["content"] = tell[i]["tell_summary"]
							sayMap["phone"] = tell[i]["tell_phone"]
							sayMap["browse"] = tell[i]["tell_browse"]
							sayMap["comments"] = comments
							sayMap["say_date"] = tell[i]["tell_date"]
							sayMap["ct_time"] = utils.CurrentTime()
							sayMapList = append(sayMapList, sayMap)
						}
					}
					if sayMapList != nil && len(sayMapList) > 0 {
						_, err = global.DB.InsertMapList("qq_say", sayMapList)
						if err != nil {
							global.Log.Error("qq_say 数据保存失败:", err.Error())
						}
					}
				}
			}
		}
		time.Sleep(time.Second * 5)

	}
}

func main() {
	go saveData()
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		"--start-maximized",
		"--disable-infobars",
		"--app=https://i.qq.com/?rd=1",
		"--webkit-text-size-adjust"}))
	driver.Start()
	page, err := driver.NewPage()
	if err != nil {
		println(err.Error())
	} else {
		flog, ce := setCookieLogin(page, "qqzone"+global.QZONEUName)
		if ce != nil {
			println(ce.Error())
			return
		}
		if flog {
			page.Navigate("https://user.qzone.qq.com/" + global.QZONEUName)
			time.Sleep(time.Second * 3)
			src, ei := page.Find(".head-avatar img").Attribute("src")
			if ei != nil {
				println(ei.Error())
			}
			if src != "" {
				println("头像", src)
			} else {
				flog = false
			}
		}
		if !flog {
			page.Navigate("https://i.qq.com/?rd=1")
			time.Sleep(time.Second * 3)
			iframe, ee := page.Find("#login_frame").Elements()
			if ee != nil {
				println(ee.Error())
			}
			e2 := page.SwitchToRootFrameByName(iframe[0])
			if e2 != nil {
				println(e2.Error())
				return
			}
			text, e3 := page.Find("#switcher_plogin").Text()
			if e3 != nil {
				println(e3.Error())
				return
			}
			println("登陆按钮", text)
			e4 := page.Find("#switcher_plogin").Click()
			if e4 != nil {
				println("登陆失败", e4.Error())
				return
			}
			page.FindByID("u").Fill(global.QZONEUName)
			time.Sleep(time.Second * 1)
			page.FindByID("p").Fill(global.QZONEPwd)
			time.Sleep(time.Second * 2)
			page.FindByID("login_button").Click()
			time.Sleep(time.Second * 2)
			page.SwitchToParentFrame()
			time.Sleep(time.Second * 3)

			iframe, ef := page.Find("#login_frame").Elements()
			if ef != nil {
				println(ef.Error())
			}
			if iframe != nil && len(iframe) > 0 {
				eef := page.SwitchToRootFrameByName(iframe[0])
				if eef != nil {
					println(eef.Error())
				}
			}
			for {
				//是否有验证码
				html, _ := page.HTML()
				if strings.Contains(html, "请您输入下图中的验证码") {
					println("等待输入验证码")
					continue
				}
				if strings.Contains(html, "QQ手机版授权") {
					println("等待QQ手机版授权")
					continue
				}
				if strings.Contains(html, "安全登录") {
					println("等待手工登陆")
					continue
				}
				time.Sleep(time.Second * 1)
				break
			}

			c, err5 := page.GetCookies()
			if err5 != nil {
				println("登陆失败", e4.Error())
				return
			}
			cookieJson, _ := json.Marshal(c)
			//println("cookie", string(cookieJson[:]))
			redis.Cache.SetString("qqzone"+global.QZONEUName, string(cookieJson[:]))

		}
		////数据
		count := 0
		////初步获取QQ号码
		getHome(page, count)
		//page.SwitchToParentFrame()
		//获取他人个人资料，个人动态，群等
		getQQ(page, global.QZONEUName)

	}
}

func getQQ(page *agouti.Page, qq string) {
	addDoneQueue(qq)
	println("爬取的QQ", qq)
	url := "https://user.qzone.qq.com/" + qq
	page.Navigate(url)
	time.Sleep(time.Second * 3)
	m := map[string]interface{}{}
	html, _ := page.HTML()
	println(len(html))
	if !strings.Contains(html, "申请访问") && !strings.Contains(html, "不符合互联网相关安全规范") && !strings.Contains(html, "对方未开通空间") && !strings.Contains(html, "暂不支持非好友访问") && !strings.Contains(html, "您访问的页面找不回来了") {
		nice_name, err := page.Find(".head-detail .user-name").Text()
		if err != nil {
			println(err.Error())
		}
		m["nice_name"] = nice_name
		m["qq"] = qq
		m["head_img"], _ = page.FindByID("QM_OwnerInfo_Icon").Attribute("src")
		m["url"] = url
		//个人动态
		tell := getTell(page, qq)
		m["tell"] = tell
		//个人资料
		info := getInfo(page, qq)
		m["info"] = info
		addDataQueue(m)
	}
	newQQ := getQueue()
	if newQQ != "" {
		getQQ(page, newQQ)
	}
}

func SwitchFrame(page *agouti.Page, frameSeletor string, f func()) {
	iframe, ee := page.Find(frameSeletor).Elements()
	if ee != nil {
		println(ee.Error())
	}
	if len(iframe) > 0 {
		ee2 := page.SwitchToRootFrameByName(iframe[0])
		if ee2 != nil {
			println(ee.Error())
		}
		f()
		page.SwitchToParentFrame()
	} else {
		println("******     iframe   切换失败")
	}
}

//个人动态
func getTell(page *agouti.Page, qq string) []map[string]interface{} {

	page.Navigate("https://user.qzone.qq.com/" + qq + "/311")
	time.Sleep(time.Second * 3)

	iframe, ee := page.Find(".app_canvas_frame").Elements()
	if ee != nil {
		println(ee.Error())
	}
	if len(iframe) > 0 {
		ee2 := page.SwitchToRootFrameByName(iframe[0])
		if ee2 != nil {
			println(ee.Error())
		}
	}
	time.Sleep(time.Second * 3)
	saveQQ(page)
	ms := []map[string]interface{}{}
	eles, _ := page.Find("#host_home_feeds li").Elements()
	if eles != nil {
		for i := range eles {
			m := map[string]interface{}{}
			//说说
			e, _ := eles[i].GetElement(api.Selector{"css selector", ".f-single-content .f-info"})
			if e != nil {
				m["tell_text"], _ = e.GetText()
			}
			//图片或转发或第三方链接
			e, _ = eles[i].GetElement(api.Selector{"css selector", ".f-single-content .qz_summary"})
			if e != nil {
				m["tell_summary"], _ = e.GetText()
				//手机型号
				e, _ = eles[i].GetElement(api.Selector{"css selector", ".phone-style"})
				if e != nil {
					m["tell_phone"], _ = e.GetText()
				}
			}
			//浏览量
			e, _ = eles[i].GetElement(api.Selector{"css selector", "a.qz_feed_plugin"})
			if e != nil {
				m["tell_browse"], _ = e.GetText()
			}
			//点赞
			es, _ := eles[i].GetElements(api.Selector{"css selector", ".f-like-list a.q_namecard"})
			if es != nil {
				var us []map[string]string
				for j := range es {
					u, _ := es[j].GetText()
					h, _ := es[j].GetAttribute("href")
					us = append(us, map[string]string{u: h})
				}
				if len(us) > 0 {
					m["tell_like_users"] = us
				}
			}
			//评论
			es, _ = eles[i].GetElements(api.Selector{"css selector", ".comments-content"})
			if es != nil {
				var cs []string
				for j := range es {
					u, _ := es[j].GetText()
					cs = append(cs, u)
				}
				if len(cs) > 0 {
					m["tell_comments"] = cs
				}
			}
			ms = append(ms, m)
		}
	} else {
		html, _ := page.Find("#msgList").Text()
		println(len(html))

		eles, _ := page.Find("#msgList").All(".feed").Elements()
		if eles != nil {
			for i := range eles {
				m := map[string]interface{}{}
				//说说
				e, _ := eles[i].GetElement(api.Selector{"css selector", ".bd pre"})
				if e != nil {
					m["tell_text"], _ = e.GetText()
				}
				//图片或转发或第三方链接
				e, _ = eles[i].GetElement(api.Selector{"css selector", ".rt_content"})
				if e != nil {
					m["tell_summary"], _ = e.GetText()
					//手机型号
					e, _ = eles[i].GetElement(api.Selector{"css selector", ".custom-tail"})
					if e != nil {
						m["tell_phone"], _ = e.GetText()
					}
				}
				//发表时间
				e, _ = eles[i].GetElement(api.Selector{"css selector", ".bgr3>.ft .info a"})
				if e != nil {
					m["tell_date"], _ = e.GetText()
				}
				//点赞的人
				es, _ := eles[i].GetElements(api.Selector{"css selector", ".feed_like a"})
				if es != nil {
					var us []map[string]string
					for j := range es {
						u, _ := es[j].GetText()
						h, _ := es[j].GetAttribute("href")
						us = append(us, map[string]string{u: h})
					}
					if len(us) > 0 {
						m["tell_like_users"] = us
					}
				}
				//浏览量
				e, _ = eles[i].GetElement(api.Selector{"css selector", "a.qz_feed_plugin"})
				if e != nil {
					m["tell_browse"], _ = e.GetText()
				}
				//评论
				es, _ = eles[i].GetElements(api.Selector{"css selector", ".comments_content"})
				if es != nil {
					var cs []string
					for j := range es {
						u, _ := es[j].GetText()
						cs = append(cs, u)
					}
					if len(cs) > 0 {
						m["tell_comments"] = cs
					}
				}
				ms = append(ms, m)

			}
		}
	}
	page.SwitchToParentFrame()
	return ms
}

//个人资料
func getInfo(page *agouti.Page, qq string) map[string]string {
	page.Navigate("https://user.qzone.qq.com/" + qq + "/1")
	time.Sleep(time.Second * 3)
	iframe, ee := page.Find(".app_canvas_frame").Elements()
	if ee != nil {
		println(ee.Error())
	}
	if len(iframe) <= 0 {
		return nil
	}
	ee2 := page.SwitchToRootFrameByName(iframe[0])
	if ee2 != nil {
		println(ee2.Error())
	}

	text, ee3 := page.FindByID("info_preview").Text()
	if ee3 != nil {
		println(ee3.Error())
		page.FindByID("info_tab").Click()
	}
	if text == "" {
		page.FindByID("info_tab").Click()
	}
	time.Sleep(time.Second * 2)

	m := map[string]string{}
	m["sex"], _ = page.FindByID("sex").Text()
	m["age"], _ = page.FindByID("age").Text()
	m["birthday"], _ = page.FindByID("birthday").Text()
	m["astro"], _ = page.FindByID("astro").Text()
	m["live_address"], _ = page.FindByID("live_address").Text()
	m["marriage"], _ = page.FindByID("marriage").Text()
	m["blood"], _ = page.FindByID("blood").Text()
	m["hometown_address"], _ = page.FindByID("hometown_address").Text()
	m["career"], _ = page.FindByID("career").Text()
	m["company"], _ = page.FindByID("company").Text()
	m["company_caddress"], _ = page.FindByID("company_caddress").Text()
	m["caddress"], _ = page.FindByID("caddress").Text()
	page.SwitchToParentFrame()
	return m
}

func saveQQ(page *agouti.Page) {
	html, _ := page.HTML()
	reg := regexp.MustCompile("http(s)?://user.qzone.qq.com/\\d+")
	strs := reg.FindAllString(html, -1)
	for i := 0; i < len(strs); i++ {
		s := strings.Replace(strs[i], "http://user.qzone.qq.com/", "", -1)
		s = strings.Replace(s, "https://user.qzone.qq.com/", "", -1)
		if s != "" {
			addQueue(s)
		}
	}
}

//个人主页获取QQ信息
func getHome(page *agouti.Page, count int) []map[string]string {
	saveQQ(page)

	list, e6 := page.Find("#feed_friend_list").All(".f-single").Elements()
	if e6 != nil {
		println("获取好友数据失败", e6.Error())
		return nil
	}
	println("总数量", len(list))
	var s string
	var e7 error
	var ele *api.Element
	for i := range list {

		println("---------------------------------------------------------\r\n")

		s, e7 = list[i].GetAttribute("id")
		if e7 != nil {
			println("错误：", e7.Error())
		}
		println("id：", s)
		ele, e7 = list[i].GetElement(api.Selector{"css selector", ".user-pto img"})
		if e7 != nil {
			println("错误：", e7.Error())
		}
		s, e7 = ele.GetAttribute("src")
		println("头像：", s)

		ele, _ = list[i].GetElement(api.Selector{"css selector", ".user-pto a"})
		if e7 != nil {
			println("错误：", e7.Error())
		}
		s, e7 = ele.GetAttribute("href")
		println("空间链接：", s)

		ele, _ = list[i].GetElement(api.Selector{"css selector", ".f-single-content"})
		if e7 != nil {
			println("错误：", e7.Error())
		}
		s, e7 = ele.GetText()
		println("发表内容：", s)

		ele, _ = list[i].GetElement(api.Selector{"css selector", ".qz_feed_plugin"})
		if e7 != nil {
			println("错误：", e7.Error())
		}
		s, e7 = ele.GetText()
		println("浏览量：", s)

		ele, _ = list[i].GetElement(api.Selector{"css selector", ".comments-list"})
		if e7 != nil {
			println("错误：", e7.Error())
		}
		s, e7 = ele.GetText()
		println("评论：", s)

		ele, _ = list[i].GetElement(api.Selector{"css selector", ".user-list"})
		if e7 != nil {
			println("错误：", e7.Error())
		}
		s, e7 = ele.GetText()
		println("点赞：", s)

		println("---------------------------------------------------------\r\n")
	}
	page.RunScript("document.documentElement.scrollTop=document.body.clientHeight;", nil, nil)
	time.Sleep(time.Second * 3)
	count++
	//分页次数
	if count < 3 {
		getHome(page, count)
	}
	return nil
}

func setCookieLogin(page *agouti.Page, key string) (bool, error) {
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
		page.SetCookie(cc)
	}
	return true, nil
}
