package pipeline

import "context"

// NewPipeline 함수는 워커를 초기화하고 이들을 연결한 다음,
// 파이프라인의 입력과 최종 출력을 반환한다.
func NewPipeline(ctx context.Context, numEncoders, numPrinters int) (chan string, chan string) {
	inEncode := make(chan string, numEncoders)
	inPrint := make(chan string, numPrinters)
	outPrint := make(chan string, numPrinters)

	for i := 0; i < numEncoders; i++ {
		w := Worker{
			in:  inEncode,
			out: inPrint,
		}
		go w.Work(ctx, Encode)
	}

	for i := 0; i < numPrinters; i++ {
		w := Worker{
			in:  inPrint,
			out: outPrint,
		}
		go w.Work(ctx, Print)
	}

	return inEncode, outPrint
}
