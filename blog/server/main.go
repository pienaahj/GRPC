//go:build !test
// +build !test

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/pienaahj/grpc/blog/proto"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var addr string //  = "0.0.0.0:50051"
var URI string

func main() {
	//  Get the config
	URI, addr = config()
	//  Build the mongo client using the config provided
	client, err := mongo.NewClient(options.Client().ApplyURI(URI)) // "mongodb://root:root@localhost:27017/"
	if err != nil {
		log.Fatal(err)
	}
	//  Connect the client
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//  Open the blogdb collection
	collection = client.Database("blogdb").Collection("blog")

	//  listen at address addr
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	// create a new blog server at the listener
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
