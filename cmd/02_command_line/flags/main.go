package main

import (
	"flag"
	"fmt"
	"golang/internal/01_io_filesystem/flags"
)

func main() {
	// 설정을 초기화한다.
	c := flags.Config{}
	c.Setup()

	// 다음 코드는 일반적으로 main에서 호출한다.
	flag.Parse()

	fmt.Println(c.GetMessage())
}
