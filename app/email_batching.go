package app

import (
	"github.com/prince1809/vchat-server/model"
	"sync"
)

type batchedNotification struct {
}

type EmailBatchingJob struct {
	server               *Server
	newNotifications     chan *batchedNotification
	pendingNotifications map[string][]*batchedNotification
	task                 *model.ScheduledTask
	taskMutex            sync.Mutex
}
