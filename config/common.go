package config

import (
	"github.com/prince1809/vchat-server/model"
	"io"
	"sync"
)

// commonStore enables code sharing between different backing implementations
type commonStore struct {
	emitter

	configLock             sync.RWMutex
	config                 *model.Config
	configWithoutOverrides *model.Config
	environmentOverrides   map[string]interface{}
}

// This function assumes no lock has been acquired, as it acquires a write lock itself.
func (cs *commonStore) load(f io.ReadCloser, needsSave bool, validate func(*model.Config) error, persist func(*model.Config) error) error {

}
