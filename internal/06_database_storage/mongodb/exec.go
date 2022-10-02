package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

// State는 예제에서 사용할 데이터 모델이다.
type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"population"`
}

// Exec 함수는 Example을 생성하고 쿼리한다.
func Exec(address string) error {
	ctx := context.Background()
	db, err := Setup(ctx, address)
	if err != nil {
		return err
	}

	coll := db.Database("gocookbook").Collection("example")
	vals := []interface{}{&State{"Washington", 7062000}, &State{"Oregon", 3970000}}

	// 한 번에 여러 열을 삽입할 수 있다.
	if _, err = coll.InsertMany(ctx, vals); err != nil {
		return err
	}

	var s State
	if err = coll.FindOne(ctx, bson.M{"name": "Washington"}).Decode(&s); err != nil {
		return err
	}

	if err = coll.Drop(ctx); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)
	return nil
}
