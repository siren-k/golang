package negotiate

import (
	"encoding/xml"
	"net/http"
)

// Payload는 xml과 json으로 데이터의 레이아웃을 정의한다.
type Payload struct {
	XMLName xml.Name `xml:"payload" json:"-"`
	Status  string   `xml:"status" json:"status"`
}

// Handler 함수는 요청을 사용해 negotiator 객체를 가져온 다음
// 데이터를 렌더링(표시)한다.
func Handler(w http.ResponseWriter, r *http.Request) {
	n := GetNegotiator(r)
	n.Respond(w, http.StatusOK, &Payload{Status: "Successful"})
}
