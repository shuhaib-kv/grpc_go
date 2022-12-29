package main

import (
	calculatorpb "bo/calculator/pb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calcuate Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
	doserverstreaming(c)
	// fmt.Printf("created a client %f", c)

}
func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to use unary rpc")
	req := &calculatorpb.SumRequest{
		Num1: 1,
		Num2: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC :%V", err)
	}
	log.Printf("response from greet :%v", res.Sum)
}
func doserverstreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to use prime decomposition severstreaming rpc")
	req := &calculatorpb.PrimenumberDecompositionRequest{
		Number: 12,
	}
	stream, err := c.PrimenumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC :%V", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("something happend :%v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
