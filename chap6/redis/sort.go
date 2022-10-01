package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

// Sort 함수는 redis 작업을 정렬한다.
func Sort() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	listkey := "list"
	if err = conn.LPush(listkey, 1).Err(); err != nil {
		return err
	}
	// 다음 명령 중 하나라도 오류가 발생하면 리스트 키를 정리(제거)한다.
	defer conn.Del(listkey)

	if err = conn.LPush(listkey, 3).Err(); err != nil {
		return err
	}

	if err = conn.LPush(listkey, 2).Err(); err != nil {
		return err
	}

	res, err := conn.Sort(listkey, &redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
