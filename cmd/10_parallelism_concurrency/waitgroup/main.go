package main

import (
	"fmt"
	"golang/internal/10_parallelism_concurrency/waitgroup"
)

// 이 에제는 작업이 완료될 때까지 대기할 경우 동기화 메커니즘으로 waitgroup을 사용하는 방법을 보여준다. 본질적으로 waitgroup.Wait()은
// 내부 카운터(counter)가 0에 도달할 때까지 대기한다. waitgroup.Add(int) 메소드는 입력된 양만큼 카운터를 증가시키며 waitgroup.Done()은
// 카운터를 1씩 감소시킨다. 따라서 다양한 고루틴이 waitgroup을 Done()으로 표시하는 동안 비동기적으로 Wait()을 실행해야 한다.
//
// 이 에제에서는 각 HTTP 요청을 전달하기 전에 카운터를 증가시킨 다음, wg.Done() 메소드를 지연 호출시켰으므로 고루틴이 종료할 때마다 카운터를 감소
// 시킬수 있었다. 그런 다움 집계한 결과를 반환하기 전에 모든 고루틴이 종료할 떄까지 대기한다.
//
// 실제로는 오류 및 응답을 전달하는데 채널을 사용하는 것이 더 좋다.
//
// 이런 방법으로 작업을 비동기로 수행하는 경우, 공유 맵 수정과 같은 작업을 할 때 스레드 안정성을 고려해야 한다. 이를 명싱하면, 모든 비동기적 작업을
// 대기시키는데 waitgroup이 유용하게 활용될 것이다.
func main() {
	sites := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://www.google.com/search?q=golang",
	}

	resps, err := waitgroup.Crawl(sites)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resps received:", resps)
}
