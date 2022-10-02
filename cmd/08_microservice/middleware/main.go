package main

import (
	"fmt"
	"golang/internal/08_microservice/middleware"
	"log"
	"net/http"
	"os"
)

// $ curl http://localhost:3333
// success
// $ curl http://localhost:3333
// success

// GOROOT=/usr/local/go #gosetup
// GOPATH=/Users/benjamin/.golang #gosetup
// /usr/local/go/bin/go build -o /private/var/folders/c7/9nkv3bcd42n6szhb2vfnfdfr0000gn/T/GoLand/___go_build_golang_cmd_08_microservice_middleware golang/cmd/08_microservice/middleware #gosetup
// /private/var/folders/c7/9nkv3bcd42n6szhb2vfnfdfr0000gn/T/GoLand/___go_build_golang_cmd_08_microservice_middleware
// Listening on port :3333
// started request to / with id 100
// completed request to / with id 100 in 16.917µs
// started request to / with id 101
// completed request to / with id 101 in 88.584µs

// 미들웨어는 로깅, 메트릭(metric) 수집, 분석과 같은 간단한 작업을 수행하는데 사용할 수 있다. 또한 미들웨어는 요청마다 동적으로
// 변수에 데이터를 설정할 수도 있다. 예를 들면, 이 예제에서 했듯이 ID를 설정하거나 ID를 생성하는 요청에서 X-Header를 수집하는
// 작업을 수행할 수 있다. 또 다른 ID 전략은 모든 요청에 대해 UUID(Universal Unique Identifier)를 생성하는 것이다.
// 모든 요청에 대해 UUID를 생성하면, 여러 마이크로서비스가 응답을 만드는 경우에 여러 로그 메시지를 쉽게 연결시키고 서로 다른
// 애플리케이션 각 요청을 추적할 수 있다.
//
// 컨텍스트 값을 활용해 작업할 때는 미들웨어의 순서를 고려하는 것이 중요하다. 일반적으로 미들웨어가 서로 의존하지 않도록 만드는 것이 좋다.
// 이번 예제로 들면, 로깅 미들웨어 자체에서 UUID를 생성하ㅡㄴ 것이 더 좋다. 하지만 이 에제는 미들웨어를 계층화하고 main.go에서
// 이 미들웨어를 초기화하는 방법에 대한 기본적인 내용을 제공하는데 중점을 뒀다.
func main() {
	// 아래에서부터 위로 적용한다.
	h := middleware.ApplyMiddleware(
		middleware.Handler,
		middleware.Logger(log.New(os.Stdout, "", 0)),
		middleware.SetID(100),
	)

	http.HandleFunc("/", h)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
