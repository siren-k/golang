package main

import (
	"fmt"
	"golang/chap7/twirp/greeter"
	"golang/chap7/twirp/server"
	"net/http"
)

func main() {
	s := &server.Greeter{}
	twirpHandler := greeter.NewGreeterServiceServer(s, nil)

	fmt.Println("Listening on port :4444")
	http.ListenAndServe(":4444", twirpHandler)
}
