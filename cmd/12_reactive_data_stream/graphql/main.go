package main

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"golang/internal/12_reactive_data_stream/graphql/cards"
	"log"
)

// cards.go 파일은 카드 객체를 정의하고 cards라는 이름의 전역 변수에 기본 카드 패를 초기화한다. 데이터베이스와 같은 장기 저장소에도
// 이 상태를 보관할 수 있다. 그런 다음, types.go 파일에 CardType을 정의해 graphql이 카드 객체로 응답할 수 있도록 설정한다.
// 이어서 resolve.go 파일로 이동해 값과 타입으로 카드를 필터링하는 방법을 정의한다. schema.go 파이에 정의한 최종 스키마에서
// 이 Resolve 함수를 사용한다.
//
// 필요에 따라서는 데이터베이스를 검색하기 위해 이 에제의 Resolve 함수를 수정할 수도 있다. 마지막으로 스키마를 불러와 이 스키마에 대해
// 쿼리를 실행한다. 적은 양의 수정을 통해 스키마를 REST 엔드포인트에 올렸지만, 코드의 간결성을 위해 하드코딩된 쿼리를 실행한다.
// http://graphql.org/learn/queries/에서 GraphQL 쿼리에 대한 더 자세한 내용을 참고할 수 있다.
func main() {
	// 스키마를 가져온다.
	schema, err := cards.Setup()
	if err != nil {
		panic(err)
	}

	// 요청한다.
	query := `
	{
		cards(value: "A") {
			value
			suit
		}
	}`
	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		panic(err)
	}

	log.Printf("%s \n", rJSON)
}
