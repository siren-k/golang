package oauthstore

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
)

// Storage는 제네릭 저장소 인터페이스다.
type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(token *oauth2.Token) error
}

// GetToken 함수는 github oauth2 토큰을 조회한다.
func GetToken(ctx context.Context, conf Config) (*oauth2.Token, error) {
	token, err := conf.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, err
	}

	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your brower and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	var code string
	if _, err = fmt.Scan(&code); err != nil {
		return nil, err
	}

	return conf.Exchange(ctx, code)
}
