package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Exec 함수는 redis 작업을 처리한다.
func Exec() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	c1 := "value"
	// value는 인터페이스이며, 마지막 매개변수가 무엇이든 간에
	// redis의 만료 시간을 지정할 수 있다.
	conn.Set("key", c1, 5*time.Second)

	var result string
	if err = conn.Get("key").Scan(&result); err != nil {
		switch err {
		// 키가 발견되지 않았다는 것을 의미한다.
		case redis.Nil:
			return nil
		default:
			return err
		}
	}

	fmt.Println("result =", result)

	return nil
}
