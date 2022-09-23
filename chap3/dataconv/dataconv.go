package dataconv

import "fmt"

// ShowConv 함수는 일부 타입의 변환 예제를 보여준다.
func ShowConv() {
	// int
	var a = 24
	// float 64
	var b = 2.0

	// 다음 계산을 위해 int를 float64로 변환하기
	c := float64(a) * b
	fmt.Println(c)
	// fmt.Sprintf는 문자열로 변환하는 좋은 방법이다.
	precision := fmt.Sprintf("%.2f", b)

	// 값과 타입 출력하기
	fmt.Printf("%s - %T\n", precision, precision)
}
