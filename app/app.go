package app

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
	"github.com/prince1809/vchat-server/mlog"
	"github.com/prince1809/vchat-server/model"
)

type App struct {
	Srv *Server

	Log             *mlog.Logger
	NotificationLog *mlog.Logger

	T              goi18n.TranslateFunc
	Session        model.Session
	RequestId      string
	IpAddress      string
	Path           string
	UserAgent      string
	AcceptLanguage string
}

func New(options ...AppOption) *App {
	app := &App{}

	for _, option := range options {
		option(app)
	}

	return app
}


// DO NOT CALL THIS
// This is to avoid having to change all the code in cmd/mattermose/commands/* for now
// shutdown should be called directly on the server
func (a *App) Shutdown()  {
	a.Srv.Shutdown()
	a.Srv = nil
}
