package plugin

import (
	"github.com/beewit/found/app/common"
)

type UserAgentPlugin struct{}

func NewUserAgentPlugin() *UserAgentPlugin {
	return &UserAgentPlugin{}
}

func (this *UserAgentPlugin) Do(pluginType PluginType, args ...interface{}) {
	if pluginType == BeforeDownloaderType {
		req := args[0].(*common.Request)
		req.Request.Header.Set("User-Agent", "MSIE (MSIE 6.0; X11; Linux; i686) Opera 7.23")
	}
}
