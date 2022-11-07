package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

func doSma(c pb.CalculatorServiceClient) {
	log.Println("doSma was invoked")

	stream, err := c.Sma(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 4, 7, 2, 19, 4, 6, 32, 7}
		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.SmaRequest{
				Number: number,
			})

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Problem while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new simple moving avarage of...: %.3f\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
