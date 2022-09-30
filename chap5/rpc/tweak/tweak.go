package tweak

import "strings"

// StringTweaker는 자기 자신을 뒤집을 수 있는 문자열 타입이다.
type StringTweaker struct{}

// Args는 문자열을 조정할 수 있는 방법에 대한 옵션이다.
type Args struct {
	String  string
	ToUpper bool
	Reverse bool
}

// Tweak 함수는 다음 내용을 준수하는 RPC 라이브러리를 준수한다.
// - 함수의 타입을 내보낸다.(exported)
// - 함수를 내보낸다.
// - 함수는 두 개의 매개변수를 가지며, 둘 모두 내보내기(또는 내장)한 타입이다.
// - 함수의 두 번째 매개변수는 포인터다.
// - 함수는 error 반환 타입을 갖는다.
func (s StringTweaker) Tweak(args *Args, resp *string) error {
	result := strings.ToUpper(args.String)
	if args.ToUpper {
		result = strings.ToUpper(result)
	}
	if args.Reverse {
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = string(runes)
	}
	*resp = result

	return nil
}
