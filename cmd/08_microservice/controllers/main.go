package main

import (
	"fmt"
	"golang/internal/08_microservice/controllers"
	"net/http"
)

// $ curl "http://localhost:3333/set" -X POST -d "value=value"
// $ curl "http:/localhost:3333/get" -X GET
// {"value":"value"}
// $ curl "http://localhost:3333/get/default" -X GET
// {"value":"default"}

// Go는 메소드(함수)가 http.HandleFunc와 같이 타입이 지정된 함수를 만족시키기 때문에 이런 전략이 효과적이다.
// 구조체를 사용하면 main.go에 데이터베이스 연결, 로깅 등 다양한 기능을 주입할 수 있다. 예제에서는 저장소(storage)
// 인터페이스를 삽입했다. 컨트롤러에 연결된 모든 핸들러는 연결된 컨트롤러의 함수와 속성을 사용할 수 있다.
//
// GetValue 함수는 http.handleFunc 서명이 없는 대신에 이 서명을 반환한다. 이것이 바로 클로저를 사용해 상태를
// 주입하는 방법이다. 예제에서는 main.go에 두 개의 경로(라우트)를 정의했다. 하나는 UseDefault를 false로 설정하고,
// 다른 하나는 UseDefault를 true로 설정한다. 여러 경로에 걸쳐 있는 함수를 정의하거나 핸들러가 너무 복잡한 곳에서
// 구조체를 사용할 때 이 방법(클로저 활용)을 활용할 수 있다.
func main() {
	storage := controllers.MemStorage{}
	c := controllers.New(&storage)
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
