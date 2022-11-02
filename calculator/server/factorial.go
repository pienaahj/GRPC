package main

import (
	"log"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

func (*Server) Factorial(in *pb.FactorialRequest, stream pb.CalculatorService_FactorialServer) error {
	log.Printf("Factorial was invoked with %v\n", in)

	number := in.Number
	result := int64(1)

	if number == 0 {
		stream.Send(&pb.FactorialResponse{
			Result: result,
		})
		return nil
	} else {
		for x := number; x > 0; x-- {
			stream.Send(&pb.FactorialResponse{
				Result: x,
			})
		}
	}

	return nil
}
