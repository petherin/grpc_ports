package proto

import (
	"context"
	"log"
)

// compile-time check for interface adherence.
var _ PortsServer = new(Service)

type Service struct {
	Ports map[string]Port
}

// NewService returns a new Service object with passed map.
func NewService(ports map[string]Port) Service {
	return Service{
		Ports: ports,
	}
}

// Save a port passed in SavePortRequest, return SavePortResponse or error if something went wrong.
func (s Service) Save(ctx context.Context, request *SavePortRequest) (*SavePortResponse, error) {
	key := request.Key
	p := request.Port
	log.Println("Save called")
	s.Ports[key] = *p
	log.Printf("len ports: %d", len(s.Ports))
	log.Printf("%v", p)
	return &SavePortResponse{Success: true}, nil
}

// GetPorts returns all ports or an error if something went wrong.
func (s Service) GetPorts(ctx context.Context, request *GetPortsRequest) (*PortList, error) {
	list := &PortList{}
	list.Ports = make(map[string]*Port)

	log.Printf("len ports: %d", len(s.Ports))
	for key, p := range s.Ports {
		ptrP := p
		list.Ports[key] = &ptrP
	}

	return list, nil
}

func (s Service) mustEmbedUnimplementedPortsServer() {}
