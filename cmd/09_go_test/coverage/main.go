package main

// ref https://www.hahwul.com/2021/07/11/easy-make-golang-test-code-with-gotests
//
// $ brew install gotests
//
// $ gotests -all -w .
//
// $ go test -cover
// PASS
// coverage: 100.0% of statements
// ok      golang/internal/09_go_test/coverage     0.356s
//
// $ go test -coverprofile=cover.out
// PASS
// coverage: 100.0% of statements
// ok      golang/internal/09_go_test/coverage     0.213s
//
// $ go tool cover -html=cover.out -o coverage.html

// go test -cover 명령은 기본 Go 설치와 함께 제공된다. Go 애플리케이션의 테스트 범위 보고서를 수집하는데 이 명령을 사용할 수
// 있다. 또한 테스트 범위 수집 방법 및 HTML 테스트 범위 보고서를 만들 수 있다. 이 도구를 다른 도구로 래핑해 사용하기도 하며,
// 이 방법은 다음 예제에서 다룰 에정이다. 이런 테이블 기반 테스트는 https://github.com/golang/go/wiki/TableDrivenTests에서
// 다루고 있으며, 많은 코드를 추가 작성하 않아도 여러 케이스(case)를 처리할 수 있는 깔끔한 테스트를 만드는 훌룡한 방법이다.
//
// 이 예제는 테스트 코드를 자동으로 생성한 다음, 테스트 범위를 더 넓게 만들기 위해 테스트 케이스를 채우는 것으로 시작한다. 이 때,
// 특히 까다로운 부분은 변수로 선언되지 않은 함수나 메소드가 호출되는 경우다. 예를 들면, 테스트 범위를 늘리기 위해 gob.Encode()에서
// 오류를 반환시키는 것은 까다로울 수 있다. 또한 이 장의 '표준 라이브러리를 사용한 모의 테스트' 절에서 설명했던 방법을 사용하고 패치를
// 하기 위해 var gobEncode = gob.Encode를 사용하는 것도 어려울 수 있다. 이런 이유로 100% 테스트 범위를 목표로 하는 대신에
// 외부 인터페이스를 더 넓은 범위에서 테스트하는 것(다양한 입출력에 대한 테스트)에 중점을 두는 편이 좋으며, 이 장의 'Go를 사용한
// 동작 테스트' 절에서 살펴볼 일부 케이스와 퍼지(fuzzy) 테스트가 유용할 수 있다.
func main() {

}
