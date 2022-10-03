package mocking

// DoStuffer는 간단한 인터페이스다.
type DoStuffer interface {
	DoStuff(input string) error
}
