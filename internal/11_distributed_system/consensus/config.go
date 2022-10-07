package consensus

import "github.com/hashicorp/raft"

// Config 함수는 num개의 인메모리 raft 노드를 생성하고 이들을 연결한다.
func Config(num int) {
	// 노드를 표현하는데 필요한 모든 내용으로 구성된 n개의 "raft-sets"를 생성한다.
	rs := getRaftSet(num)

	// 모든 통신 수단(transport)을 연결한다.
	for _, r1 := range rs {
		for _, r2 := range rs {
			r1.Transport.Connect(r2.Transport.LocalAddr(), r2.Transport)
		}
	}

	// 각 노드를 실행하고 연결한다.
	for _, r := range rs {
		if err := raft.BootstrapCluster(r.Config, r.Store, r.Store, r.SnapShotStore, r.Transport, r.Configuration); err != nil {
			panic(err)
		}
		raft, err := raft.NewRaft(r.Config, r.FSM, r.Store, r.Store, r.SnapShotStore, r.Transport)
		if err != nil {
			panic(err)
		}
		rafts[r.Transport.LocalAddr()] = raft
	}
}
