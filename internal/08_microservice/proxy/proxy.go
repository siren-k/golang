package proxy

import (
	"log"
	"net/http"
)

// Proxy는 프록시를 설정하기 위해 구성된 클라이언트와 BaseURL을 저장한다.
type Proxy struct {
	Client  *http.Client
	BaseURL string
}

// ServeHTTP 함수는 프록시가 핸들러(Handler) 인터페이스를 구현한다는 것을 의미한다.
// 요청을 조직해 BaseURL올 전달한 다음 응답을 반환한다.
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error occurred during process request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error occurred during client operation: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	CopyResponse(w, resp)
}
