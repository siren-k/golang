package pool

import "github.com/pkg/errors"

type op string

const (
	// Hash는 bcrypt 작업 타입이다.
	Hash op = "encrypt"
	// Compare는 bcrypt 비교 작업이다.
	Compare = "decrypt"
)

// WorkRequest는 워커의 요청이다.
type WorkRequest struct {
	Op      op
	Text    []byte
	Compare []byte // 옵션, 선택사항
}

// WorkResponse는 워커의 응답이다.
type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

// Process 함수는 워커 풀 채널에 작업을 전달한다.
func Process(wr WorkRequest) WorkResponse {
	switch wr.Op {
	case Hash:
		return hashWork(wr)
	case Compare:
		return compareWork(wr)
	default:
		return WorkResponse{Err: errors.New("unsupported operation")}
	}
}
