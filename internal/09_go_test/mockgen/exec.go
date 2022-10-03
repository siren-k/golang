package mockgen

// Controller는 인터페이스를 초기화하는 한 방법을 보여주는 구조체다.
type Controller struct {
	GetSetter
}

// GetThenSet 함수는 값을 확인한 다음 설정한다.
// 설정되 않은 경우 값을 설정한다.
func (c *Controller) GetThenSet(key, value string) error {
	val, err := c.Get(key)
	if err != nil {
		return err
	}

	if val != value {
		return c.Set(key, value)
	}

	return nil
}
