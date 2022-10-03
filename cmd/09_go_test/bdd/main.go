package main

// $ go install github.com/cucumber/godog/cmd/godog@latest
//
// $ godog
// Use of godog without a sub-command is deprecated. Please use godog build, godog run or godog version.
// Feature: Bad Method
//
//   Scenario: Good request                           # features/handler.feature:2
//     Given we create a HandlerRequest payload with:
//       | reader |
//       | coder  |
//       | other  |
//     And we POST the HandlerRequest to /hello
//     Then the response code should be 200
//     And the response body should be:
//       | BDD testing reader |
//       | BDD testing coder  |
//       | BDD testing other  |
//
// 1 scenarios (1 undefined)
// 4 steps (4 undefined)
// 154.875µs
//
// You can implement step definitions for undefined steps with these snippets:
//
// func theResponseBodyShouldBe(arg1 *godog.Table) error {
//     return godog.ErrPending
// }
//
// func theResponseCodeShouldBe(arg1 int) error {
//     return godog.ErrPending
// }
//
// func weCreateAHandlerRequestPayloadWith(arg1 *godog.Table) error {
//     return godog.ErrPending
// }
//
// func wePOSTTheHandlerRequestToHello() error {
//        return godog.ErrPending
// }
//
// func InitializeScenario(ctx *godog.ScenarioContext) {
//     ctx.Step(`^the response body should be:$`, theResponseBodyShouldBe)
//     ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
//     ctx.Step(`^we create a HandlerRequest payload with:$`, weCreateAHandlerRequestPayloadWith)
//     ctx.Step(`^we POST the HandlerRequest to \/hello$`, wePOSTTheHandlerRequestToHello)
// }
//
// $ godog
// Use of godog without a sub-command is deprecated. Please use godog build, godog run or godog version.
// Feature: Bad Method
//
//   Scenario: Good request                           # features/handler.feature:2
//     Given we create a HandlerRequest payload with: # handler_test.go:15 -> weCreateAHandlerRequestPayloadWith
//       | reader |
//       | coder  |
//       | other  |
//     And we POST the HandlerRequest to /hello       # handler_test.go:25 -> wePOSTTheHandlerRequestToHello
//     Then the response code should be 200           # handler_test.go:40 -> theResponseCodeShouldBe
//     And the response body should be:               # handler_test.go:49 -> theResponseBodyShouldBe
//       | BDD testing reader |
//       | BDD testing coder  |
//       | BDD testing other  |
//
// 1 scenarios (1 passed)
// 4 steps (4 passed)
// 492.083µs
//
// Cucumber 프레임워크는 페어 프로그래밍(pair programming), 엔드 투 엔드 테스트, 작성된 지침을 통한 의사소통, 비기술 종사자들이 이해할 수 있도록
// 의사소통하는데 가장 적합한 모든 종류의 테스트에 훌룡하게 동작한다. 한 stepㅇ ㅣ구현되면, 필요한 곳에서 이 Step을 재사용할 수 있다. 서비스 간 통합을
// 테스트할 때는 먼저 여러분의 환경이 HTTP 연결을 받도록 설정해 확인한다. 그러면 실제 HTTP 클라이언트를 사용하도록 테스트를 작성할 수 있다.
//
// 동작 기반 개발(BDD, Behavior-Driven Development)의 datadog 구현은 다른 Cucumber 프레임워크를 사용할 때 기대할 수 있는 몇 가지 기능이
// 부족하다. 여기에는 예제의 부족, 함수간 컨텍스트 전달, 다른 키워드 수 등이 포함된다. 하지만 이는 좋은 출발점이며, 상태 추적을 위한 전역 변수(시나리오
// 사이에 이런 전역 변수의 정리를 보장해야 한다.)와 같이 이 예제의 몇 가지 기법을 사용하면 상당히 강력한 테스트를 제작할 수 있다. datadog 테스트 패키지는
// 서드파티 테스트 실행기를 사용하므로 gocov나 go test -cover와 같은 패키지와 함께 사용할 수 없다.
func main() {

}
