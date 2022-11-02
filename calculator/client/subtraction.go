package main

import (
	"context"
	"log"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

func doSubtraction(c pb.CalculatorServiceClient) {
	log.Println("doSubtraction was invoked")
	r, err := c.Subtract(context.Background(), &pb.SubtractionRequest{FirstNumber: 1, SecondNumber: 1})

	if err != nil {
		log.Fatalf("Could not subtract: %v\n", err)
	}

	log.Printf("Subtraction result: %d\n", r.Result)
}
