package main

import (
	"fmt"
	"golang/chap4/basicerrors"
)

// errors.New, fmt.Errorf, 사용자 정의 error를 사용하는지 여부를 떠나, 코드에서 오류를 처리하지 않은
// 상태로 두지 않아야 한다는 점이 무엇보다 중요하다. 이렇게 다양한 오류의 정의 방법은 많은 유연성을 제공한다.
// 예를 들면, 구조체에 함수를 추가해 오류를 분석하고 호출 함수에서 여러분이 정의한 오류 타입으로 인터페이스를
// 형 변환해 추가한 기능을 사용할 수 있다.
// 인터페이스 자체는 매우 단순하며, 유효한 문자열을 반환해야 한다는 것이 유일한 요구사항이다. 전체적으로 일관된
// 오류 처리가 가능한지만, 다른 애플리케이션들과 잘 동작해야 하는 일부 고수준(high-level) 애플리케이션에는
// 인터페이스를 구조체에 연결하는 것이 유용할 수 있다.
func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
