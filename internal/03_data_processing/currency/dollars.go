package currency

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// ConvertStringDollarsToPennies 함수는 문자열(예: 1.00, 55.12 등)로
// 달러의 양을 입력받아 int64로 변환한다.
func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// amount 인자가 float로 변환 가능한지 확인한다.
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// . 문자를 기준으로 값을 나눈다(분할)
	groups := strings.Split(amount, ".")

	// . 문자가 없는 경우에는 다음 코드에서 값을 result에 저장한다.
	result := groups[0]

	// 기본 문자열
	r := ""

	// "." 다음의 데이터를 처리한다.
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}

	// 0으로 채운다. . 문자가 없는 경우
	// 두 개의 0이 채워질 것이다.
	for len(r) < 2 {
		r += "0"
	}

	result += r

	// int로 변환한다.
	return strconv.ParseInt(result, 10, 64)
}
