package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware는 모든 미들웨어 함수가 리턴 타입이다.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyMiddleware 함수는 모든 미들웨어를 적용하며
// 마지막 매개변수는 컨텍스트 전달을 위해 외부로 래핑한다.
func ApplyMiddleware(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	applied := h
	for _, m := range middleware {
		applied = m(applied)
	}
	return applied
}

// Logger 함수는 요청을 로그로 기록하며 SetID()를 통해 전달된 id를 사용한다.
func Logger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s with id %s", r.URL, GetID(r.Context()))
			next(w, r)
			l.Printf("completed request to %s with id %s in %s", r.URL, GetID(r.Context()), time.Since(start))
		}
	}
}
