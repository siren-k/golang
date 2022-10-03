package atomic

import (
	"github.com/pkg/errors"
	"sync"
)

// SafeMap 스레드 안정성을 유지하면서 값을 읽고 설정할 수 있도록 뮤텍스를 사용한다.
type SafeMap struct {
	m  map[string]string
	mu *sync.RWMutex
}

// NewSafeMap 함수는 SafeMap을 생성한다.
func NewSafeMap() SafeMap {
	return SafeMap{
		m:  make(map[string]string),
		mu: &sync.RWMutex{},
	}
}

// Set 함수는 lock을 작성하고 전달될 키로 값을 설정한다.
func (t *SafeMap) Set(key, value string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.m[key] = value
}

// Get 함수는 RW lock을 사용해 존재하는 경우 값을 읽고 존재하지 않는 경우 오류를 반환한다.
func (t *SafeMap) Get(key string) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if v, ok := t.m[key]; ok {
		return v, nil
	}

	return "", errors.New("Key not found")
}
