package jobs

import (
	"github.com/prince1809/vchat-server/services/configservice"
	"github.com/prince1809/vchat-server/store"
)

type JobServer struct {
	ConfigService configservice.ConfigService
	Store         store.Store
	Workers       *Workers
	Schedulers    *Schedulers
}
