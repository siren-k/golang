package main

import (
	"context"
	"fmt"
	greeter2 "golang/internal/07_web_client_and_api/twirp/greeter"
	"net/http"
)

// twitchtv/twirp RPC 프레임워크는 프로토콜 버퍼(https://developers.google.com/protocol-buffers)를 사용해 모델을 구축하는 등
// gRPC의 다양한 이점을 제공하며 HTTP 1.1을 통한 통신을 허용한다. 또한 JSON을 사용해 통신할 수도 있으므로 curl 명령을 사용해 twirp RPC
// 서비스와 통신할 수 있다.
// 이 에제는 twitchtv/twirp의 다른 기능을 살펴보지 않고, 기본적인 클라이언트-서버 통신에 대한 내용에만 집중한다. 지원되는 기능에 대한 더 자세한
// 내용은 깃허브 페이지(https://github.com/twitchtv/twirp)를 방문해 확인하길 바란다.
//
// 4444번 포트에서 요청을 대기하도록 twitchtv/twirp RPC 서버를 설정했다. gRPC와 마찬가지로 protoc 명령은 다양한 언어에 대한 클라이언트를
// 생성하는데 사용할 수 있다. 예를 들면 Swagger(https://swagger.io/)를 생성할 수 있다.
//
// gRPC와 마찬가지로 .proto 파일에 먼저 모델을 정의하고 Go 바인딩을 생성한 다음, protoc 명령에서 생성한 인터페이스를 구현했다. .proto 파일을
// 사용한 덕분에 gRPC와 twitchtv/twirp 간에 코드를 이식할 수 있다.(두 프레임워크의 고급 기능을 사용하지 않는 한).
func main() {
	// 타임아웃 등, 좀 더 세부적인 제어를 위해 사용자 정의 클라이언트를 전달할 수도 있다.
	client := greeter2.NewGreeterServiceProtobufClient("http://localhost:4444", &http.Client{})

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
