package main

import (
	"fmt"
	collections2 "golang/internal/03_data_processing/collections"
)

// Go의 클로저는 매우 강력하다. 에제의 collections 함수는 제네릭이 아니지만, 비교적 작으며 다양한 함수를
// 사용해 최소한의 코드만 추가하면 WorkWith 구조체에 쉽게 적용할 수 있다. 예제를 살펴보면 어디에서도 오류를
// 반환하지 않는다는 사실을 알 수 있을 것이다. 이 함수들은 순수 함수(pure function)다. 즉, 각 함수 호출
// 이후에 리스트를 덮어 쓴다는 것을 제외하면 원래의 리스트에는 어떠한 영향도 미치지 않는다.
// 리스트나 리스트 구조체에 여러 레이어의 수정본을 적용해야 한다면 이 패턴을 사용해 많은 혼란을 피하고 테스트를
// 매우 직관적으로(간단하게) 만들 수 있다. 또한 매우 표현력 있는 코딩 스카일을 위해 맵(map)과 필터(filter)를
// 연결할 수도 있다.
func main() {
	ws := []collections2.WorkWith{
		collections2.WorkWith{"Example", 1},
		collections2.WorkWith{"Example 2", 2},
	}
	fmt.Printf("Initial list: %#v\n", ws)

	// 먼저 리스트의 값을 소문자로 만든다.
	ws = collections2.Map(ws, collections2.LowerCaseData)
	fmt.Printf("After LowerCaseData Map: %#v\n", ws)

	// 그 다음 모든 버전을 증가시킨다.
	ws = collections2.Map(ws, collections2.IncrementVersion)
	fmt.Printf("After IncrementVersion Map: %#v\n", ws)

	// 마지막으로 3보다 작은 모든 버전을 제거한다.
	ws = collections2.Filter(ws, collections2.OldVersion(3))
	fmt.Printf("After OldVersion Filter: %#v\n", ws)
}
