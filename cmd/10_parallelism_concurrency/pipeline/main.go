package main

import (
	"context"
	"fmt"
	"golang/internal/10_parallelism_concurrency/pipeline"
)

// main 패키지는 열 개의 인코더와 두 개의 프린터로 구성된 파이프라인을 만든다. in 채널에서 20개의 문자열을 큐에 넣고
// out 채널에서 20개의 응답을 대기한다. 메시지가 out 채널에 도달하면, 전체 파이프라인을 성공적으로 통과했다는 것을 나타낸다.
//
// 풀을 연결하는데 NewPipiline 함수를 사용한다. NewPipiline 함수는 적절한 버퍼의 크기로 채널이 생성되고 일부 풀의 출력
// 채널이 다른 풀의 입력 채널에 적절하게 연결되도록 만든다. 또한 in 채널, out 채널의 배열을 사용해 각 워커, 이름이 설정된
// 여러 채널, 채널의 맵에 파이프라인을 연결할 수도 있다. 이를 통해 각 단계마다 로거에 메시지를 전달할 수도 있다.
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	in, out := pipeline.NewPipeline(ctx, 10, 2)

	go func() {
		for i := 0; i < 20; i++ {
			in <- fmt.Sprint("Message", i)
		}
	}()

	for i := 0; i < 20; i++ {
		<-out
	}
}
