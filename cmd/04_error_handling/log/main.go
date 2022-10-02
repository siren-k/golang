package main

import (
	"fmt"
	log2 "golang/internal/04_error_handling/log"
)

// 로거를 초기화하고 log.NewLogger()를 사용해 로그를 전달하거나 log 패키지 수준 로거를 사용해
// 메시지를 로그로 출력할 수 있다. 이 에제의 log.go 파일은 첫 번째 방법을 사용하고 error.go 파일은
// 두 번째 방법을 사용한다. 또한 오류가 최종 목적지에 도달한 후 로그를 기록해야 한다는 것을 보여준다.
// 그렇게 하지 않으면 한 이벤트에 로그가 여러 번 기록될 수 있다.
// 이 접근 방법에는 몇 가지 문제가 있다. 먼저 중간 함수 중 하나에 컨텍스트를 추가하려는 경우가 있을 수 있다.
// 둘때, 많은 변수를 로그에 기록하면 코드가 정리되지 않아 혼란스럽고 읽기 어렵게 만들 수 있다. 다음 예제에서는
// 변수를 로그에 기록할 때 유연함을 제공하는 구조화된 로깅을 살펴본다. 그 다음 예제에서는 전역 패키지 수준
// 로거의 구현도 살펴볼 것이다.
func main() {
	fmt.Println("basic logging and modification of logger:")
	log2.Log()
	fmt.Println("logging 'handled' errors:")
	log2.FinalDestination()
}
