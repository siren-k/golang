package pool

import (
	"context"
	"log"
)

// Dispatch 함수는 매개변수로 전달된 numWorker 수만큼의 워커를 생성하고
// 작업 및 응답 추가를 위한 함수와 cancel 함수를 반환한다.
// cancel 함수는 반드시 호출돼야 한다.
func Dispatch(numWorker int) (context.CancelFunc, chan WorkRequest, chan WorkResponse) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	in := make(chan WorkRequest, 10)
	out := make(chan WorkResponse, 10)

	for i := 0; i < numWorker; i++ {
		go Worker(ctx, i, in, out)
	}
	return cancel, in, out
}

// Worker 함수는 계속 반복 실행되며 워커 풀의 일부다.
func Worker(ctx context.Context, id int, in chan WorkRequest, out chan WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			log.Printf("worker id: %d, performing %s work\n", id, wr.Op)
			out <- Process(wr)
		}
	}
}
