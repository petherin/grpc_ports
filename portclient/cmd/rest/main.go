package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"portclient/internal/app/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"portsvc/proto"
)

func main() {
	const defaultPort = ":8080"

	// Create context that we can pass to go routines and cancel.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// A WaitGroup for the goroutines to tell us they've stopped.
	wg := sync.WaitGroup{}

	///////////////////////////////////
	// gRPC portsClient
	defaultAddress := "0.0.0.0:50051"
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Println("Dialling port service...")
	conn, err := grpc.Dial(defaultAddress, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	portsClient := proto.NewPortsClient(conn)
	///////////////////////////////////

	///////////////////////////////////
	// HTTP server
	svr := server.New(defaultPort, portsClient)
	wg.Add(1)
	go svr.Run(ctx, &wg)
	///////////////////////////////////

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
