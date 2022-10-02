package main

import (
	"context"
	"fmt"
	greeter2 "golang/internal/07_web_client_and_api/grpc/greeter"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greeter2.NewGreeterServiceClient(conn)
	ctx := context.Background()
	req := greeter2.GreetRequest{Greeting: "Hello", Name: "Reader"}
	resp, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	req.Greeting = "Goodbye"
	resp, err = client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
