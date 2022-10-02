package oauthstore

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

// Config는 기본 oauth2.Config를 래핑해 예제의 저장소를 추가한다.
type Config struct {
	*oauth2.Config
	Storage
}

// Exchange 함수는 토큰을 조회한 다음 이를 저장한다.
func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	if err = c.Storage.SetToken(token); err != nil {
		return nil, err
	}

	return token, nil
}

// TokenSource 함수는 저장된 토큰을 전달받을 수 있으며
// 새로운 토큰이 조회되면 이 토큰을 저장한다.
func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return StorageTokenSource(ctx, c, t)
}

// Client 함수는 TokenSource에 추가된다.
func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}
