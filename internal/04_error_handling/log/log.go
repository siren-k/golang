package log

import (
	"bytes"
	"fmt"
	"log"
)

func Log() {
	// bytes.Buffer에 기록하기 위해 로거를 설정한다.
	buf := bytes.Buffer{}

	// 두 번째 인자는 접두사이며, 마지막 인자는 논리 연산자 or를 결합해 설정하는 옵션이다.
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Function to log and crash", true)

	fmt.Println(buf.String())
}
