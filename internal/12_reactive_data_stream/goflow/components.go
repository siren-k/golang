package goflow

import (
	"encoding/base64"
	"fmt"
)

// Encoder는 모든 입력을 base64로 인코딩한다.
type Encoder struct {
	Val <-chan string
	Res chan<- string
}

// Process 함수는 인코딩을 한 다음 Res로 결과를 보낸다.
func (e *Encoder) Process() {
	for val := range e.Val {
		encoded := base64.StdEncoding.EncodeToString([]byte(val))
		e.Res <- fmt.Sprintf("%s => %s", val, encoded)
	}
}

// Printer는 표준 출력(stdout)으로 출력하기 위한 컴포넌트다.
type Printer struct {
	Line <-chan string
}

// Process 함수는 전달받은 현재 라인(line)을 출력한다.
func (p *Printer) Process() {
	for line := range p.Line {
		fmt.Println(line)
	}
}
