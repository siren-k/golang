package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler 함수는 GET 요청으로 매개변수 "name"을 받아
// Hello <name>!을 일반 텍스트로 응답한다.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
}
