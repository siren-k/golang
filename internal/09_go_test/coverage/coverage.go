package coverage

import "github.com/pkg/errors"

// Coverage 함수는 분기 조건을 갖는 단순한 함수다.
func Coverage(condition bool) error {
	if condition {
		return errors.New("condition was set")
	}
	return nil
}
