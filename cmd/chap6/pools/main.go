package main

import "golang/chap6/pools"

// 연결 풀의 깊이를 제어할 수 있는 기능은 매우 유용하다. 이를 통해 데이터베이스에 과부하가 걸리는 것을 방지할 수 있다.
// 하지만 이런 점이 시간 초과 상황에서 어떤 의미를 갖는지 고려하는 것 또한 중요하다. 이 예제에서 했던 것과 같이,
// 설정된 수의 연결과 엄격한 컨텍스트 기반 타임아웃을 모두 적용하면 너무 많은 연결을 시도해 과부하가 걸린 애플리케이션에
// 대한 요청이 시간 초과에 자주 걸리는 경우가 발생한다.
//
// 연결을 사용할 수 있을 때까지 대기하다가 시간 초과가 발생하기 때문이다. database/sql 패키지에 새로 추가된 컨텍스트
// 기능을 사용하면, 쿼리 수행과 관련된 단계를 포함해 전체 요청에 대한 공유 타임아웃을 사용할 수 있어 처리가 훨씬 더
// 간단해진다.
func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}