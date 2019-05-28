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

// AddListener adds a callback function to invoke when the configuration is modified.
func (e *emitter) AddListener(listener Listener) string {
	id := model.NewId()

	e.listeners.Store(id, listener)

	return id
}

// RemoveListener removes a callback function using an id returned from AddListener
func (e *emitter) RemoveListener(id string) {
	e.listeners.Delete(id)
}
