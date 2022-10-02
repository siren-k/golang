package main

import (
	"context"
	oauthcli2 "golang/internal/07_web_client_and_api/oauthcli"
)

// 표준 OAuth2 동작은 리다이렉트(redirect) 기반이며, 서버는 지정한 엔드포인트로 경로를 재지정한다. 그런 다음
// 서버는 코드를 가져와 토큰을 교환해야 한다. 이 예제는 https://localhost나 https://a-domain-you-own과 같은
// URL을 사용하고 코드를 수동으로 복사/붙여넣기를 한 다음 엔터 키를 누르는 것을 허용하는 것으로 이 요구 사항을 우회한다.
// 토큰 교환이 완료되면 클라이언트는 필요에 따라 지능적으로 토큰을 갱신한다.
//
// 어떤 식으로든 토큰을 저장하지 않는다는 점에 주의해야 한다. 프로그램에 문제가 생기면 토큰을 다시 교환해야 한다. 또한
// 토큰이 만료되거나 토큰을 잃어버렸거나 토큰이 손상되지 않는 한 토큰을 명시적으로 한 번만 조회한다는 점에 주의한다.
// 클라이언트의 설정을 완료하면, OAuth2 작업을 진핸하는 동안 적절한 범위(scope)가 요청되는 한 API에 대한 일반적인
// HTTP 동작은 모두 수행할 수 있다. 이 에제는 "repo"와 "user" 범위를 요청하지만, 필요하다면 더 많은(또는 더 적은)
// 범위도 요청할 수 있다.
func main() {
	ctx := context.Background()
	conf := oauthcli2.Setup()

	token, err := oauthcli2.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}

	client := conf.Client(ctx, token)
	if err = oauthcli2.GetUser(client); err != nil {
		panic(err)
	}
}
