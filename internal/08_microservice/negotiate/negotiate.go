package negotiate

import (
	"github.com/unrolled/render"
	"net/http"
)

// Negotiator는 render를 래핑해 ContentType으로 내용을 전환한다.
type Negotiator struct {
	ContentType string
	*render.Render
}

// GetNegotiator 함수는 요청을 받고 Content-Type 헤더로부터 ContentType을 판별한다.
func GetNegotiator(r *http.Request) *Negotiator {
	contentType := r.Header.Get("Content-Type")
	return &Negotiator{
		ContentType: contentType,
		Render:      render.New(),
	}
}
