package main

import (
	client2 "golang/internal/07_web_client_and_api/client"
)

// net/http 패키지는 DefaultClient 패키지 변수를 노출시키며, Go, Get, POST 등의 내부 작업에 사용된다.
// 예제의 Setup() 함수는 클라이언트를 반환하고 기본 클라이언트를 반환한 클라이언트와 동일한 클라이언트로 설정한다.
// 클라이언트를 설정할 때 대부분의 수정 사항은 전송 과정에서 이루어지기 때문에 RoundTripper 인터페이스만
// 구현하면 된다.
//
// 이 예제는 항상 418 상태 코드를 반환하는 아무런 작업도 수행하지 않는 RoundTripper의 예시를 보여준다.
// 이 예제가 테스트에 어떻게 유용하게 활용될 수 있을지 상상할 수 있을 것이다. 또한 클라이언트를 함수의 매개변수로
// 전달하는 방법, 전달된 클라이언트를 구조체 매개변수로 사용하는 방법, 요청 처리를 위해 기본 클라이언트를 사용하는
// 방법을 보여준다.
func main() {
	// Setup 함수의 첫 번째 인자는 true로 설정하고 두 번째 인자는 false로 설정한다.
	cli := client2.Setup(true, false)
	if err := client2.DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := client2.DoOps(cli); err != nil {
		panic(err)
	}

	c := client2.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	// Setup 함수의 첫 번째 인자를 true로 설정하고 두 번째 인자도 true로 설정한다.
	client2.Setup(true, true)
	if err := client2.DefaultGetGolang(); err != nil {
		panic(err)
	}
}
