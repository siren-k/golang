package context

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Exec 함수는 두 개의 랜덤 타이머를 설정하고 먼저 발생하는 타이머의 컨텍스트 값을 출력한다.
func Exec() {
	// 기본 컨텍스트
	ctx := context.Background()
	ctx = Setup(ctx)

	rand.Seed(time.Now().UnixNano())

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(rand.Intn(2))*time.Millisecond)
	defer cancel()

	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Duration(rand.Intn(2))*time.Millisecond))
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			fmt.Println(GetValue(ctx, timeoutKey))
			return
		case <-deadlineCtx.Done():
			fmt.Println(GetValue(ctx, deadlineKey))
			return
		}
	}
}
