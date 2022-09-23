package math

import (
	"fmt"
	"math"
)

// Examples 함수는 math 패키지에서 제공하는
// 일부 함수의 사용법을 보여준다.
func Examples() {
	// sqrt 예제
	i := 25

	// i는 int 타입이기 때문에 변환한다.
	result := math.Sqrt(float64(i))
	fmt.Println(result) // 25의 제곰근은 5

	// 올림 연산 처리
	result = math.Ceil(9.5)
	fmt.Println(result)

	// 내림 연산 처리
	result = math.Floor(9.5)
	println(result)

	// 또한 math 패키지는 몇 가지 상수(const)도 제공한다.
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}
