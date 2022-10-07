package main

import (
	consul "github.com/hashicorp/consul/api"
	"golang/internal/11_distributed_system/discovery"
)

// $ consul agent -dev -node=localhost

// Consul은 강력한 Go API 라이브러리를 제공한다. Consul을 처음 시작할 때는 기능이 너무 많기 때문에 위압감을 느낄 수 있지만,
// 이 예제는 Consul을 래핑하는 방법을 보여준다. Consul을 추가로 구성하는 것은 이 예제의 범위를 벗어난다. 예제는 서비스를 등록하고
// 다른 서비스를 요청하는 기본적인 내용에 중점을 뒀다.
//
// Consul을 사용해 시작 시점에 새로운 마이크로서비스를 등록하고, 이와 의존 관계에 있는 모든 서비스를 요청하고, 종료할 때 등록을
// 해제하는 것이 가능하다. 모든 요청에 대해 Consul에 접근하지 않고록 이런 정보를 캐시에 저장할 수도 있지만, 이 예제는 여러분이 나중에
// 확장할 수 있도록 기본적인 도구만 제공한다. Consul 에이전트는 이런 요청을 빠르고 효율적으로 처리한다.
// (https://www.consul.io/intro/getting-started/agent.html)
func main() {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"

	cli, err := discovery.NewClient(config, "localhost", "discovery", 8080)
	if err != nil {
		panic(err)
	}

	if err = discovery.Exec(cli); err != nil {
		panic(err)
	}
}
