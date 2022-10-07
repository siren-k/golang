package state

import "context"

// Processor 함수는 처리할 작업을 분기한다.
func Processor(ctx context.Context, in chan *WorkRequest, out chan *WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			out <- Process(wr)
		}
	}
}
