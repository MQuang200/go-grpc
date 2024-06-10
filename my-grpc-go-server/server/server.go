package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/MQuang200/my-grpc-proto/protogen/hello"
	"google.golang.org/grpc"
)

type server struct {
	hello.CalculatorServer
}

func (*server) Sum(ctx context.Context, req *hello.SumRequest) (*hello.SumResponse, error) {
	return &hello.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}, nil
}

func (*server) PrimeNumDecompose(req *hello.PNDRequest, stream hello.Calculator_PrimeNumDecomposeServer) error {
	log.Println("Calling to decompose api")
	N := req.GetNum()
	k := int32(2)

	for N > 1 {
		time.Sleep(time.Millisecond * 1000)
		if N%k == 0 {
			stream.Send(&hello.SumResponse{
				Result: k,
			})
			N = N / k
		} else {
			k++
			log.Printf("K increases to %v", k)
		}
	}
	return nil
}

func (*server) Average(stream hello.Calculator_AverageServer) error {
	var total float32
	var count float32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&hello.AvgResponse{
				Avg: total / count,
			})
		}
		log.Printf("Receive number: %v\n", req.GetNum())
		total += req.GetNum()
		count++
	}
}

func (*server) FindMax(stream hello.Calculator_FindMaxServer) error {
	var max int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while Recv Max %v\n", err)
		}
		num := req.GetNum()
		log.Printf("Received: %v\n", num)
		if num > max {
			max = num
			err := stream.Send(&hello.FMaxResponse{
				Max: max,
			})
			if err != nil {
				log.Fatalf("error while returning Max %v\n", err)
			}
		}

	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9090")

	if err != nil {
		log.Fatalf("Failed to listen on port tcp:%d %v\n", 9090, err)
	}

	s := grpc.NewServer()
	hello.RegisterCalculatorServer(s, &server{})

	log.Println("Server is listening on port 9090")
	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to listen on port tcp:%d %v\n", 9090, err)
	}

}
