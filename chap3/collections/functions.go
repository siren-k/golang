package collections

import "strings"

// LowerCaseData 함수는 WorkWith의 Data 문자열에 ToLower를 수행한다.
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion 함수는 WorkWith의 버전을 증가시킨다.
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// OldVersion 함수는 버전의 지정된 값보다 큰지 확인하는 클로저를 반환한다.
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}
