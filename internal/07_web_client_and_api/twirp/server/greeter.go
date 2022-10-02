package server

import (
	"context"
	"fmt"
	"golang/internal/07_web_client_and_api/twirp/greeter"
)

// Greeter는 protoc에서 생성된 인터페이스를 구현한다.
type Greeter struct {
	Exclaim bool
}

// Greet 함수는 grpc Greet을 구현한다.
func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}
	return &greeter.GreetResponse{Response: msg}, nil
}
