package server

import (
	"context"
	"fmt"
	greeter2 "golang/internal/07_web_client_and_api/grpc/greeter"
)

// Greeter는 protoc에서 생성된 인터페이스를 구현한다.
type Greeter struct {
	Exclaim bool
	greeter2.UnimplementedGreeterServiceServer
}

// Greet 함수는 grpc Greet을 구현한다.
func (g *Greeter) Greet(ctx context.Context, r *greeter2.GreetRequest) (*greeter2.GreetResponse, error) {
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}
	return &greeter2.GreetResponse{Response: msg}, nil
}
