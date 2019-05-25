package app

import "github.com/prince1809/vchat-server/model"

type Hub struct {
	// ConnectionCount should be kept first
	ConnectionCount int64
	app             *App
	ConnectionIndex int
	register        chan *WebConn
	unregister      chan *WebConn
	broadcast       chan *model.WebSocketEvent
}
