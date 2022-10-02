package main

import (
	"fmt"
	"golang/internal/07_web_client_and_api/twirp/greeter"
	"golang/internal/07_web_client_and_api/twirp/server"
	"net/http"
)

func main() {
	s := &server.Greeter{}
	twirpHandler := greeter.NewGreeterServiceServer(s, nil)

	fmt.Println("Listening on port :4444")
	http.ListenAndServe(":4444", twirpHandler)
}
