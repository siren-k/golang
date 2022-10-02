package main

import (
	"golang/internal/06_database_storage/storage"
)

// 이 예제를 설명하는 가장 중요한 함수는 PerformOperations 함수다. 이 함수는 저장소(storage) 인터페이스를
// 매개변수로 받는다. 즉, 이 함수를 수정하지 않고도 사용할 저장소를 동적으로 교체할 수 있다. 예를 들어 저장소에서
// 데이터를 가져오고 수정하기 위한 별도의 API를 저장소에 연결하는 작업은 간딘하다.
//
// 예제에서는 이런 인터페이스에 대한 컨텍스트를 사용해 유연성을 더하고 인터페이스에서 타임아웃도 처리할 수 있도록
// 했다. 애플리케이션 로직을 저장소와 분리하면 다양한 이점이 있다. 하지만 이를 분리시킬 명확한 경계(위치)를 선택하기
// 어려울 수 있으며, 애플리케이션에 따라 매우 다를 수 있다.
func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}