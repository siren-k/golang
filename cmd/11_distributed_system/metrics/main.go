package main

import (
	"fmt"
	"golang/internal/11_distributed_system/metrics"
	"net/http"
)

// $ curl localhost:8080/counter
// success
//
// $ curl localhost:8080/timer
// success
//
// $ curl localhost:8080/report
// {"counterhandler.counter":{"count":1},"reporthandler.writemetrics":{"15m.rate":0,"1m.rate":0,"5m.rate":0,"75%":0,"95%":0,"99%":0,"99.9%":0,"count":0,"max":0,"mean":0,"mean.rate":0,"median":0,"min":0,"stddev":0},"timerhandler.timer":{"15m.rate":0.19779007785878447,"1m.rate":0.16929634497812282,"5m.rate":0.1934432200964012,"75%":154417,"95%":154417,"99%":154417,"99.9%":154417,"count":1,"max":154417,"mean":154417,"mean.rate":0.15205803420049435,"median":154417,"min":154417,"stddev":0}}

// gometrics는 수집 정보를 모두 레지스트리에 보관한다. gometrics가 설정되면 카운터나 타이머와 같은 수집 정보 발행 옵션을 사용할 수 있으며, 이에 대한
// 업데이트 정보를 레지스트리에 저장한다. 이런 수집 정보들은 서드파티 도구에 내보내는 여러 도구가 있다. 예제에서는 모든 수집 정보를 JSON 포맷으로 생성하는
// 핸들러를 설정했다.
//
// 카운터를 증가시키는 핸들러, 핸들러 종료 시간을 기록하는 핸들러, 보고서를 출력하는 핸들러(또한, 추가 카운터를 증가시킴) 이렇게 세 개의 핸들러를 설정했다.
// GetOrRegister 함수는 스레드 안전한 방식으로 동작하며, 원자적(atomically)으로 수집 정보를 가져오거나 현재 존재하지 않는 경우에는 정보 발행기를
// 생성하는데 유용하다. 또한 사전에 한 번 모든 정보를 등록할 수도 있다.
func main() {
	// 정보 생성을 위한 핸들러
	http.HandleFunc("/counter", metrics.CounterHanlder)
	http.HandleFunc("/timer", metrics.TimerHandler)
	http.HandleFunc("/report", metrics.ReportHandler)

	fmt.Println("listening on :8080")
	panic(http.ListenAndServe(":8080", nil))
}
