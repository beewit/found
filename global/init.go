package global

import (
	"github.com/beewit/beekit/conf"
	"github.com/beewit/beekit/log"
	"github.com/beewit/beekit/utils/convert"
	"github.com/beewit/beekit/redis"
	"github.com/beewit/beekit/mysql"
	"github.com/beewit/beekit/utils"
)

var (
	CFG        = conf.New("config.json")
	DB         = mysql.DB
	Log        = log.Logger
	RD         = redis.Cache
	QZONEUName = convert.ToString(CFG.Get("qzone.u"))
	QZONEPwd   = convert.ToString(CFG.Get("qzone.p"))
	SinaWeiboUName = convert.ToString(CFG.Get("sinaWeibo.u"))
	SinaWeiboPwd   = convert.ToString(CFG.Get("sinaWeibo.p"))
	Page       = new(utils.AgoutiPage)
)
