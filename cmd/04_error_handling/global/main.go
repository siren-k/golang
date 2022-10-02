package main

import (
	"golang/internal/04_error_handling/global"
)

// 전역 패키지 수준 객체를 사용하는 일반적인 패턴은 전역 변수를 노출시키지 않은 채 함수를 통해 원하는 기능만
// 노출시키는 것이다. 일반적으로 로거 객체를 필요로 하는 패키지를 위해 전역로거의 사본을 반환하는 함수를
// 포함시킬 수도 있다.
//
// sync.Once 타입은 새로 도입된 구조체다. 이 구조체는 Do 함수와 함께 코드에서 한 번만 실행된다.
// sync.Once를 초기화 코드에서 사용하고 Init 함수가 한 번 이상 호출되면 오류를 발생시킨다. 예제에서는
// 전역 로그에 매개변수를 전달할 때 내장 init() 함수 대신 사용자 정의 Init 함수를 사용했다.
func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
