package server

import (
	"context"
	"log"
	"net/http"
	"portsvc/proto"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router     mux.Router
	HTTPServer *http.Server
	c          *http.Client
	client     proto.PortsClient
}

// New creates a new web server. Takes in a client to talk to the ports gRPC service.
func New(addr string, client proto.PortsClient) Server {
	svr := Server{c: http.DefaultClient, client: client}
	svr.routes()

	svr.HTTPServer = &http.Server{
		Addr:    addr,
		Handler: &svr.Router,
	}

	return svr
}

// Run starts the HTTP server listening
func (s Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	// Tell the caller that we've stopped.
	defer wg.Done()

	log.Printf("HTTP server started on http://localhost%s", s.HTTPServer.Addr)

	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			log.Printf("Listen : %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("server: caller has told us to stop")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ignore error since it will be "Err shutting down server : context canceled"
	s.HTTPServer.Shutdown(shutdownCtx)

	log.Println("server: gracefully stopped")
}
