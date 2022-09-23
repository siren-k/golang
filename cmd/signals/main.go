package main

import (
	"fmt"
	signals2 "golang/signals"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 채널을 초기화한다.
	signals := make(chan os.Signal)
	done := make(chan bool)

	// signals 라이브러리에 연결한다.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 이 Go 루틴에 의해 신호가 잡혔으면
	// done으로 쓴다.
	go signals2.CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}
