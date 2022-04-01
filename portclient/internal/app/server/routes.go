package server

import "net/http"

func (s *Server) routes() {
	s.Router.MethodNotAllowedHandler = s.handleUnsupportedMethod()
	s.Router.HandleFunc("/", s.handleLanding()).Methods(http.MethodGet)
	s.Router.HandleFunc("/ports", s.handleGetPorts()).Methods(http.MethodGet)
}
