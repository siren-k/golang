package cards

import (
	"github.com/graphql-go/graphql"
	"strings"
)

// Resolve 함수는 패(suit)와 값(value)으로 카드 필터링을 처리한다.
func Resolve(p graphql.ResolveParams) (interface{}, error) {
	finalCards := []Card{}
	suit, suitOK := p.Args["suit"].(string)
	suit = strings.ToLower(suit)

	value, valueOK := p.Args["value"].(string)
	value = strings.ToLower(value)

	for _, card := range cards {
		if suitOK && suit != strings.ToLower(card.Suit) {
			continue
		}
		if valueOK && value != strings.ToLower(card.Value) {
			continue
		}

		finalCards = append(finalCards, card)
	}
	return finalCards, nil
}
