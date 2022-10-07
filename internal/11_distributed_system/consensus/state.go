package consensus

type state string

const (
	first  state = "first"
	second       = "second"
	third        = "third"
)

var allowedState map[state][]state

func init() {
	// 유효한 상태를 설정한다.
	allowedState = make(map[state][]state)
	allowedState[first] = []state{second, third}
	allowedState[second] = []state{third}
	allowedState[third] = []state{first}
}

// CanTransition 함수는 새 상태가 유효한지 확인한다.
func (s *state) CanTransition(next state) bool {
	for _, n := range allowedState[*s] {
		if n == next {
			return true
		}
	}
	return false
}

// Transition 함수는 가능한 경우 한 상태르 다음 상태로 이동시킨다.
func (s *state) Transition(next state) {
	if s.CanTransition(next) {
		*s = next
	}
}
