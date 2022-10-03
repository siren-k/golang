package mockgen

// GetSetter는 키-값 쌍의 데이터 모음을 가져온다.(Get)
type GetSetter interface {
	Set(key, val string) error
	Get(key string) (string, error)
}
