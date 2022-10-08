package kafkaflow

import (
	"fmt"
	flow "github.com/trustmaster/goflow"
	"strings"
)

// Upper는 대문자로 들어오는 문자열을 저장한다.
type Upper struct {
	Val <-chan string
	Res chan<- string
}

// Process 함수는 입력 값을 루프로 처리하면서 입력 값의 대문자 문자열 버전을 Res에 쓴다.
func (e *Upper) Process() {
	for val := range e.Val {
		e.Res <- strings.ToUpper(val)
	}
}

// Printer는 표준 출력(stdout)에 출력하기 위한 컴포넌트다.
type Printer struct {
	flow.Component
	Line <-chan string
}

// Process 함수는 현재 받은 라인(line)을 출력한다.
func (p *Printer) Process() {
	for line := range p.Line {
		fmt.Println(line)
	}
}
