package state

type op string

const (
	// 값을 더한다.
	Add op = "add"
	// 값을 뺀다.
	Subtract = "sub"
	// 값을 곱한다.
	Multiply = "mult"
	// 값을 나눈다.
	Divide = "div"
)

// WorkRequest 구조체는 두 값(Value1, Value2)에 대한 연산(op)을 처리한다.
type WorkRequest struct {
	Operation op
	Value1    int64
	Value2    int64
}

// WorkResponse 구조체는 연산의 결과와 발생한 오류를 반환한다.
type WorkResponse struct {
	Wr     *WorkRequest
	Result int64
	Err    error
}
