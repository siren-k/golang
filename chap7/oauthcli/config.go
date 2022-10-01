package oauthcli

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

// Setup 함수는 github와 소통하기 위한 구성이 완료된 oauth2Config를 반환한다.
// id와 secret 값이 설정된 환경 변수가 필요하다.
func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}

// GetToken 함수는 github oauth2 토큰을 가져온다.
func GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your brower and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}
