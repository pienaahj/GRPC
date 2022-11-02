package main

import (
	"context"
	"io"
	"log"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

func doFactorial(c pb.CalculatorServiceClient) {
	log.Println("doFactorial was invoked")
	req := &pb.FactorialRequest{
		Number: 6,
	}
	stream, err := c.Factorial(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Factorial: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res.Result)
	}
}
