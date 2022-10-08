package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

// State는 데이터 모델이다.
type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"pop"`
}

// Exec 함수는 예제를 생성하고 이를 쿼리한다.
func Exec(address string) error {
	ctx := context.Background()
	db, err := Setup(ctx, address)
	if err != nil {
		return err
	}

	conn := db.Database("gocookbook").Collection("example")

	vals := []interface{}{&State{"Washington", 7062000}, &State{"Oregon", 3970000}}

	// 여러 열(row)을 한 번에 삽입한다.
	if _, err := conn.InsertMany(ctx, vals); err != nil {
		return err
	}

	var s State
	if err := conn.FindOne(ctx, bson.M{"name": "Washington"}).Decode(&s); err != nil {
		return err
	}

	if err := conn.Drop(ctx); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)
	return nil
}
