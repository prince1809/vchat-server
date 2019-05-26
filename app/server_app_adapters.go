package app

func (s *Server) RunOldAppShutdown()  {
	s.FakeApp()
}

func (s *Server) FakeApp() *App {
	a := New(
		ServerConnector)
}
