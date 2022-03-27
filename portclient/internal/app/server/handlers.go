package server

import (
	"context"
	"encoding/json"
	"net/http"
	"portsvc/proto"
)

func (s *Server) handleLanding() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "only GET allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello I'm the rest client"))
	}
}

func (s *Server) handleGetPorts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "only GET allowed", http.StatusMethodNotAllowed)
			return
		}

		req := proto.GetPortsRequest{}
		list, err := s.client.GetPorts(context.Background(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(list.Ports)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
