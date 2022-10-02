package main

import (
	"fmt"
	errwrap2 "golang/internal/04_error_handling/errwrap"
)

// pkg/errors 패키지는 매우 유용한 도구이다. 이 패키지를 이용하면 반환되는 모든 오류를
// 래핑해 로깅 및 오류 디버깅 과정에서 추가 정보를 제공할 수 있다. 오류가 발생했을 때
// 전체 스택 추적(stack trace) 정보를 출력하거나 오류를 출력할 때 오류에 접두사만
// 추가할 수 있을 정도로 유연하다. 또한 래핑된 nil은 nil 값을 반환하기 때문에 코드를
// 정리할 수도 있다. 다음 예제를 살펴보자
//
// func RetError() error {
//     err := ThisReturnAnError()
//     return errors.Wrap(err, "This only does something if err != nil")
// }
//
// 경우에 따라서는 오류를 반환하기 전에 먼저 오류가 nil인지 확인하지 않아도 된다. 이 예제는
// 오류를 래핑하고 래핑을 해제하는 방법과 기본적인 스택 추적 기능을 보여준다. 이 패키지에 대한
// 문서는 부분 스택 출력 등 유용한 다른 에제도 제공한다. 이 라이브러리는 만든 데이브 체니(Dave Cheney)는
// 여려 유용한 블로그와 관련 주제에 대한 내용도 작성했다.
// https://dave.chency.net/2016/04 에서 더 자세한 정보를 확인할 수 있다.
func main() {
	errwrap2.Wrap()
	fmt.Println()
	errwrap2.Unwrap()
	fmt.Println()
	errwrap2.StackTrace()
}
