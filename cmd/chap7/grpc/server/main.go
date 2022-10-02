package main

import (
	"fmt"
	"golang/chap7/grpc/greeter"
	"golang/chap7/grpc/server"
	"google.golang.org/grpc"
	"net"
)

// gRPC 서버는 4444번 포트에서 요청을 대기하도록 설정돼 있다. 클라이언트가 연결되면 서버에 요청을 전달하고 서버로부터
// 응답을 받을 수 있다. 요청(request), 응답(response)의 구조와 지원되는 함수는 4번 단계에서 생성한 .proto 파일에 의해
// 결정된다. 실제로 gRPC 서버와 통합할 때는 클라이언트를 자동으로 생성하는데 사용할 수 있는 .proto 파일을 제공해야 한다.
//
// 클라이언트 외에도 protoc 명령은 서버에 대한 스텁(stub)을 생성하며 구현 세부 사항만 작성하면 된다. 생성된 Go 코드에는
// JSON 태그가 있으며, 동일한 구조를 JSON REST 서비스에 재사용할 수 있다. 예제 코드는 안전하지 않은 클라이언트를 설정한다.
// GRPC를 안전하게 처리하면 SSL 인증서를 사용해야 한다.
func main() {
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &server.Greeter{Exclaim: false})
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port :4444")
	grpcServer.Serve(lis)
}
