package main

import (
	"bo/greet/greetpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("bo")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	// fmt.Printf("created a client %f", c)

}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to use unary rpc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "shuhaib",
			LastName:  "kv",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC :%V", err)
	}
	log.Printf("response from greet :%v", res.Result)
}
