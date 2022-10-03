package channels

import (
	"context"
	"fmt"
	"time"
)

// Printer 함수는 ch chan에 있는 모든 값을 출력하며 200밀리초마다 tock을 출력한다.
// 이 동작은 타임아웃이나 취소 등으로 컨텍스트(ctx 매개변수)가 완료될 떄까지 반복된다.
func Printer(ctx context.Context, ch chan string) {
	t := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printer done.")
			return
		case res := <-ch:
			fmt.Println(res)
		case <-t:
			fmt.Println("tock")
		}
	}
}
