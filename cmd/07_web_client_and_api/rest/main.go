package main

import (
	"golang/internal/07_web_client_and_api/rest"
)

// 이 코드는 인증과 같은 로직을 숨기는 방법과 Transport 인터페이스를 사용해 토큰 갱신을 처리하는 방법을 보여준다.
// 또한, 메소드(함수)를 통해 API 호출을 노출시키는 방법도 보여준다. 사용자 API 등을 사용하지 않고 구현한다면
// 다음과 같은 메소드가 필요하다.
//
// type API interface {
//     GetUsers() (Users, error)
//     CreateUser(User) error
//     UpdateUser(User) error
//     DeleteUser(User)
// }
//
// 6장, '데이터베이스와 저장소의 모든 것'을 읽었다면, '데이터베이스 트랙잭션 인터페이스 실행하기'라는 제목의 예제와
// 비슷하다는 것을 알 수 있습니다. 이러한 구성, 특히 RoundTripper 인터페이스와 같은 공통 인터페이스를 통한 구성은
// API를 작성할 때 많은 유연성을 제공한다. 또한 이전에 했듯이 최상위 인터페이스를 작성하고 이 인터페이스를 전달하는
// 것이 클라이언트에 직접 전달하는 것보다 유용할 수 있다.
func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
