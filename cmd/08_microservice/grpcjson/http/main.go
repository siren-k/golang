package main

import (
	"fmt"
	"golang/internal/08_microservice/grpcjson/http"
	"golang/internal/08_microservice/grpcjson/keyvalue"
	h "net/http"
)

func main() {
	c := http.Controller{KeyValue: keyvalue.NewKeyValue()}
	h.HandleFunc("/set", c.SetHandler)
	h.HandleFunc("/get", c.GetHandler)

	fmt.Println("Listening on port :3333")
	err := h.ListenAndServe(":3333", nil)
	panic(err)
}
