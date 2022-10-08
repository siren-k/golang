package cards

import "github.com/graphql-go/graphql"

// Setup 함수는 카드 스키마를 준비하고 반환한다.
func Setup() (graphql.Schema, error) {
	cardType := CardType()
	// 스키마
	fields := graphql.Fields{
		"cards": &graphql.Field{
			Type: graphql.NewList(cardType),
			Args: graphql.FieldConfigArgument{
				"suit": &graphql.ArgumentConfig{
					Description: "Filter cards by card suit(hearts, clubs, diamonds, spades)",
					Type:        graphql.String,
				},
				"value": &graphql.ArgumentConfig{
					Description: "Filter cards by card value(A-K)",
					Type:        graphql.String,
				},
			},
			Resolve: Resolve,
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	schema, err := graphql.NewSchema(schemaConfig)

	return schema, err
}
