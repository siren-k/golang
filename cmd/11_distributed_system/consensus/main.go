package main

import (
	"golang/internal/11_distributed_system/consensus"
	"net/http"
)

// $ curl "http://localhost:3333/?next=second"
// second
//
// $ curl "http://localhost:3333/?next=third"
// third
//
// $ curl "http://localhost:3333/?next=second"
// invalid transition
//
// $ curl "http://localhost:3333/?next=first"
// first

// 애플리케이션이 시작되면 다수의 Raft 객체를 초기화한다. 이 Raft 객체들은 각각 자신의 주소와 전송 수단을 가진다. InmemTransport{} 함수는
// 다른 운송 수단을 연결하기 위한 메소드를 제공하며 Connect() 함수를 호출한다. 이 연결들이 설정되면 Raft 클러스터는 투표(election)를
// 지정한다. Raft 클러스터에서 통신할 때 클라이언트는 반드시 리더(leader)와 통신해야 한다. 예제에서는 하나의 핸들러가 모든 노드와 대화할 수
// 있었다. 따라서 이 핸들러가 Raft 리더의 Apply() 객체 호출을 담당했다. 그러면 다른 모든 노드에서 차례대로 apply()를 실행한다.
//
// InmemTransport{} 함수는 모든 것을 메모리에 상주시켜 투표와 프로세스 시작을 단순화한다. 실제로는 고루틴이 공유 메모리에 자유롭게 접근할 수
// 있기 때문에 테스트를 하고 개념을 증명하는 경우를 제외하면 이 방법은 그리 유용하지 않다. 좀 더 생산적인 구현의 경우, 서비스 인스턴스가 장치들
// 사이에서 통신할 수 있는 HTTP 전송 등의 기능을 사용한다. 이를 위해서는 서비스 인스턴스가 수신 대기 및 서비스를 해야 할 뿐만 아니라, 서비스
// 인스턴스들이 서로를 연결하고 검색할 수 있어야 하기 때문에 추가적인 기록 또는 서비스 검색이 필요할 수 있다.
func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)

	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
