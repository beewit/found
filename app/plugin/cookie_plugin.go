package plugin

import (
	"net/http/cookiejar"

	"github.com/beewit/found/app/common"
)

type GetCookieFunc func(*common.Request) (*cookiejar.Jar, error)

type CookiePlugin struct {
	getCookieFunc GetCookieFunc
}

func NewCookiePlugin(getCookieFunc GetCookieFunc) *CookiePlugin {
	return &CookiePlugin{getCookieFunc: getCookieFunc}
}

func (this *CookiePlugin) Do(pluginType PluginType, args ...interface{}) {
	if pluginType == BeforeDownloaderType {
		req := args[0].(*common.Request)
		var err error
		req.Jar, err = this.getCookieFunc(req)
		if err != nil {
			req.Error = err
		}
	}
}
