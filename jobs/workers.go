package jobs

import (
	"github.com/prince1809/vchat-server/services/configservice"
	"sync"
)

type Workers struct {
	startOnce     sync.Once
	ConfigService configservice.ConfigService
}
