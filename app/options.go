package app

type Option func(s *Server) error

func StoreOverride(override interface{}) Option {
	panic("implement me")
}

type AppOption func(a *App)
type AppOptionCreator func() []AppOption
