package global

import (
	"github.com/beewit/beekit/conf"
	"github.com/beewit/beekit/log"
	"github.com/beewit/beekit/utils/convert"
	"github.com/beewit/beekit/redis"
)

var (
	CFG        = conf.New("config.json")
	Log        = log.Logger
	RD         = redis.Cache
	QZONEUName = convert.ToString(CFG.Get("qzone.u"))
	QZONEPwd   = convert.ToString(CFG.Get("qzone.p"))
)
