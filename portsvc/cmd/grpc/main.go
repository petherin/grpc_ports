package main

import (
	"log"
	"net"

	"github.com/petherin/grpc_ports_svc/proto"

	"google.golang.org/grpc"
)

func main() {
	const defaultPort = ":50051"

	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	//TODO replace later, pass in test list of ports
	testList := map[string]proto.Port{
		"AEAJM": {
			Name:    "Ajman",
			City:    "Ajman",
			Country: "United Arab Emirates",
			Alias:   nil,
			Regions: nil,
			Coordinates: []float32{
				55.5136433,
				25.4052165,
			},
			Timezone: "Asia/Dubai",
			Unlocs:   []string{"AEAJM"},
			Code:     "52000",
		},
	}

	proto.RegisterPortsServer(grpcServer, proto.NewService(testList))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
