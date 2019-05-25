package plugin

import (
	"github.com/prince1809/vchat-server/model"
	"sync"
)

type apiImplCreatorFunc func(manifest *model.Manifest) API

type Environment struct {
	activePlugins        sync.Map
	pluginHealthStatuses sync.Map
	PluginHealthCheckJob
}
