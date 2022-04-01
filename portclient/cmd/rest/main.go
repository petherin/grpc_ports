package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"portclient/internal/app/config"
	"portclient/internal/app/server"
	"portclient/internal/domains/reader/json"
	saver "portclient/internal/domains/repo/grpc"
	"portsvc/proto"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	const defaultPort = ":8080"
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Create context that we can pass to go routines and cancel.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// A WaitGroup for the goroutines to tell us they've stopped.
	wg := sync.WaitGroup{}

	///////////////////////////////////
	// gRPC portsClient. Required by HTTP server and repo object to save to gRPC service.
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	log.Println("Dialling port service")
	conn, err := grpc.Dial(cfg.SvcURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	log.Println("Port service connection established")
	portsClient := proto.NewPortsClient(conn)
	///////////////////////////////////

	////////////////////////////////////////
	// Run saver
	saver, portCh := saver.New(portsClient)
	wg.Add(1)
	go saver.Run(ctx, &wg)
	////////////////////////////////////////

	////////////////////////////////////////
	// Run json reader
	file, err := os.Open(cfg.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonReader := json.New(file, portCh)
	wg.Add(1)
	go jsonReader.Run(ctx, &wg)
	////////////////////////////////////////

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
