package app

import (
	"encoding/json"
	"fmt"
	"github.com/beewit/found/global"
	"github.com/pkg/errors"
	"time"
	"strings"
	"github.com/beewit/beekit/utils"
	"github.com/sclevine/agouti/api"
	"github.com/sclevine/agouti"
	"github.com/beewit/beekit/utils/convert"
)

var (
	SPIDER        = "%s_SPIDER"
	SPIDER_DONE   = "%s_SPIDER_DONE"
	SPIDER_FAILED = "%s_SPIDER_FAILED"
	SPIDER_DATA   = "%s_SPIDER_DATA"
)

type Spider struct {
	Table    string
	Name     string
	StartUrl []string
	Domain   string
	Sleep    int64
}

func Start(spider *Spider, downloadFunc func(spider *Spider), saveData func(spider *Spider)) (err error) {
	global.WechatGroupPage, err = global.Driver.NewPage()
	if err != nil {
		global.Log.Info(err.Error())
	}
	if len(spider.StartUrl) > 0 {
		for i := 0; i < len(spider.StartUrl); i++ {
			spider.AddQueue(spider.StartUrl[i])
			//列表页面1小时更新一次
			spider.ExpireKey(spider.StartUrl[i], 60*60)
		}
		global.WechatGroupPage.Navigate(spider.StartUrl[0])
	} else {
		err = errors.New("StartUrl为空")
		return
	}
	go saveData(spider)
	spider.Run(downloadFunc)
	defer spider.Close()
	return
}

func (spider *Spider) Close() {
	global.WechatGroupPage.CloseWindow()
}

func (spider *Spider) Run(downloadFunc func(spider *Spider)) {
	url := spider.getQueue()
	if url != "" {
		global.WechatGroupPage.Navigate(url)
		downloadFunc(spider)
		if spider.Sleep > 0 {
			time.Sleep(time.Duration(spider.Sleep) * time.Second)
		}
		spider.Run(downloadFunc)
	} else {
		//spider.Close()
		//1小时后重启
		time.Sleep(time.Hour)
		spider.Run(downloadFunc)
	}

}

func (spider *Spider) TrimSpace(s string, err error) string {
	if err != nil {
		global.Log.Error(err.Error())
		return ""
	}
	return strings.TrimSpace(s)
}

type Elements struct {
	Ele []*api.Element
}
type Element struct {
	Ele *api.Element
}
type MultiSelection struct {
	Multi *agouti.MultiSelection
}

func (spider *Spider) GetSelectorAll(page *agouti.Page, allSelector string) (*Elements) {
	multi := MultiSelection{}
	multi.Multi = page.All(allSelector)
	return multi.GetElements()
}

func (spider *Spider) GetSelector(selection *agouti.Selection, allSelector string) (*Elements) {
	multi := MultiSelection{}
	multi.Multi = selection.All(allSelector)
	return multi.GetElements()
}

func (multi *MultiSelection) GetElements() (*Elements) {
	ele, err := multi.Multi.Elements()
	if err != nil {
		global.Log.Error(err.Error())
		return nil
	}
	if len(ele) <= 0 {
		return nil
	}
	thisEle := Elements{}
	thisEle.Ele = ele
	return &thisEle
}

func (ele *Elements) Eq(index int) (*Element) {
	if ele == nil || len(ele.Ele) <= index {
		return nil
	}
	thisEle := Element{}
	thisEle.Ele = ele.Ele[index]
	return &thisEle
}

func (ele *Element) GetFirstFind(selector string) (*Element) {
	return ele.GetFind(0, selector)
}

func (ele *Element) GetFind(index int, selector string) (*Element) {
	if ele == nil || ele.Ele == nil {
		return nil
	}
	text, err2 := ele.Ele.GetText()
	if err2 != nil {
		println(err2.Error())
	}
	println(text)
	newEle, err := ele.Ele.GetElements(api.Selector{"css selector", selector})
	if err != nil {
		global.Log.Error(err.Error())
		return nil
	}
	if len(newEle) <= index {
		return nil
	}
	thisEle := Element{}
	thisEle.Ele = newEle[index]
	return &thisEle
}

func (ele *Element) GetAttr(attr string) (value string) {
	if ele == nil || ele.Ele == nil {
		return
	}
	value, err := ele.Ele.GetAttribute(attr)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	value = strings.TrimSpace(value)
	return
}
func (ele *Element) GetText() (value string) {
	if ele == nil || ele.Ele == nil {
		return
	}
	value, err := ele.Ele.GetText()
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	value = strings.TrimSpace(value)
	return
}

func (ele *Elements) GetFirstText() (value string) {
	value = ele.GetText(0)
	return
}
func (ele *Elements) GetText(index int) (value string) {
	if ele == nil || ele.Ele == nil || len(ele.Ele) <= 0 || len(ele.Ele) <= index {
		return
	}
	value, err := ele.Ele[index].GetText()
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	value = strings.TrimSpace(value)
	return
}
func (ele *Elements) GetFirstFind(selector string) (*Elements) {
	return ele.GetFind(0, selector)
}

