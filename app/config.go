package app


// Removes a listener function by the unique ID returned when AddConfigListener was called
func (s *Server) RemoveConfigListener(id string) {
	s.configStore.RemoveListener(id)
}
