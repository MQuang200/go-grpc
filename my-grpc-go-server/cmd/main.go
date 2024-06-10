package main

import (
	"github.com/MQuang200/my-grpc-go-server/internal/adapter/grpc"
	"github.com/MQuang200/my-grpc-go-server/internal/application"
)

func main() {
	hs := &application.HelloService{}

	grpcAdapter := grpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}
