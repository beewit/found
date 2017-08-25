package plugin

import (
	"github.com/beewit/found/app/common"
)

type ProxyPlugin struct{}

func NewProxyPlugin() *ProxyPlugin {
	return &ProxyPlugin{}
}

func (this *ProxyPlugin) Do(pluginType PluginType, args ...interface{}) {
	if pluginType == BeforeDownloaderType {
		req := args[0].(*common.Request)
		var err error
		req.ProxyUrl, err = common.NewProxy().GetOneProxy(req.Url)
		if err != nil {
			req.Error = err
		}
	}
}
