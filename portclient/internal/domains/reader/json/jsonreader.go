package json

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"portsvc/proto"
	"sync"
)

type JsonReader struct {
	c        *http.Client
	Ports    chan map[string]proto.Port
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
		Ports:    portCh,
	}, portCh
}

// Run starts the JsonReader. It reads through the json file and at each port sends
// it on the channel for something else to save to the service.
func (r JsonReader) Run(ctx context.Context, wg *sync.WaitGroup) {
	// tell the caller that we've stopped
	defer wg.Done()

	log.Println("JsonReader started")

	// Default to path that works locally unless we were given an alternative.
	path := "files/ports.json"
	if len(r.filePath) > 0 {
		path = fmt.Sprintf("%sports.json", r.filePath)
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

		r.Ports <- map[string]proto.Port{id: p}
		i++
	}

	// This goroutine created the channel so it is best to close it
	close(r.Ports)
	log.Printf("%d Ports loaded", i)

	<-ctx.Done()
	log.Println("reader: caller has told us to stop")

	// Any cleanup needed

	log.Println("reader: gracefully stopped")
}
