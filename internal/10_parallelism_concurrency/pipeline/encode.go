package pipeline

import (
	"context"
	"encoding/base64"
	"fmt"
)

// Encode 함수는 일반 텍스트를 int로 받고 base64 문자열 인코딩으로
// 인코딩된 문자열(string => <base64 string encoding>)을 출력으로 반환한다.
func (w *Worker) Encode(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			w.out <- fmt.Sprintf("%s => %s", val, base64.StdEncoding.EncodeToString([]byte(val)))
		}
	}
}
