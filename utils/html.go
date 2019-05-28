package utils

import "sync/atomic"

type HTMLTemplateWatcher struct {
	templates atomic.Value
	stop      chan struct{}
	stopped   chan struct{}
}

func (w *HTMLTemplateWatcher) Close() {
	close(w.stop)
	<-w.stopped
}
