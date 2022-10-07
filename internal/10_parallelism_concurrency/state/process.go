package state

import "github.com/pkg/errors"

// Process 함수는 연산 타입별로 작업을 분기한 뒤 해당 연산을 처리한다.
func Process(wr *WorkRequest) *WorkResponse {
	resp := WorkResponse{Wr: wr}

	switch wr.Operation {
	case Add:
		resp.Result = wr.Value1 + wr.Value2
	case Subtract:
		resp.Result = wr.Value1 - wr.Value2
	case Multiply:
		resp.Result = wr.Value1 * wr.Value2
	case Divide:
		if wr.Value2 == 0 {
			resp.Err = errors.New("divide by 0")
			break
		}
		resp.Result = wr.Value1 / wr.Value2
	}
	return &resp
}
