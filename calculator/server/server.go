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
func (*server) PrimenumberDecomposition(req *calculatorpb.PrimenumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimenumberDecompositionServer) error {
	fmt.Printf("recived primenumber decompodition rpc : %v", req)
	number := req.GetNumber()
	divisor := 2
	for number > 1 {
		if number%int64(divisor) == 0 {
			stream.Send(&calculatorpb.PrimenumberDecompositionResponse{
				PrimeFactor: int64(divisor),
			})
			number = number / int64(divisor)
		} else {
			divisor++
			fmt.Printf("divisor has increased to %v", divisor)
		}
	}
	return nil
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
