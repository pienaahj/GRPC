package main

import pb "github.com/pienaahj/grpc/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
