package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// GetByName 함수는 정확한 이름을 사용해 mongodb에서 항목(아이템)을 조회한다.
func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	c := m.Client.Database(m.DB).Collection(m.Collection)
	var i Item
	if err := c.FindOne(ctx, bson.M{"name": name}).Decode(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

// Put 함수는 예제의 mongo 인터페이스에 항목을 추가한다.
func (m *MongoStorage) Put(ctx context.Context, i *Item) error {
	c := m.Client.Database(m.DB).Collection(m.Collection)
	_, err := c.InsertOne(ctx, i)
	return err
}
