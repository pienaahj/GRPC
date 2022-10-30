package main

import pb "github.com/pienaahj/grpc/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
