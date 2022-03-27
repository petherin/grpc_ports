package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"portclient/internal/app/server"
	"portclient/internal/domains/reader/json"
	saver "portclient/internal/domains/repo/grpc"

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
	// gRPC portsClient needed by HTTP server and repo object to save to gRPC service
	// Get service URL
	address := os.Getenv("SVC_URL")
	if address == "" {
		// Default to address that works locally unless given alternative.
		address = "0.0.0.0:50051"
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Println("Dialling port service...")
	conn, err := grpc.Dial(address, opts...)
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

	////////////////////////////////////////
	// Run json reader
	filePath := os.Getenv("FILE_PATH")
	jsonReader, portCh := json.New(filePath)
	wg.Add(1)
	go jsonReader.Run(ctx, &wg)
	////////////////////////////////////////

	////////////////////////////////////////
	// Run saver
	saver := saver.New(portCh, portsClient)
	wg.Add(1)
	go saver.Run(ctx, &wg)
	////////////////////////////////////////

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
