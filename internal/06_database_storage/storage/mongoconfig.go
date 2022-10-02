package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MongoStorage는 에제의 저장소 인터페이스를 구현한다.
type MongoStorage struct {
	*mongo.Client
	DB         string
	Collection string
}

// NewMongoStorage 함수는 MongoStorage를 초기화한다.
func NewMongoStorage(ctx context.Context, connection, db, collection string) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	ms := MongoStorage{
		Client:     client,
		DB:         db,
		Collection: collection,
	}

	return &ms, nil
}
