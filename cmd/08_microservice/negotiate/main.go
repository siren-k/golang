package main

import (
	"fmt"
	"golang/internal/08_microservice/negotiate"
	"net/http"
)

// $ curl "http://localhost:3333" -H "Content-Type: text/xml"
// <payload><status>Successful</status></payload>
// $ curl "http://localhost:3333" -H "Content-Type: application/json"
// {"status":"Successful"}

// github.com/unrolled/render 패키지는 이 예제를 위해 중요한 역할을 하며, HTML 템플릿 등으로 작업해야 할 경우에
// 입력할 수 있는 다양한 옵션이 있다. 이 예제는 웹 핸들러로 작업할 때 다양한 Content-Type 헤더를 전달하거나 구조체를
// 직접 조작할 방법을 자동 협상에 활용할 수 있다는 점을 보여준다.
//
// 헤더를 수락할 때 비슷한 패턴을 적용할 수 있다. 하지만 이런 헤더는 종종 여러 값을 포함하기 때문에 코드에서 이를
// 고려해야 한다는 점을 주의한다.
func main() {
	http.HandleFunc("/", negotiate.Handler)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
