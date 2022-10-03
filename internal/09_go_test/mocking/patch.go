package mocking

import "reflect"

// Restorer는 이전 상태를 복귀하는데 사용하는 함수를 가진다.
type Restorer func()

// Restore 함수는 이전 상태를 복귀한다.
func (r Restorer) Restore() {
	r()
}

// Patch 함수는 매개변수로 가져온 destination이 가리키는 값을 지정된 값으로 설정하고,
// 변환한 값을 원래의 값으로 복원하기 위한 함수를 반환한다.
// 이 값은 destination의 항목의 타입으로 할당이 가능해야 한다.
func Patch(dest, value interface{}) Restorer {
	destv := reflect.ValueOf(dest).Elem()
	oldv := reflect.New(destv.Type()).Elem()
	oldv.Set(destv)
	valuev := reflect.ValueOf(value)
	if !valuev.IsValid() {
		// 이 방법은 destination 타입이 nilable이 아닌 경우에는
		// 적합하지 않지만, 다른 복잡한 방법보다는 좋은 방법이다.
		valuev = reflect.Zero(destv.Type())
	}
	destv.Set(valuev)
	return func() {
		destv.Set(oldv)
	}
}
