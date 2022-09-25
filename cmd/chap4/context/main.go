package main

import "golang/chap4/context"

// context 패키지는 database와 HTTP 패키지를 포함한 다양한 패키지로 나타난다. 이 에제에서
// 보여준 방법을 사용하면, context에 로그 필드를 첨부(추가)하고 이를 로깅 목적으로 사용할 수 있다.
// 예제에서 보여준 아이디어는 context가 별도의 함수에 전달될 때 더 많은 필드를 첨부한 후 최종
// 호출 위치에서 로깅과 변수 관련 작업을 수행할 수 있다는 것이다.
// 이 예제는 이전 예제의 logging 패키지에서 볼 수 있는 WithField와 WithFields 함수를
// 모방했다. 이 두 함수는 context에 있는 값 하나를 수정하며 context 사용의
// 다른 이점(취소, 타임아웃, 스레드 안정성)도 제공한다.
func main() {
	context.Initialize()
}
