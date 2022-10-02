package main

import (
	"golang/internal/07_web_client_and_api/decorator"
)

// 이 예제는 일급(first-class citizen) 객체인 클로저(closure)와 인터페이스의 장점을 활용한다. 인터페이스를
// 구현하는 함수를 작성하는 것이 이 기법의 핵심이다. 이렇게 하면 구조체로 구현한 인터페이스를 함수로 구현한 인터페이스로
// 래핑할 수 있다.
//
// middleware.go 파일은 클라이언트 미들웨어 함수의 두 가지 예제를 포함한다. 좀 더 정교한 인증과 기능을 위해
// 미들웨어를 추가할 수 있도록 확장하는 것도 가능하다. 또한 이 예제는 이전 예제와 결합해 추가 미들웨어에서 확장할 수 있는
// OAuth2 클라이언트를 생성할 수도 있다.
//
// Decorator 함수는 다음을 가능하게 만드는 편의 함수다.
//
// Decorate(RoundTripper, Middleware1, Middleware2, etc)
//
// vs
//
// var t RoundTripper
// t = Middleware1(t)
// t = Middleware2(t)
// etc
//
// 클라이언트를 래핑하는 방법과 비교해 이 방법의 장점은 인터페이스를 분리해 관리할 수 있다는 것이다. 완전한 기능을 갖춘
// 클라이언트를 작성하려면 GET, POST, PostForm과 같은 함수도 구현해야 한다.
func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
