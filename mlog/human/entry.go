package human

import (
	"github.com/prince1809/vchat-server/mlog"
	"time"
)

type LogEntry struct {
	Time    time.Time
	Level   string
	Message string
	Caller  string
	Fields  []mlog.Field
}
