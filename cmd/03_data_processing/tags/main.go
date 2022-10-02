package main

import (
	"fmt"
	"golang/internal/03_data_processing/tags"
)

// 이 예제는 구조체값을 문자열 직렬화 포맷으로 만들고 모든 문자열 필드를 파싱 가능한 포맷으로 직렬화한다.
// 이 예제는 특별한 경우를 다루지 않는다. 특히 문자열은 콜론(:)이나 세미콜론(;) 문자는 포함되면 안 된다.
// 다음은 예제 동작을 요역한 내용이다.
// - 필드가 문자열이면 직렬화(serialization)/역직렬화(deserialization)한다.
// - 필드가 문자열이 아니면, 무시한다.
// - 필드의 문자열 태그가 직렬화 '키(key)'를 가지면, 이 키가 직렬화/역직렬화 환경을 반환한다.
// - 중복은 처리하지 않는다.
// - 구조체 태그가 설정되지 않았으면, 필드 대신 이름을 대신 사용한다.
// - 구조체 태그 값이 하이픈(-)이면, 필드가 문자열이라도 무시한다.
// 주목해야 할 또 다른 점은 리플렉션이 내보내지 않은 값(nonexported value)에는 전혀 동작하지 않는다는 점이다.
func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