func (ele *Elements) GetFind(index int, selector string) (*Elements) {
	if ele == nil || ele.Ele == nil || len(ele.Ele) <= 0 || len(ele.Ele) <= index {
		return nil
	}
	newEle, err := ele.Ele[index].GetElements(api.Selector{"css selector", selector})
	if err != nil {
		global.Log.Error(err.Error())
		return nil
	}

	ele.Ele = newEle
	return ele
}

func (ele *Elements) GetFirstAttr(attr string) (value string) {
	value = ele.GetAttr(0, attr)
	return
}
func (ele *Elements) GetAttr(index int, attr string) (value string) {
	if ele == nil || ele.Ele == nil || len(ele.Ele) <= 0 || len(ele.Ele) <= index {
		return
	}
	value, err := ele.Ele[index].GetAttribute(attr)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	value = strings.TrimSpace(value)
	return
}

func GetPushJson(rule string) (*Spider, error) {
	var spider Spider
	err := json.Unmarshal([]byte(rule), &spider)
	if err != nil {
		global.Log.Error("规则解析错误：%s", err.Error())
		return &Spider{}, err
	} else {
		return &spider, nil
	}
}

func (spider *Spider) AddQueue(value string) bool {
	if !spider.checkDoneQueue(value) && !spider.checkFailedQueue(value) && !spider.checkQueue(value) {
		global.RD.SetSETString(fmt.Sprintf(SPIDER, spider.Name), value)
		return true
	}
	global.Log.Info("Redis key：%s 已存在", value)
	return false
}
func (spider *Spider) ExpireKey(key string, seconds int64) bool {
	_, err := global.RD.ExpireKey(fmt.Sprintf(SPIDER, spider.Name), seconds)
	if err != nil {
		return false
	}
	return true
}
func (spider *Spider) getQueue() string {
	v, _ := global.RD.GetSETRandStringRm(fmt.Sprintf(SPIDER, spider.Name))
	return v
}
func (spider *Spider) checkQueue(value string) bool {
	x, _ := global.RD.CheckSETString(fmt.Sprintf(SPIDER, spider.Name), value)
	return x > 0
}

func (spider *Spider) AddDoneQueue(value string) {
	global.RD.SetSETString(fmt.Sprintf(SPIDER_DONE, spider.Name), value)
}
func (spider *Spider) getDoneQueue(value string) string {
	v, err := global.RD.GetSETRandStringRm(fmt.Sprintf(SPIDER_DONE, spider.Name))
	if err != nil {
		global.Log.Error(err.Error())
	}
	return v
}
func (spider *Spider) checkDoneQueue(value string) bool {
	key := fmt.Sprintf(SPIDER_DONE, spider.Name)
	x, err := global.RD.CheckSETString(key, value)
	if err != nil {
		global.Log.Error(err.Error())
	}
	return x > 0
}

func (spider *Spider) AddFailedQueue(value string) {
	global.RD.SetSETString(fmt.Sprintf(SPIDER_FAILED, spider.Name), value)
}
func (spider *Spider) getFailedQueue(value string) string {
	v, err := global.RD.GetSETRandStringRm(fmt.Sprintf(SPIDER_FAILED, spider.Name))
	if err != nil {
		global.Log.Error(err.Error())
	}
	return v
}
func (spider *Spider) checkFailedQueue(value string) bool {
	x, err := global.RD.CheckSETString(fmt.Sprintf(SPIDER_FAILED, spider.Name), value)
	if err != nil {
		global.Log.Error(err.Error())
	}
	return x > 0
}

func (spider *Spider) AddDataQueue(value map[string]interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		global.Log.Info(err.Error())
		return
	}
	x, err := global.RD.SetSETString(fmt.Sprintf(SPIDER_DATA, spider.Name), string(b))
	if err != nil {
		global.Log.Info(err.Error())
	}
	global.Log.Info("【%S】添加数据到Redis结果%v", spider.Name, x)
}

func (spider *Spider) getDataQueue() (string, error) {
	return global.RD.GetSETRandStringRm(fmt.Sprintf(SPIDER_DATA, spider.Name))
}

func (spider *Spider) GetData() map[string]interface{} {
	d, err := spider.getDataQueue()
	if err != nil {
		global.Log.Warning("【%S】无数据，error：", spider.Name, err.Error())
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

func (spider *Spider) SaveData() {
	m := spider.GetData()
	if m != nil {
		count, err := global.DB.Query("SELECT re_time FROM wechat_group WHERE url=? AND domain=? LIMIT 1", m["url"], spider.Domain)
		if err != nil {
			global.Log.Error(err.Error())
		} else {
			if len(count) == 1 {
				if convert.ToString("re_time") != convert.ToString(m["re_time"]) {
					_, err := global.DB.Update("UPDATE wechat_group SET ut_time=?,re_time=? WHERE url=? AND domain=?",
						utils.CurrentTime(), m["re_time"], m["url"], spider.Domain)
					if err != nil {
						global.Log.Error("Table【 %s 】数据修改失败，Error:", spider.Table, err.Error())
					}
				}
			} else {
				_, err := global.DB.InsertMap(spider.Table, m)
				if err != nil {
					global.Log.Error("Table【 %s 】数据保存失败，Error:", spider.Table, err.Error())
				}
			}
		}
		spider.SaveData()
	}
}
