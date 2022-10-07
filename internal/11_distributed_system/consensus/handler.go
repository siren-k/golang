package consensus

import (
	"net/http"
	"time"
)

// Handler 함수는 get 요청의 매개변수 ?next=의 값을 가져와 이 값에 포함된 상태로 전환을 시도한다.
func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.FormValue("next")
	for address, raft := range rafts {
		if address != raft.Leader() {
			continue
		}

		result := raft.Apply([]byte(state), 1*time.Second)
		if result.Error() != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newState, ok := result.Response().(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newState != state {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid transition"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(newState))
		return
	}
}
