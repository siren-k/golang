package bdd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerRequest는 Handler 함수가 json으로 디코딩한다.
type HandlerRequest struct {
	Name string `json:"name"`
}

// Handler 함수는 요청을 매개변수로 받아 응답을 렌더링한다.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(r.Body)
	var req HandlerRequest
	if err := dec.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("BDD testing %s", req.Name)))
}
