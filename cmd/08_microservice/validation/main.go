package main

import (
	"fmt"
	"golang/internal/08_microservice/validation"
	"net/http"
)

// $ curl "http://localhost:3333/" -X POST -d '{}'
// name is required
// $ curl "http://localhost:3333/" -X POST -d '{"name":"test"}'
// age is required and must be a value greater than 0 and less than 120
// $ curl "http://localhost:3333/" -X POST -d '{"name":"test","age":5}' -v
// Note: Unnecessary use of -X or --request, POST is already inferred.
// *   Trying 127.0.0.1:3333...
// * Connected to localhost (127.0.0.1) port 3333 (#0)
// > POST / HTTP/1.1
// > Host: localhost:3333
// > User-Agent: curl/7.79.1
// > Accept: */*
// > Content-Length: 23
// > Content-Type: application/x-www-form-urlencoded
// >
// * Mark bundle as not supporting multiuse
// < HTTP/1.1 200 OK
// < Date: Sun, 02 Oct 2022 12:11:50 GMT
// < Content-Length: 0
// <
// * Connection #0 to host localhost left intact

// 컨트롤러 구조체에 클로저를 전달해 유효성 검사를 처리했다. 컨트롤러에서 검증이 필요한 모든 입력에 대해 이런 클로저 중
// 하나가 필요하다. 이 접근 방법의 장점은 실시간으로 검증 함수를 테스트하고 변경할 수 있기 때무에 테스트가 휠씬 간단하다는
// 점이다. 또한 단일 함수 서명에 국한되지 않아도 되므로 데이터베이스 연결 등의 검증 함수에 전달할 수 있다.
//
// 이 예제에서 보여주는 다른 예는 Verror라는 타입의 오류를 반환하는 것이다. 이 타입은 사용자에게 표시할 수 있는 검증
// 오류 메시지를 가진다. 이 방법의 한 가지 단점은 한 번에 여러 검증 메시지를 처리하지 못한다는 것이다. 이 문제는 Verror
// 타입이 더 많은 상태를 가질 수 있도록 변경하는 것으로 해결할 수 있다. 예를 들면, ValidatePayload 함수에서 반환하기
// 저에 여러 오류를 저장하기 위해 맵을 사용하도록 변경할 수 있을 것이다.
func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
