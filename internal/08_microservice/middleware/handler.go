package middleware

import "net/http"

// 매우 기초적인 핸들러
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
