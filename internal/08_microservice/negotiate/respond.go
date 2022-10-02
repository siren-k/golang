package negotiate

import (
	"github.com/unrolled/render"
	"io"
)

// Respond 함수는 내용 유형(Content Type)에 따라 적절하게 응답한다.
func (n *Negotiator) Respond(w io.Writer, status int, v interface{}) {
	switch n.ContentType {
	case render.ContentJSON:
		n.Render.JSON(w, status, v)
	case render.ContentXML:
		n.Render.XML(w, status, v)
	default:
		n.Render.JSON(w, status, v)
	}
}
