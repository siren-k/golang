package main

import (
	"fmt"
	"golang/chap3/nulls"
)

// 어떤 값을 마샬링하거나 언마샬링(unmarshaling)할 때 null 값을 표현하는 빠른 방법은 값을
// 포인터로 전환하는 것이다. 포인터네 값을 직접 할달할 수 없으므로(예: --*a: 1등) 이런 값들을
// 설정하는 것이 까다로울 수 있지만, 그렇지 않은 경우에도 유연하게 처리할 수 있다.
func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}
