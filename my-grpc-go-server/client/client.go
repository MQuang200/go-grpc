package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/MQuang200/my-grpc-proto/protogen/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error when creating new client\n")
	}
	defer c.Close()

	client := hello.NewCalculatorClient(c)
	// callSum(client)
	// callPND(client)
	// callAvg(client)
	callFindMax(client)

	log.Printf("service client %f", client)

}

func callSum(client hello.CalculatorClient) {
	res, err := client.Sum(context.Background(), &hello.SumRequest{
		Num1: 5,
		Num2: 6,
	})

	if err != nil {
		log.Fatalf("error when calling api %v\n", err)
	}

	log.Printf("Result: %v", res.Result)
}

func callPND(client hello.CalculatorClient) {
	stream, err := client.PrimeNumDecompose(context.Background(), &hello.PNDRequest{
		Num: 120,
	})

	if err != nil {
		log.Fatalf("error while calling PND api %v\n", err)
	}

	for {
		result, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Eof\n")
			return
		}
		log.Printf("Received: %v\n", result)
	}
}

func callAvg(client hello.CalculatorClient) {
	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("error when calling Avg api: %v", err)
	}

	requests := []*hello.AvgRequest{
		{Num: 5},
		{Num: 10},
		{Num: 15},
		{Num: 20.6},
		{Num: 25},
	}

	for _, request := range requests {
		err := stream.Send(request)
		if err != nil {
			log.Fatalf("error while callling Avg api: %v", err)
		}
		log.Printf("Sent: %v", request.GetNum())
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while ending Avg api: %v", err)
	}
	log.Printf("Average: %.2f", response.GetAvg())
}

func callFindMax(client hello.CalculatorClient) {
	log.Println("Calling find max api")
	stream, err := client.FindMax(context.Background())
	if err != nil {
		log.Fatalf("call average error %v\n", err)
	}
	waitChannel := make(chan struct{})

	go func() {
		requests := []*hello.FMaxRequest{
			{Num: 5},
			{Num: 10},
			{Num: 30},
			{Num: 20},
			{Num: 25},
			{Num: 70},
		}
		for _, request := range requests {
			err := stream.Send(request)
			if err != nil {
				log.Fatalf("error while sending request %v\n", err)
			}
			time.Sleep(time.Millisecond * 1000)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resposne, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while Recv FindMax %v\n", err)
				break
			}
			log.Printf("Max: %v\n", resposne.GetMax())
		}
		close(waitChannel)
	}()
	<-waitChannel
}
