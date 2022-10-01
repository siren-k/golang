package pools

import (
	"context"
	"time"
)

// ExecWithTimeout 함수는 현재 시간을 얻어오려다 시간 초과를 발생시킨다.
func ExecWithTimeout() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx := context.Background()
	// 즉시 시간 초과를 발생시킨다.
	ctx, cancel := context.WithDeadline(ctx, time.Now())
	// 완료한 다음, cancel을 호출한다.
	defer cancel()

	// 트랜잭션이 컨텍스트를 인지한다.
	_, err = db.BeginTx(ctx, nil)

	return err
}
