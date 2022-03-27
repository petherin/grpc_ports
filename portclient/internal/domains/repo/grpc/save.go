package grpc

import (
	"context"
	"log"
	"portsvc/proto"
	"sync"
)

type Repo struct {
	portCh chan map[string]proto.Port
	client proto.PortsClient
}

// New returns a Repo with the ability to save to the ports gRPC service.
func New(portCh chan map[string]proto.Port, client proto.PortsClient) Repo {
	return Repo{
		portCh: portCh,
		client: client,
	}
}

// Run starts waiting for ports on the channel and when it gets one, saves it to the port service.
func (r Repo) Run(ctx context.Context, wg *sync.WaitGroup) {
	// tell the caller that we've stopped
	defer wg.Done()

	log.Println("Saver started...")
	
	for {
		select {
		case port := <-r.portCh:
			// Should only be 1 entry in the map, ranging over it easier than getting key out some other way.
			for k, p := range port {
				req := proto.SavePortRequest{Port: &p, Key: k}
				_, err := r.client.Save(ctx, &req)
				if err != nil {
					log.Fatal(err.Error())
				}
			}

		case <-ctx.Done():
			log.Println("saver: caller has told us to stop")

			// Any cleanup needed

			log.Println("saver: gracefully stopped")
			return
		}
	}
}
