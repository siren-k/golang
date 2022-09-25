package main

import (
	"fmt"
	"golang/chap4/panic"
)

// 더 복잡한 미들웨어를 사용해 중첩된 여러 함수를 실행한 후 recover의 지연 호출(defer)을 통해
// 문제를 처리하는 방법을 상상할 수 있을 것이다. recover 안에서는 기본적으로 어떤 작업이든 할 수
// 있지만 로그를 내보내는 것이 일반적이다.
// 대부분의 웹 애플리케이션에서는 패닉이 발생하면 해당 패닉을 포착해 http.Internal.ServerError
// 메시지를 표시하는 것이 일반적이다.
func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
