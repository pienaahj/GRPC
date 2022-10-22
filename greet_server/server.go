package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"gitbub.com/pienaahj/GRPC/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create a server type
type server struct{}

// Implement unary GreetWithDeadline API
func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline Request received in server %v\n", req.GetGreeting().GetFirstName())
	for i := 0; i < 3; i++ {
		// Check if the client cancelled yet
		if ctx.Err() == context.Canceled {
			fmt.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, fmt.Sprintf("The client cancelled the request: %v\n", req))
		}
		time.Sleep(1 * time.Second)
	}

	fullName := req.Greeting.GetFirstName() + " " + req.Greeting.GetLastName()
	result := "Hello " + fullName
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}
	return res, nil
}

// LongGreet client stream
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Greet stream Request received in server")
	result := ""
	for {
		req, rErr := stream.Recv()
		if rErr == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if rErr != nil {
			log.Fatalf("Could not receive from client: %v\n", rErr)
			return rErr
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

// Implement unary server function
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Request received in server %v", req)
	fullName := req.Greeting.GetFirstName() + " " + req.Greeting.GetLastName()
	result := "Hello " + fullName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

// Implement server streaming function
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet Request received in GreetManyTimes server %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello from gRPC-Server")

	// Create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//Create a new gRPC server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// Check if the server is serving the listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// BiDi Greet server
func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("GreetEveryone stream Request received in server")
	for {
		req, rErr := stream.Recv()
		if rErr == io.EOF {
			return nil
		}
		if rErr != nil {
			log.Fatalf("Error while reading client stream: %v\n", rErr)
			return rErr
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sErr != nil {
			log.Fatalf("Error while sending data to client: %v\n", sErr)
			return sErr
		}
	}
}
