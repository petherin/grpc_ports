package proto

import (
	"context"
	"log"
	"sort"
	"sync"
)

// compile-time check for interface adherence.
var _ PortsServer = new(Service)

type Service struct {
	Ports map[string]Port
	lock  sync.RWMutex
}

// NewService returns a new Service object with passed map.
func NewService(ports map[string]Port) Service {
	return Service{
		Ports: ports,
		lock:  sync.RWMutex{},
	}
}

// Save a port passed in SavePortRequest, return SavePortResponse or error if something went wrong.
func (s Service) Save(ctx context.Context, request *SavePortRequest) (*SavePortResponse, error) {
	key := request.Key
	p := request.Port
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.Ports[key] = *p

	return &SavePortResponse{Success: true}, nil
}

// GetPorts returns all ports or an error if something went wrong.
func (s Service) GetPorts(ctx context.Context, request *GetPortsRequest) (*PortList, error) {
	list := &PortList{}
	list.Ports = make([]*Port, len(s.Ports))

	log.Printf("len ports: %d", len(s.Ports))
	i := 0

	s.lock.Lock()
	defer s.lock.Unlock()
	for _, p := range s.Ports {
		ptrP := p
		list.Ports[i] = &ptrP
		i++
	}

	sort.Slice(list.Ports, func(i, j int) bool {
		portj := *list.Ports[j]
		porti := *list.Ports[i]
		return porti.Unlocs[0] < portj.Unlocs[0]
	})

	return list, nil
}

func (s Service) mustEmbedUnimplementedPortsServer() {}
