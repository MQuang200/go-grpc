package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/MQuang200/my-grpc-go-server/internal/port"
	"github.com/MQuang200/my-grpc-proto/protogen/hello"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (adapter *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v\n", adapter.grpcPort, err)
	}

	log.Printf("Listening on port %d\n", adapter.grpcPort)
	grpcServer := grpc.NewServer()
	adapter.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, adapter)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d: %v\n", adapter.grpcPort, err)
	}
}

func (adapter *GrpcAdapter) stop() {
	adapter.server.Stop()
}
