package main

import (
	"fmt"
	"golang/internal/08_microservice/handlers"
	"net/http"
)

// curl "http://localhost:3333/greeting" -X POST -d 'name=Reader;greeting=Goodbye'

// 이 예제에서는 두 개의 핸들러를 설정했다.
//
// 첫 번째 핸들러는 name이라는 GET 매개변수를 가진 GET 요청을 예상한다. curl 명령을 통해 실행하면 일반 텍스트로
// Hello <name>!을 반환한다.
//
// $ curl "http://localhost:3333/name?name=Reader" -X GET
// Hello Reader!
//
// 두 번째 핸들러는 PostForm 요청을 갖는 POST 함수를 예상한다. AJAX 호출 없이 표준 HTML 폼(form)을 사용해
// 응답을 얻을 수 있다. 다른 방법으로 요청 본문에서 JSON을 해석할 수도 있다. 이 과정은 일반적으로 json.Decoder를
// 통해 처리한다. JSON을 처리하는 것도 연습으로 시도해볼 것을 권한다. 마지막으로, 핸들러는 JSON 포맷의 응답을 전송하고
// 모든 헤더를 적절하게 설정한다.
//
// 이 모든 내용을 명시적으로 작성했지만, 다음과 같이 코드를 덜 작성하는 방법이 있다.
// - 응답 처리를 위해 https://github.com/unrolled/render를 사용한다.
// - 이 장의 '웹 핸들러, 요청, ResponseWriter 인스턴스를 활용한 작업' 절에서 언급했던 다양한 웹 프레임워크를 사용해
//   라우트 매개변수 해석, 특정 HTTP 함수로 경로 제한, 안전하게 종료하기 등을 처리한다.
func main() {
	http.HandleFunc("/name", handlers.HelloHandler)
	http.HandleFunc("/greeting", handlers.GreetingHandler)

	fmt.Println("Listening on prot :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
