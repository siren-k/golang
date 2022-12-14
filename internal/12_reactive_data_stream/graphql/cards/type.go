package cards

import "github.com/graphql-go/graphql"

// CardType 함수는 card graphql 객체를 반환한다.
func CardType() *graphql.Object {
	cardType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Card",
		Description: "A Playing Card",
		Fields: graphql.Fields{
			"value": &graphql.Field{
				Type:        graphql.String,
				Description: "Ace through King",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if card, ok := p.Source.(Card); ok {
						return card.Value, nil
					}
					return nil, nil
				},
			},
			"suit": &graphql.Field{
				Type:        graphql.String,
				Description: "Hearts, Diamonds, Clubs, Spades",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if card, ok := p.Source.(Card); ok {
						return card.Suit, nil
					}
					return nil, nil
				},
			},
		},
	})
	return cardType
}
