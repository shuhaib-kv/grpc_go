package main

import (
	"bo/greet/greetpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("hello")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("faild tom connecr %v", err)

	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild:%v", err)
	}
}
