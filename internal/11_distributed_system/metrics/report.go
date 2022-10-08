package metrics

import (
	"github.com/rcrowley/go-metrics"
	"net/http"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	t := metrics.GetOrRegisterTimer("reporthandler.writemetrics", nil)
	t.Time(func() {
		metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
	})
}
