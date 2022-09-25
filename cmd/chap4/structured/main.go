package main

import (
	"fmt"
	"golang/chap4/structured"
)

// sirupsen/logrus와 apex/log 패키지 모두 훌룡한 구조화된 로거(structured logger)를 제공한다.
// 두 패키지 모두 여러 이벤트가 발생했을 때 로그를 기록하거나 로그 항목에 필드(정보)를 추가할 수 있는 기능을
// 제공한다. 이는 비교적 간단한데, 예를 들면 logrus의 hook 기능이나 apex의 사용자 정의 핸들러(custom handler)를
// 사용해 모든 로그에 줄 번호와 서비스 이름을 추가하는 것이 비교적 간단하다. 이 기능의 다른 용도로는 서로 다른
// 서비스에서 요청된 사항을 추적하기 위한 tracdID가 있다.
//
// logrus는 hook과 formatter를 분리해 제공하지만 apex는 이 둘을 결합해 제공한다. 이외에도 apex는 오류
// 필드와 오류 추적(tracing)을 추가할 수 있는 WithError와 같은 편의 기능도 제공한다. 두 가지 모두 예제에서
// 살펴봤다. 또한 logrus와 apex 핸들러로 hook을 적용하는 것도 비교적 간단하다. 두 솔루션 모두 색상이 적용된
// ANSI 텍스트 대신 JSON 포맷으로 변환하는 것은 간단하다.
func main() {
	fmt.Println("Logrus:")
	structured.Logrus()

	fmt.Println()
	fmt.Println("Apex:")
	structured.Apex()
}
