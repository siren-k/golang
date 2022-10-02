package storage

import (
	"context"
	"fmt"
)

// Exec 함수는 저장소를 초기화한 다음, 저장소 인터페이스를 사용해
// 작업을 처리한다.
func Exec() error {
	ctx := context.Background()
	m, err := NewMongoStorage(ctx, "localhost", "gocookbook", "items")
	if err != nil {
		return err
	}

	if err = PerformOperations(m); err != nil {
		return err
	}

	if err = m.Client.Database(m.DB).Collection(m.Collection).Drop(ctx); err != nil {
		return err
	}

	return nil
}

// PerformOperations 함수는 양초(candle) 항목을 생성하고
// 데이터베이스에서 이 값을 조회한다.
func PerformOperations(s Storage) error {
	ctx := context.Background()
	i := Item{Name: "candles", Price: 100}
	if err := s.Put(ctx, &i); err != nil {
		return err
	}

	candles, err := s.GetByName(ctx, "candles")
	if err != nil {
		return err
	}

	fmt.Printf("Result: %#v\n", candles)
	return nil
}
