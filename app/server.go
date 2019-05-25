package app

import (
	"github.com/gorilla/mux"
	"github.com/prince1809/vchat-server/plugin"
	"github.com/prince1809/vchat-server/store"
	"github.com/throttled/throttled"
	"net"
	"net/http"
	"sync"
)

var MaxNotificationsPerChannelDefault int64 = 1000000

type Server struct {
	Store           store.Store
	WebSocketRouter *WebSocketRouter

	// RootRouter is the starting point for all HTTP requests to the server.
	RootRouter *mux.Router

	Server      *http.Server
	ListenAddr  *net.TCPAddr
	RateLimiter *RateLimiter

	didFinishListen chan struct{}

	goroutineCount      int32
	goroutineExitSignal chan struct{}

	PluginsEnvironment     *plugin.Environment
	PluginConfigListenerId string
	PluginsLock            sync.RWMutex

	EmailBatching    *EmailBatchingJob
	EmailRateLimiter *throttled.GCRARateLimiter

	Hubs                        []*Hub
	HubsStopCheckingForDeadlock chan bool
}

func NewServer(options ...Option) (*Server, error) {
	rootRouter := mux.NewRouter()

	s := &Server{
		RootRouter: rootRouter,
	}

	return s, nil
}
