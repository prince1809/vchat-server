package config

import (
	"github.com/prince1809/vchat-server/model"
	"sync"
)

// emitter enables threadsafe registration and broadcasting to configuration listeners
type emitter struct {
	listeners sync.Map
}

// invokeConfigListeners synchronously notifies all listeners about the configuration change.
func (e *emitter) invokeConfigListeners(oldCfg, newCfg *model.Config) {
	e.listeners.Range(func(key, value interface{}) bool {
		Listener := value.(Listener)
		Listener(oldCfg, newCfg)
		return true
	})
}
