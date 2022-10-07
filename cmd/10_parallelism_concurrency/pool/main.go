package main

import (
	"fmt"
	"golang/internal/10_parallelism_concurrency/pool"
)

// 이 예제는 하나의 입력 채널, 출력 채널에 여러 워커를 생성하고 워커들을 하나의 cancel() 함수에 연결하기 위해 Dispatch()를 사용한다.
// 다른 목적으로 여러 풀을 만들려고 할 때 이 방법을 활용할 수 있다. 예를 들면, 별도의 풀을 사용해 열 개의 crypto와 20개의 비교 워커를
// 생성할 수 있다. 예를 들면, 별도의 풀을 사용해 해시 요청을 워커에 전달하고 응답을 받은 다음, 같은 풀에서 비교 요청을 전달했다. 이로 인해
// 작업을 수행하는 워커가 매번 달라지지만, 워커는 두 유형의 작업을 모두 처리할 수 있다.
//
// 이 접근 방법의 장점은 두 요청 모두 병렬 처리를 허용하며 최대 동시성도 제어할 수 있다는 것이다. 고루틴의 최대 수를 제한하는 것 또한 메모리
// 제한에 중요할 수 있다. 에제에서는 crypto를 선택했다. crypto는 웹 서비스 등에서 새로운 요청마다 새 고루틴을 가동하려고 할 때 CPU나
// 메모리를 최대로 활용하는 좋은 예제 코드다.
func main() {
	cancel, in, out := pool.Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		in <- pool.WorkRequest{
			Op:   pool.Hash,
			Text: []byte(fmt.Sprintf("messages %d", i)),
		}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- pool.WorkRequest{
			Op:      pool.Compare,
			Text:    res.Wr.Text,
			Compare: res.Result,
		}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string: \"%s\"; matched: %v\n", string(res.Wr.Text), res.Matched)
	}
}
