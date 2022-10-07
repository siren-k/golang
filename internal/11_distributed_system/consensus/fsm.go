package consensus

import (
	"github.com/hashicorp/raft"
	"io"
)

// FSM은 raft FSM 인터페이스를 구현하고 상태를 저장한다.
type FSM struct {
	state state
}

// NewFSM 함수는 "first"라는 시작 상태 값을 갖는 새로운 FSM을 생성한다.
func NewFSM() *FSM {
	return &FSM{state: first}
}

// Apply 함수는 FSM을 갱신한다.
func (f *FSM) Apply(r *raft.Log) interface{} {
	f.state.Transition(state(r.Data))
	return string(f.state)
}

// Sanpshot 함수는 raft FSM 인터페이스의 요구를 충족하기 위해 필요하다.
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore 함수는 raft FSM 인터페이스의 요구를 충족하기 위해 필요하다.
func (f *FSM) Restore(closer io.ReadCloser) error {
	return nil
}
