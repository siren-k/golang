package main

import (
	"context"
	"golang/internal/10_parallelism_concurrency/channels"
	"time"
)

// 이 예제는 채널에서 값을 읽거나 채널에 쓸 수 있고 (잠재적으로) 두 가지를 모두 수행할 수 있는 워크 프로세스(worker process)를 실행하는
// 두 가지 방법을 보여준다. 채널에 done이 작성되거나 cancel 함수의 호출 또는 타임아웃을 통해 컨텍스트가 취소되면 워커가 종료된다. context
// 패키지는 'context 패키지 사용하기' 예제에서 더 자세히 다룰 예정이다.
//
// main 패키지는 별도의 함수를 서로 연결하는데 사용한다. main 패키지 덕분에 채널에 공유하지 않는 한 여러 쌍을 설정할 수 있다. 이 외에도
// 같은 채널에 대한 동작을 처리하는 여러 고루틴을 만들 수 있으며, 이 내용은 '워커 풀 디자인 패턴 사용하기' 예제에서 살펴볼 것이다.
//
// 마지막으로 고루틴의 비동기적 특성으로 인해 정리 작업 및 종료 조건 설정이 까자로울 수 있다. 예를 들어 다음 코드는 흔히 저지를 수 있는 실수를
// 보여준다.
//
// select {
//     case <-time.Tick(200 * time.Millisecond):
//     // 이 코드는 다른 case 구문이 선택될 때마다 재설정된다.
// }
//
// select 구분에 Tick을 넣으면 이 case 구문이 실행되는 것을 막을 수 있다. 또한 select 구문에서 트래픽의 우선순위를 지정하는 간단한
// 방법도 없다.
func main() {
	ch := make(chan string)
	done := make(chan bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go channels.Printer(ctx, ch)
	go channels.Sender(ch, done)

	time.Sleep(2 * time.Second)
	done <- true
	cancel()
	time.Sleep(3 * time.Second)
}
