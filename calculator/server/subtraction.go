package main

import (
	"context"
	"log"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

func (*Server) Subtract(ctx context.Context, in *pb.SubtractionRequest) (*pb.SubtractionResponse, error) {
	log.Printf("Subtraction was invoked with %v\n", in)
	return &pb.SubtractionResponse{Result: in.FirstNumber - in.SecondNumber}, nil
}
