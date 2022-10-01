package main

import "golang/chap6/redis"

// Go에서 Redis를 활용하는 작업은 MySQL을 활용한 작업과 비슷하다. 표준 라이브러리는 없지만,
// 많은 라이브러리가 Redis에서 Go 타입으로 데이터를 읽는 Scan()과 같은 함수와 동일한 규칙을 따른다.
// 이와 같은 경우에는 가장 적합한 라이브러리를 선택하는 것이 어려울 수 있다. 이런 내용은 빠르게 바뀔 수
// 있기 때문에 주기적으로 어떤 라이브러리를 사용할 수 있는지 조사하는 것이 좋다.
//
// 이 예제에서는 redis 패키지를 사용해 데이터 설정, 읽기, 좀 더 복잡한 정렬 함수와 기본 구성을
// 살펴봤다. database/sql 패키지와 같이 타임아웃, 풀(pool) 크기 등의 형태로 추가 구성을 설정할
// 수 있다. 또한, Redis 패키지는 Redis 클러스터 지원, Zscore 및 counter 객체,
// 분삭 락(distributed lock)을 포함해 사용하는 것이 좋다.
//
// 이전 예제와 마찬가지로, 설정과 보안 작업을 쉽계 처리하기 위해 Redis 설정 및 세부 구성 정보를
// 저장하는 설정 객체(config)를 사용하는 것이 좋다.
func main() {
	if err := redis.Exec(); err != nil {
		panic(err)
	}

	if err := redis.Sort(); err != nil {
		panic(err)
	}
}
