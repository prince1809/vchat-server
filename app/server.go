package app

import (
	"context"
	"crypto/ecdsa"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/services/timezones"
	"github.com/prince1809/vchat-server/config"
	"github.com/prince1809/vchat-server/jobs"
	"github.com/prince1809/vchat-server/mlog"
	"github.com/prince1809/vchat-server/plugin"
	"github.com/prince1809/vchat-server/store"
	"github.com/prince1809/vchat-server/utils"
	"github.com/throttled/throttled"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
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

	PushNotificationHub PushNotificationHub

	runJobs bool
	Jobs    *jobs.JobServer

	clusterLeaderListeners sync.Map

	licenseValue       atomic.Value
	clientLicenseValue atomic.Value
	licenseListeners   map[string]func()

	timezones *timezones.Timezones

	newStore func() store.Store

	htmlTemplateWatcher     *utils.HTMLTemplateWatcher
	sessionCache            *utils.Cache
	seenPendingPostIdsCache *utils.Cache
	configListenerId        string
	licenseListenerId       string
	logListenerId           string
	clusterLeaderListenerId string
	configStore             config.Store
	asymmetricSigningKey    *ecdsa.PrivateKey
	postActionCookieSecret  []byte
}

func NewServer(options ...Option) (*Server, error) {
	rootRouter := mux.NewRouter()

	s := &Server{
		RootRouter: rootRouter,
	}

	return s, nil
}

const TIME_TO_WAIT_FOR_CONNECTIONS_TO_CLOSE_ON_SERVER_SHUTDOWN = time.Second

func (s *Server) StopHTTPServer() {
	if s.Server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), TIME_TO_WAIT_FOR_CONNECTIONS_TO_CLOSE_ON_SERVER_SHUTDOWN)
		defer cancel()
		didShutdown := false
		for s.didFinishListen != nil && !didShutdown {
			if err := s.Server.Shutdown(ctx); err != nil {
				mlog.Warn(err.Error())
			}
			timer := time.NewTimer(time.Millisecond * 50)
			select {
			case <-s.didFinishListen:
				didShutdown = true
			case <-timer.C:
			}
			timer.Stop()
		}
		s.Server.Close()
		s.Server = nil
	}
}

func (s *Server) Shutdown() error {
	mlog.Info("Stopping Server...")

	s.RunOldAppShutdown()

	s.StopHTTPServer()
	s.WaitForGoroutines()

	if s.Store != nil {
		s.Store.Close()
	}

	if s.htmlTemplateWatcher != nil {
		s.htmlTemplateWatcher.Close()
	}

	s.RemoveConfigListener(s.configListenerId)
	s.RemoveConfigListener(s.logListenerId)

	s.configStore.Close()

	mlog.Info("Server stopped")
	return nil
}

func (s *Server) WaitForGoroutines() {
	for atomic.LoadInt32(&s.goroutineCount) != 0 {
		<-s.goroutineExitSignal
	}
}
