package pipeline

import "context"

// Worker는 Work 함수를 호출하는 시점을 결정하는 한 가지 역할을 맏는다.
type Worker struct {
	in  chan string
	out chan string
}

// Job은 워커가 수행할 수 있는 작업이다.
type Job string

const (
	// Print는 모든 입력을 표준 출력(stdout)으로 똑같이 출력한다.
	Print Job = "print"
	// 'encode'라는 입력 값을 base64로 인코딩한다.
	Encode Job = "encode"
)

// Work 함수는 워커를 연결하는 방법을 보여주며, 여기서 작업이 할당된다.
func (w *Worker) Work(ctx context.Context, j Job) {
	switch j {
	case Print:
		w.Print(ctx)
	case Encode:
		w.Encode(ctx)
	default:
		return
	}
}
