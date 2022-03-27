package server

func (s *Server) routes() {
	s.Router.HandleFunc("/", s.handleLanding())
}
