package main

import pb "github.com/pienaahj/grpc/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
