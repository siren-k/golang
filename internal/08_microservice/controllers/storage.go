package controllers

// Storage 인터페이스는 단일 값의 Get과 Put을 지원한다.
type Storage interface {
	Get() string
	Put(string)
}

// MemStorage는 Storage를 구현한다.
type MemStorage struct {
	value string
}

// 인메모리 값을 가져온다.
func (m *MemStorage) Get() string {
	return m.value
}

// 인메모리 값을 설정한다.
func (m *MemStorage) Put(s string) {
	m.value = s
}
