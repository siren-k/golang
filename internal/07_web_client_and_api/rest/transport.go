package rest

import "net/http"

// APITransport는 모든 요청에 대해 SetBasicAuth를 수행한다.
type APITransport struct {
	*http.Transport
	username, password string
}

// RoundTrip 함수는 기본 전송을 수행하기 전에 기본적인 인증을 진행한다.
func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.Transport.RoundTrip(req)
}
