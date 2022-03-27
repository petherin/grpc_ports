package main

import (
	"log"
	"net"

	"portsvc/proto"

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

	ports := make(map[string]proto.Port)
	proto.RegisterPortsServer(grpcServer, proto.NewService(ports))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
