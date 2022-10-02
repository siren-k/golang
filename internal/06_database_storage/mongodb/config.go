package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Setup mongo 클라이언트를 초기화한다.
func Setup(ctx context.Context, address string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// cancel은 설정이 종료될 때 호출된다.(지연 호출)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
