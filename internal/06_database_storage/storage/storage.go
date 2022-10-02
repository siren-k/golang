package storage

import "context"

// Item은 상점의 항목을 표현한다
type Item struct {
	Name  string
	Price int64
}

// Storage는 저장소 인터페이스이며
// Mongo 저장소를 활용해 구현할 예정이다
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}
