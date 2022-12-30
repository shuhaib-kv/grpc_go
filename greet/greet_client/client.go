package main

import (
	"bo/greet/greetpb"
	"context"
	"fmt"
	"io"
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
	doserverstreaming(c)

}
func doserverstreaming(c greetpb.GreetServiceClient) {
	fmt.Println("server streaming rpc...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "shuhaib",
			LastName:  "kv",
		},
	}
	resstream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("erroe while calling greet :%v", err)
	}
	for {
		msg, err := resstream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading ")
		}
		log.Printf("Response from greet many times:%v", msg.GetResult())
	}
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
