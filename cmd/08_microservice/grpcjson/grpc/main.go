package main

import (
	"golang/internal/08_microservice/grpcjson/keyvalue"
	"google.golang.org/grpc"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	keyvalue.RegisterKeyValueServer(grpcServer, keyvalue.NewKeyValue())

	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}

	println("Listening on port :4444")
	grpcServer.Serve(lis)
}
