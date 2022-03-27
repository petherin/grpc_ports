package json

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"portsvc/proto"
)

type JsonReader struct {
	c        *http.Client
	ports    chan map[string]proto.Port
	filePath string
}

// New creates a JsonReader object and a channel to pass back ports as they're read.
// This is unbuffered so will block as each port is processed. Would be better to have
// a buffer so a few ports can be processed without blocking. Would need to be able
// to save more than one port at a time to the service, maybe using gRPC streaming.
func New(filePath string) (JsonReader, chan map[string]proto.Port) {
	portCh := make(chan map[string]proto.Port)
	return JsonReader{
		c:        http.DefaultClient,
		filePath: filePath,
		ports:    portCh,
	}, portCh
}

// Read processes the configured json file and puts ports onto a channel for another process to deal with.
func (r JsonReader) Read() {
	// This goroutine created the channel, so it should be responsibe for closing it.
	defer close(r.ports)

	// Default to path that works locally unless we were given an alternative.
	path := "files/ports.json"
	if len(r.filePath) > 0 {
		path = r.filePath
	}

	file, err := os.Open(path)
	if err != nil {
		log.Panicf("cannot open %s: %s", path, err.Error())
	}

	dec := json.NewDecoder(file)

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		id, ok := key.(string)
		if !ok {
			log.Fatal(fmt.Errorf("cannot convert port key to string"))
		}

		var p proto.Port
		err = dec.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}

		r.ports <- map[string]proto.Port{id: p}
		i++

		log.Printf("%s processed\n", id)
	}

	log.Printf("%d Port(s) loaded", i)
}

// Run starts reading data and waits until the passed ctx is cancelled.
func (r JsonReader) Run(ctx context.Context, wg *sync.WaitGroup) {
	// tell the caller that we've stopped
	defer wg.Done()

	log.Println("JsonReader started")

	go r.Read()

	// Code will block here, so when the data has all been read it'll wait here.
	// Wise to have a way of cancelling this long-running process if the
	// data is still being read but we want to leave this func.
	<-ctx.Done()
	log.Println("reader: caller has told us to stop")

	// Any cleanup needed

	log.Println("reader: gracefully stopped")
}
