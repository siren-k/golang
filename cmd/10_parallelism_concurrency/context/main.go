package main

import "golang/internal/10_parallelism_concurrency/context"

// 컨텍스트 값을 사용하는 작업을 할 때는 키를 표현하는 새로운 타입을 생성하는 것이 좋다. 이 예제에서는 key 타입을 생성한 다음,
// 가능한 모든 키를 표현하는 상수(const) 값을 선업했다.
//
// Setup() 함수를 사용해 동시에 키-값 쌍의 테이터를 초기화한다. 함수가 컨텍스트를 수정할 때는 일반적으로 컨텍스트를 매개별수로
// 받고 컨텍스트 값을 반환한다. 따라서 함수의 서명은 대부분 다음과 비슷하다.
//
// func ModifyContext(ctx context.Context) context.Context
//
// 경우에 따라 이 메소드는 오류나 context.WithCancel, context.WithTimeout, context.WithDeadline과 같은 cancel()
// 함수를 반환한다. 모든 자식 컨텍스르는 부모의 속성을 상속한다.
//
// 이 예제에서는 두 개의 자식 컨텍스트를 생성했다. 하나는 데드라인을 설정했고, 다른 하나는 타임아웃을 설정했다. 두 컨텍스트에서 임의의
// 범위로 타임아웃을 설정했고, 타임아웃이 발생하면 종료하도록 설정했다. 마지막으로, 전달된 키를 사용해 값을 가져와 출력했다.
func main() {
	context.Exec()
}
