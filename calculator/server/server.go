package main

import (
	calculatorpb "bo/calculator/pb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("recived sum rpc : %v", req)
	firstnum := req.Num1
	secondnum := req.Num2
	sum := firstnum + secondnum
	res := &calculatorpb.SumResponse{
		Sum: sum,
	}
	return res, nil

}

func main() {
	fmt.Println("claculater server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("faild tom connecr %v", err)

	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild:%v", err)
	}
}
