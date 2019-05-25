package app

import "github.com/prince1809/vchat-server/model"

type WebSocketHandler interface {
	ServeWebSocket(conn *WebConn, request *model.WebSocketRequest)
}

type WebSocketRouter struct {
	app      *App
	handlers map[string]WebSocketHandler
}


