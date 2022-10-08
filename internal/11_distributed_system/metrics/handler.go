package metrics

import (
	"github.com/rcrowley/go-metrics"
	"net/http"
	"time"
)

// CounterHandler 함수는 호출될 때마다 카운터(counter)를 갱신한다.
func CounterHanlder(w http.ResponseWriter, r *http.Request) {
	c := metrics.GetOrRegisterCounter("counterhandler.counter", nil)
	c.Inc(1)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

// TimerHandler 함수는 완료하는 데까지 필요한 시간(기간)을 기록한다.
func TimerHandler(w http.ResponseWriter, r *http.Request) {
	currt := time.Now()
	t := metrics.GetOrRegisterTimer("timerhandler.timer", nil)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
	t.UpdateSince(currt)
}
