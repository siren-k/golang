package decorator

import "net/http"

// TransportFunc 함수는 RoundTripper 인터페이스를 구현한다.
type TransportFunc func(*http.Request) (*http.Response, error)

// RoundTrip 함수는 원래의 함소를 호출만 한다.
func (tf TransportFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return tf(r)
}

// Decorator 함수는 미들웨어 내부 함수를 표현하는 편의 함수다.
type Decorator func(http.RoundTripper) http.RoundTripper

// Decorate 함수는 모든 미들웨어를 래핑하는 도움 함수다.
func Decorate(t http.RoundTripper, rts ...Decorator) http.RoundTripper {
	decorated := t
	for _, rt := range rts {
		decorated = rt(decorated)
	}
	return decorated
}
