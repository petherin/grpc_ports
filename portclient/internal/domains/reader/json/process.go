package json

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"portsvc/proto"
	"sync"
)

type JsonReader struct {
	c          *http.Client
	ports      chan map[string]proto.Port
	fileReader io.Reader
}

// New creates a JsonReader object.
func New(fileReader io.Reader, portCh chan map[string]proto.Port) JsonReader {
	return JsonReader{
		c:          http.DefaultClient,
		fileReader: fileReader,
		ports:      portCh,
	}
}

// Read processes the configured json file and puts ports onto a channel for another process to deal with.
func (r JsonReader) Read() {
	// This goroutine created the channel, so it should be responsible for closing it.
	defer close(r.ports)

	dec := json.NewDecoder(r.fileReader)

	_, err := dec.Token()
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
