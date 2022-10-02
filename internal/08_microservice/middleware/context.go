package middleware

import (
	"context"
	"net/http"
	"strconv"
)

// ContextID는 컨텍스트 객체를 조회하기 위한 타입이다.
type ContextID int

// ID는 예제에서 정의한 유일한 ID다
const ID ContextID = 0

// SetID 함수는 id로 컨택스트를 갱신한 다음 id를 증가시킨다.
func SetID(start int64) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ID, strconv.FormatInt(start, 10))
			start++
			r = r.WithContext(ctx)
			next(w, r)
		}
	}
}

// GetID 함수는 ID가 설정된 경우 컨텍스트에서 ID를 가져오고
// ID가 설정되지 않았으면 빈 문자열을 반환한다.
func GetID(ctx context.Context) string {
	if val, ok := ctx.Value(ID).(string); ok {
		return val
	}
	return ""
}
