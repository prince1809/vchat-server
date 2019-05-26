package app

type Option func(s *Server) error

func StoreOverride(override interface{}) Option {
	panic("implement me")
}

type AppOption func(a *App)
type AppOptionCreator func() []AppOption

func ServerConnector(s *Server) AppOption  {
	return func(a *App) {
		a.Srv = s

		a.Log = s.Log

	}
}
