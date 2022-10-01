package main

import (
	"fmt"
	"net/http"

	"golang/chap7/async"
)

// 이 예제는 단일 클라이언트를 사용해 팬 아웃(fan-out) 비동기 방식으로  비동기 방식으로 요청을 처리하기 위한 프레임워크를 작성한다.
// 이 클라이언트는 지정한 URL을 최대한 빨리 검색하려고 시도한다. 대부분의 경우 워커 풀 등으로 이를 제한하는 것이 좋다. 또한 이런
// 비동기 Go 루틴을 클라이언트 외부에서 처리하는 것이 좋으며, 특정 저장소나 조회 인터페이스에 대해 처리하는 것이 좋다.
//
// 이 예제는 또한 case 문을 사용해 여러 채널을 분기하는 방법을 보여준다. 처리 로직 연결을 비동기로 실행하기 때문에 작업을 완료할
// 떄까지 대기하는 매커니즘이 필요하다. 예제에서는 main 함수에서 원래 목록에 있는 URL과 동일한 수의 응답과 오류를 읽었을 때만
// 프로그램을 종료한다. 이와 같은 경우에는 애플리케이션에 타임아웃이 필요한지, 작업을 일찍 취소하는 방법으 지원해야 하는지 여부를
// 고려하는 것도 중요하다.
func main() {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"https://www.github.com",
	}
	c := async.NewClient(http.DefaultClient, len(urls))
	async.FetchAll(urls, c)

	for i := 0; i < len(urls); i++ {
		select {
		case resp := <-c.Resp:
			fmt.Printf("Status received for %s: %d\n", resp.Request.URL, resp.StatusCode)
		case err := <-c.Err:
			fmt.Printf("Error received: %s\n", err)
		}
	}
}
