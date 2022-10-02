package main

import (
	"fmt"
	"golang/internal/08_microservice/proxy"
	"net/http"
)

// 이번 예제에서는 리버스 프록시 애플리케이션을 개발한다. 브라우저에서 http://localhost:3333이 입력되면 모든 트래픽을 구성 가능한
// 호스트로 전달하고, 브라우저로 응답을 전달한다. 최종 결과는 프록시 애플리케이션을 통해 브라우저에서 https://www.golang.org가 표시되는
// 것이다. 이 과정에서 중간 서버를 통해 웹 사이트에 안전하게 접속하기 위해 포트 포워딩과 SSH 터널을 결합해야 한다. 이 예제는 리버스 프록시를
// 처음부터 제작한다. 하지만 이 기능은 net/http/httputil 패키지에서도 제공한다. 이 패키지를 사용하면 들어오는 요청을
// Director func(*http.Request) 함수로 보낼 수 있고, 내보내는 응답을 ModifyResponse func(*http.Response) error 함수로
// 수정할 수 있다. 또한 응답 버퍼링도 지원한다.
//
// 클라이언트와 핸들러에서 Go의 요청 및 응답 객체를 광범위하게 공유할 수 있다. 이 코드는 핸들러 인터페이스를 만족하는 Proxy 구조체로부터 얻어온
// 요청 객체를 받는다. main.go 파일은 다른 곳에서 사용했던 handleFunc 대신 Handle을 사용한다. 요청이 들어오면 이 요청에 대해
// Proxy.BaseURL을 추가한다. 그런 다음, 클라이언트는 이를 배포한다. 마지막으로 ResponseWriter 인터페이스에 응답을 역으로(reverse)
// 복사한다. 여기에는 헤더, 본문, 상태 코드가 모두 포함된다.
//
// 또한, 필요하다면 요청에 대한 기본적인 인증, 토큰 관리 등과 같은 몇 가지 기능을 추가할 수도 있다. 이렇게 하면 프록시가 자바스크립트나 다른
// 클라이언트 애플리켕션의 세션을 관리하는 곳으로 토큰을 관리하는데 유용할 수 있다.
func main() {
	p := &proxy.Proxy{
		Client:  http.DefaultClient,
		BaseURL: "https://www.golang.org",
	}
	http.Handle("/", p)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
