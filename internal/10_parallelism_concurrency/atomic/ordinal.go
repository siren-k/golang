package atomic

import (
	"sync"
	"sync/atomic"
)

// Ordinal 한 번만 초기화할 수 있는 전역 값을 갖는다.
type Ordinal struct {
	ordinal uint64
	once    *sync.Once
}

// NewOrdinal 함수는 설정된 ordinal을 반환한다.
func NewOrdinal() *Ordinal {
	return &Ordinal{once: &sync.Once{}}
}

// Init 함수는 ordinal 값을 설명하며 한 번만 실행될 수 있다.
func (o *Ordinal) Init(val uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, val)
	})
}

// GetOrdinal 함수는 현재의 ordinal 값을 반환한다.
func (o *Ordinal) GetOrdinal() uint64 {
	return atomic.LoadUint64(&o.ordinal)
}

// Increment 함수는 현재의 ordinal 값을 증가시킨다.
func (o *Ordinal) Increment() {
	atomic.AddUint64(&o.ordinal, 1)
}
