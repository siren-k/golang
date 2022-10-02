package main

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	oauthstore2 "golang/internal/07_web_client_and_api/oauthstore"
	"io"
	"os"
)

// 이 예제에서는 토큰의 내용을 파일에 저장하고 조회한다. 예제를 처음 실행하면 전체 코드 교환을 실행한다.
// 하지만 이후의 실행해서는 접근 토큰을 재사용하고 가능하다면 토큰을 갱신한다.
//
// config.go 파일은 표준 OAuth2 config를 래핑한다. 토큰 조회와 관련된 모든 메소드(함수)에 대해 먼저
// 로컬 저장소에 유효한 토큰이 있는지 확인한다. 유효한 토큰이 없다면 표준 config를 사용해 토큰을 조회한 다음
// 파일에 저장한다.
//
// tokensource.go 파일은 Config와 함꼐 사용자 정의 TokenSource 인터페이스를 구현한다. Config와
// 마찬가지로 먼저 파일에서 토큰 조회를 시도하고 실패하면 새로운 토큰을 설정한다.
//
// storage.go 파일은 Config와 TokenSource에서 사용하는 저장소 인터페이스다. 두 개의 메소드(함수)만
// 정의하면 이전 예제에서 했던 것과 비슷한 OAuth2 기반 동작을 수행하는 도움 함수도 포함한다. 하지만 유효한
// 토큰을 가진 파일이 존재하면 파일을 대신 사용한다.
//
// filestorage.go 파일은 저장소 인터페이스를 구현한다. 새 토큰을 저장할 때 먼저 파일에서 필요한 내용을 자르고
// 토큰 구조의 JSON 표현을 파일에 쓴다. 토큰이 있는 경우에는 파일을 해석해 토큰을 반환한다.
func main() {
	conf := oauthstore2.Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT"),
			ClientSecret: os.Getenv("GITHUB_SECRET"),
			Scopes:       []string{"repo", "user"},
			Endpoint:     github.Endpoint,
		},
		Storage: &oauthstore2.FileStorage{Path: "token.txt"},
	}

	ctx := context.Background()
	token, err := oauthstore2.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}

	client := conf.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
