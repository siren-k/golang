package signals

import (
	"fmt"
	"os"
	"syscall"
)

// CatchSig 함수는 SIGINT 인터럽트에 대한
// 리스너를 설정한다.
func CatchSig(ch chan os.Signal, done chan bool) {
	// 신호에 대기하기 위해 블록(block)한다.
	sig := <-ch
	// 신호를 받으면 출력한다.
	fmt.Println("\nsig received:", sig)

	// 모든 신호에 대한 핸들러(handlers)를 설정할 수 있다.
	switch sig {
	case syscall.SIGINT:
		println("handling a SIGINT now!")
	case syscall.SIGTERM:
		println("handling a SIGTERM in an entirely different way!")
	default:
		println("unexpected signal received")
	}

	// 종료한다.
	done <- true
}
