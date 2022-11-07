package main

import (
	"io"
	"log"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

// mean returns the mean of the slice provided
func mean(slice []int32) float64 {
	log.Printf("from mean, slice: %v", slice)
	var output float64
	var sum int32
	for _, i := range slice {
		sum += i
	}
	log.Printf("slice: %v sum: %v, slice lenght: %d\n", slice, sum, len(slice))
	output = float64(sum / int32(len(slice)))
	return output
}

// sma return the simple moving average of the slice provided with K=3
func sma(s []int32, k int) float64 {
	var (
		output float64
		tmpX   []int32
	)
	if len(s) > k {
		//  delete the first items(extra) from the slice
		tmpX = s[len(s)-k:]
	} else {
		tmpX = s
	}

	output = mean(tmpX) / float64(k)

	return output
}

func (*Server) Sma(stream pb.CalculatorService_SmaServer) error {
	log.Println("Sma was invoked")
	var ma float64
	mX := []int32{}
	k := 3

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		err = stream.Send(&pb.SmaResponse{
			Result: ma,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
		mX = append(mX, req.Number)

		log.Printf("slice sent for processing: %v\n", mX)
		ma = sma(mX, k)
		// ensure only k elements remain
		if len(mX) > k {
			mX = mX[len(mX)-k:]
		}
	}
}
