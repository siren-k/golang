package kafkaflow

import "github.com/trustmaster/goflow"

// NewUpperApp 함수는 컴포넌트들을 서로 연결한다.
func NewUpperApp() *goflow.Graph {
	u := goflow.NewGraph()

	u.Add("upper", new(Upper))
	u.Add("printer", new(Printer))

	u.Connect("upper", "Res", "printer", "Line")
	u.MapInPort("In", "upper", "Val")

	return u
}
