package config

import (
	"github.com/prince1809/vchat-server/model"
	"strings"
)

// Listener is a callback function invoked when the configuration changes.
type Listener func(oldConfig *model.Config, newConfig *model.Config)

// Store abstracts the act of getting and setting the configuration.
type Store interface {
	// Get fetches the current, cached configuration.
	Get() *model.Config

	// AddListener adds a callback function to invoke when the configuration is modified.
	AddListener(listener Listener) string

	// RemoveListener removes a callback function using an id returned from AddListener.
	RemoveListener(id string)

	// Close cleans up resources associated with the store.
	Close() error
}

// NewStore creates a database or file store given a data source name by which to connect.
func NewStore(dsn string, watch bool) (Store, error) {
	if strings.HasPrefix(dsn, "mysql://") || strings.HasPrefix(dsn, "postges://") {
		return NewDatabaseStore(dsn)
	}
	return NewDatabaseStore(dsn)
}
