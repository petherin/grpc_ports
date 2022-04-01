package server

import (
	"encoding/json"
	"net/http"
	"portsvc/proto"
)

func (s *Server) handleLanding() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html><html><body>Ports Client<BR><BR>Browse to <a href="http://localhost:8080/ports">http://localhost:8080/ports</a> to get ports.</body></html>`))
	}
}

func (s *Server) handleGetPorts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := proto.GetPortsRequest{}
		list, err := s.client.GetPorts(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(list.Ports)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func (s *Server) handleUnsupportedMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "only GET allowed", http.StatusMethodNotAllowed)
	}
}
