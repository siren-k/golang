package channels

import "time"

// Sender 함수는 ch에 "done"이라는 문자열 값이 작성될 떄까지 '틱(tick)' 단위로 계속 확인하다가
// done 값이 설정되면 "sender done." 메시지를 전송하고 종료한다.
func Sender(ch chan string, done chan bool) {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-done:
			ch <- "sender done."
			return
		case <-t:
			ch <- "tick"
		}
	}
}
