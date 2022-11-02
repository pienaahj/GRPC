package main

import (
	"fmt"
	"io"
	"log"
	"math"

	pb "github.com/pienaahj/grpc/calculator/proto"
)

// gets the primes of nummer
func getPrime(number int32) map[int32]int {
	count := 1
	divisor := int32(2)
	primeTmp := map[int32]int{}

	for number > 1 {
		if number%divisor == 0 { //  is the current number a prime?
			if _, ok := primeTmp[divisor]; ok { //  check if divisor exist in map
				primeTmp[divisor] = count //  replace it with new count
			} else {
				primeTmp[divisor] = count //  add the divisor with current count
			}
			fmt.Printf("current prime: %v\n", primeTmp)

			number /= divisor // take the current prime out of number
			count++
		} else {
			count = 1
			divisor++
		}
	}
	return primeTmp
}

// processPrimes combines all prime numbers from requests and determine the LCM
func processPrimes(in []int32) int32 {
	lcm := int32(1)
	primeHolder := map[int32]int{}
	log.Printf("Incoming numbers in LCM request: %v\n", in)
	//  find the primes for every number received
	for _, num := range in {
		primeTmp := getPrime(num)
		for k, v := range primeTmp {
			if value, ok := primeHolder[k]; ok { //  store the larger of the powers
				if v > value {
					primeHolder[k] = v
				}
			} else {
				primeHolder[k] = v //  add the prime and power
			}
		}
	}
	// calculate the lcm
	for k, v := range primeHolder {
		lcm *= int32(math.Pow(float64(k), float64(v)))
	}
	log.Printf("LCF returned: %v\n", lcm)
	return lcm
}

func (*Server) Lcm(stream pb.CalculatorService_LcmServer) error {
	log.Println("Lcm was invoked")
	reqX := []int32{}

	for {
		req, err := stream.Recv()
		if req.GetNumber() != 0 {
			reqX = append(reqX, req.GetNumber())
		}
		if err == io.EOF {
			out := processPrimes(reqX)
			return stream.SendAndClose(&pb.LcmResponse{
				Result: out,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
	}
}
