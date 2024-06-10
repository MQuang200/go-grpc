package grpc

import (
	"context"

	"github.com/MQuang200/my-grpc-proto/protogen/hello"
)

func (adapter *GrpcAdapter) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	greet := adapter.helloService.GenerateHello(req.Name)
	return &hello.HelloResponse{
		Greet: greet,
	}, nil
}
