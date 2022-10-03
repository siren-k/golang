package main

import (
	"fmt"
	"golang/internal/10_parallelism_concurrency/atomic"
	"sync"
)

// 맵(map) 예제에서는 ReadWrite 뮤텍스를 사용했다. 값을 읽으려는 이 뮤텍스의 기본 개념은 여러 Reader가 읽기(read) 락을
// 획득할 수 있지만 Writer는 하나만 쓰기(write) 락을 획득할 수 있다는 것이다. 또한 누군가(reader 또는 writer) 쓰기 락을
// 소유하면 다른 Writer는 쓰기 락을 획득할 수 없다. 이는 표준 뮤텍스와 비교했을 때 읽기 작업이 매우 빠르며 논블로킹(non-blocking)
// 방식이므로 유용하다. 데이터를 설정하려고 할 때마다 Lock() 객체를 사용하고 데이터를 읽으려고 할 때마나 RLock(0을 사용한다. 따라서
// Unlock과 RUnlock()을 사요해 애플리케이션이 교착 상태(deadlock)에 빠지지 않도록 하는 것이 중요하다. Unlock() 객체를 지연
// 호출하는 것이 유용하지만 Unlock()을 직접 호출하는 것보다 느릴 수 있다.
//
// 락이 설정된 값으로 추가 작업을 그룹화하려는 경우에는 이 패턴이 유연하지 않을 수 있다. 예를 들어, 경우에 따라 락을 설정하고 추가적인
// 작업을 처리하면 이 작업을 완료한 후에만 락을 해제할 수 있다. 따라서 설계 시 이를 고려하는 것이 중요하다.
//
// Ordinal 값을 읽고 설정할 때 sync/atomic 패키지를 사용했다. 또한 atomic.CompareAndSwapUInt64()와 같은 원자적 비교
// 연산도 매우 유용하다. 이 예제를 사용하면 Init 함수를 Ordinal 객체에 대해 한 번만 호출할 수 있기 때문에 Ordinal 값을
// 증가시킬 수 있고 원자적으로 처리할 수 있다.
//
// 이 에제는 루프로 열 개의 고루틴(sync.Waitgroup으로 동기화)을 생성하고, Ordinal 값이 정확하게 열 번 증가하고 맵의 모든 키가
// 적절하게 설정된 것을 보여줬다.
func main() {
	o := atomic.NewOrdinal()
	m := atomic.NewSafeMap()
	o.Init(1123)
	fmt.Println("Initial ordinal is:", o.GetOrdinal())

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(fmt.Sprint(i), "success")
			o.Increment()
		}(i)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		v, err := m.Get(fmt.Sprint(1))
		if err != nil || v != "success" {
			panic(err)
		}
	}

	fmt.Println("final ordinal is:", o.GetOrdinal())
	fmt.Println("all keys found an marked as: 'success'")
}
