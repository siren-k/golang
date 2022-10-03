package main

// $ go install github.com/axw/gocov/gocov@latest
// $ go install github.com/smartystreets/goconvey@latest
//
// $ gocov test | gocov report
// ok      golang/internal/09_go_test/tools        0.119s  coverage: 100.0% of statements
//
// golang/internal/09_go_test/tools/structs.go      c.example3      100.00% (5/5)
// golang/internal/09_go_test/tools/funcs.go        example         100.00% (2/2)
// golang/internal/09_go_test/tools/funcs.go        @10:16          100.00% (2/2)
// golang/internal/09_go_test/tools                 ----------      100.00% (9/9)
//
// Total Coverage: 100.00% (9/9)

// 이 에제는 goconvey 명령을 테스트에 연결하는 방법을 보여주다. Convey 키워드는 기본적으로 t.Run을 대체하고 goconvey 웹 UI에
// 레이블을 추가하지만 약간 다르게 동작한다. 중청된 Convey 블록이 있다면 다음과 같이 항상 순서대로 재실행된다.
//
// Convey("Outer loop", t, func() {
//     a := 1
//     Convey("Inner loop", t, func() {
//         a = 2
//     })
//     Convey("Inner loop2", t, func() {
//         fmt.Println(a)
//     })
// })
//
// goconvey 명령을 사용하는 앞의 코드는 1을 출력한다. 여기서 내장 t.Run을 대신 사용했다면 2를 출력했을 것이다. 다시 말해
// Go의 t.Run 테스트는 순차적으로 실행되며 반복되지 않는다. 이런 동작을 설정 코드를 외부 Convey 블록에 넣는데 유용할 수 있다.
// 하지만 goconvey 명령과 t.Run 모두를 사용해야 할 때는 이 차이점을 명심하는 것이 중요하다.
//
// Convey 어설션을 사용하면 웹 UI 및 추가 통계 자료에서 성고에 대한 확인 표시가 나타난다. 또한 검사의 크기를 한 줄로 줄일 수
// 있고 사용자 정의 어설션을 생성하는 것도 가능하다.
//
// goconvey 웹 인터페이스를 그대로 두고 알림을 켜놓은 상태로 두면, 코드를 저장할 때 테스트가 자동으로 실행되며 테스트 범위가
// 늘어나거나 줄어들 때뿐만 아니라 빌드에 실패할 때도 알림을 받는다.
//
// 어셜션, 테스트 실행기, 웹UI라는 세 가지 도구는 모두 독립적으로 사용하거나 함께 사용할 수 있다.
//
// gocov는 더 높은 테스트 범위로 작업할 때 유용할 수 있으며, 테스트 범위가 부족한 함수를 빠르게 식별할 수 있고 테스트 범위 보고서를
// 자세히 살펴볼 수 있도록 도와준다. gocov는 또한 github.com/matm/gocov-html 패키지를 사용해 Go 코드와 함께 제공되는
// 대체 HTML 보고서를 생성하는데 사용할 수도 있다.
func main() {

}
