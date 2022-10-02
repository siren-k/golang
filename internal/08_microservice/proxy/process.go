package proxy

import (
	"bytes"
	"net/http"
	"net/url"
)

// ProcessRequest 함수는 프록시 설정에 따라 요청을 수정한다.
func (p *Proxy) ProcessRequest(r *http.Request) error {
	proxyURLRaw := p.BaseURL + r.URL.String()

	proxyURL, err := url.Parse(proxyURLRaw)
	if err != nil {
		return err
	}
	r.URL = proxyURL
	r.Host = proxyURL.Host
	r.RequestURI = ""
	return nil
}

// CopyResponse 함수는 클라이언트 응답을 받아 모든 내용을 ResponseWriter에 쓴다.
func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	var out bytes.Buffer
	out.ReadFrom(resp.Body)

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(out.Bytes())
}
