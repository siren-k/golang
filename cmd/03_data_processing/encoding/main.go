package main

import (
	encoding2 "golang/internal/03_data_processing/encoding"
)

// Gob 인코딩은 Go 데이터 타입으 염두에 두고 제작된 스트리밍 포맷이다. Gob 인코딩은 연속된
// 여러 항목을 보내고 인코딩할 때 가장 효율적이다. 단일 항목에 대해서는 JSON과 같은 다른 인코딩
// 포맷이 잠재적으로 더 효율적이며 가볍다. 그럼에도 gob 인코딩을 통해 복잡한 구조체를 마살링하고
// 마샬링된 데이터를 다시 구조체로 별도의 프로세스에서 재구성할 수 있다. 에제에서는 살펴보지 않았지만,
// gob 패키지는 MarshalBinary와 UnmarshalBinary 함수를 사용해 사용자 정의 타입이다.
// 예측하지 못한 타입에서도 동작할 수 있다.
// Base64 인코딩은 GET 요청에서 URL을 통해 데이터를 주고받거나 바이너리 데이터의 문자열 표현 인코딩을
// 생성할 때 유용하다. 대부분의 언어에서 이 포맷을 지원하며 데이터를 주고받는 다른 쪽 끝에서 언마샬링(unmarshaling)을
// 지원한다. 따라서 JSON 포맷이 지원되지 않을 때 JSON 포맷의 데이터를 인코딩하는 방법을 널리 사용한다.
func main() {
	if err := encoding2.Base64Example(); err != nil {
		panic(err)
	}

	if err := encoding2.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := encoding2.GobExample(); err != nil {
		panic(err)
	}
}
