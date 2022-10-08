package goflow

import "github.com/trustmaster/goflow"

// NewEncodingApp 함수는 컴포넌트들을 연결한다.
func NewEncodingApp() *goflow.Graph {
	e := goflow.NewGraph()

	// 컴포넌트 타입을 정의한다.
	e.Add("encoder", new(Encoder))
	e.Add("printer", new(Printer))

	// 채널을 사용해 컴포넌트들을 연결한다.
	e.Connect("encoder", "Res", "printer", "Line")

	// in 채널을 OnVal 함수에 연결된 Val에 대응시킨다.
	e.MapInPort("In", "encoder", "Val")

	return e
}
