package mocking

import (
	"errors"
	"testing"
)

func TestDoSomeStuff(t *testing.T) {
	tests := []struct {
		name       string
		DoStuff    error
		ThrowError error
		wantErr    bool
	}{
		{"base-case", nil, nil, false},
		{"DoStuff error", errors.New("failed"), nil, true},
		{"ThrowError error", nil, errors.New("failed"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 모의 구조체를 활용해 인터페이스를 테스트하는 예제
			d := MockDoStuffer{}
			d.MockDoStuff = func(string) error {
				return tt.DoStuff
			}

			// 변수로 선언된 함수를 모의 객체를 활용해 테스트하면
			// func A()는 동작하지 않고
			// 반드시 var A = func()로 해야 동작한다.
			defer Patch(&ThrowError, func() error {
				return tt.ThrowError
			}).Restore()

			if err := DoSomeStuff(&d); (err != nil) != tt.wantErr {
				t.Errorf("DoSomeStuff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
