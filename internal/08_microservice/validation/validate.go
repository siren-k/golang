package validation

import "github.com/pkg/errors"

// Verror는 검증 과정에서 발생하는 오류이며 사용자에게 반환할 수 있다.
type Verror struct {
	error
}

// Payload는 예제에서 처리하는 값이다.
type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ValidatePayload 함수는 컨트롤러 안에서 클로저를 구현하는 한 가지 방업이다.
func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return Verror{errors.New("name is required")}
	}

	if p.Age <= 0 || p.Age >= 120 {
		return Verror{errors.New("age is required and must be a value greater than 0 and less than 120")}
	}

	return nil
}
