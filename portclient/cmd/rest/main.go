package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/petherin/grpc_ports_cli/internal/app/server"
)

func main() {
	const defaultPort = ":8080"

	// Create context that we can pass to go routines and cancel.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// A WaitGroup for the goroutines to tell us they've stopped.
	wg := sync.WaitGroup{}

	svr := server.New(defaultPort)
	wg.Add(1)
	go svr.Run(ctx, &wg)

	// Now all the go routines are running, listen for Ctrl-c.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	var ctrlC bool
	for {
		select {
		case <-c:
			ctrlC = true
			log.Println("main: received Ctrl-c - shutting down")
		}

		if ctrlC {
			break
		}
	}

	// Tell the goroutines to stop
	log.Println("main: telling goroutines to stop")
	cancel()

	// And wait for them to reply back
	wg.Wait()
	log.Println("main: all goroutines have told us they've finished")
}
