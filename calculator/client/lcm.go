package main

import (
	"context"
	"log"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

// doLcm calaculates the least common multiple of the given range
func doLcm(c pb.CalculatorServiceClient) {
	log.Println("doLcm was invoked")
	stream, err := c.Lcm(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	numbers := []int32{60, 84, 108}

	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)

		stream.Send(&pb.LcmRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Least Common Multiple of %v is: %v\n", numbers, res.Result)
}
