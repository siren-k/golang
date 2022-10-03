package mocking

type MockDoStuffer struct {
	// 모의 객체를 활용한 테스트를 위한 클로저
	MockDoStuff func(input string) error
}

func (m *MockDoStuffer) DoStuff(input string) error {
	if m.MockDoStuff != nil {
		return m.MockDoStuff(input)
	}

	// 모의 테스트를 진행하지 않을 때는 일반적인 방법으로 반환한다.
	return nil
}
