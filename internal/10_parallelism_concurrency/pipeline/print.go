package pipeline

import (
	"context"
	"fmt"
)

// Print 함수는 w.in을 출력하고 w.out에 이를 반복한다.
func (w *Worker) Print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			fmt.Println(val)
			w.out <- val
		}
	}
}
