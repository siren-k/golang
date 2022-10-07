package consensus

import (
	"fmt"
	"github.com/hashicorp/raft"
)

// 나중을 위해 raft의 맵을 유지한다.
var rafts map[raft.ServerAddress]*raft.Raft

func init() {
	rafts = make(map[raft.ServerAddress]*raft.Raft)
}

// reftSet은 필요한 모든 설정 값을 저장한다.
type raftSet struct {
	Config        *raft.Config
	Store         *raft.InmemStore
	SnapShotStore raft.SnapshotStore
	FSM           *FSM
	Transport     raft.LoopbackTransport
	Configuration raft.Configuration
}

// raft 클러스터를 실행시키기 위해 n개의 raft 설정을 생성한다.
func getRaftSet(num int) []*raftSet {
	rs := make([]*raftSet, num)
	servers := make([]raft.Server, num)
	for i := 0; i < num; i++ {
		addr := raft.ServerAddress(fmt.Sprint(i))
		_, transport := raft.NewInmemTransport(addr)
		servers[i] = raft.Server{
			Suffrage: raft.Voter,
			ID:       raft.ServerID(addr),
			Address:  addr,
		}
		config := raft.DefaultConfig()
		config.LocalID = raft.ServerID(addr)

		rs[i] = &raftSet{
			Config:        config,
			Store:         raft.NewInmemStore(),
			SnapShotStore: raft.NewInmemSnapshotStore(),
			FSM:           NewFSM(),
			Transport:     transport,
		}
	}

	// 구성 값(configuration)은 서비스 사이에서 일관되게 유지해야 하므로
	// 이 경우, 전체 서버 목록이 필요하다.
	for _, r := range rs {
		r.Configuration = raft.Configuration{Servers: servers}
	}

	return rs
}
