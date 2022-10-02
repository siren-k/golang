package controllers

// Controllers는 핸들러의 상태를 전달한다.
type Controller struct {
	storage Storage
}

// New 함수는 Controller의 '생성자(constructor)'다.
func New(storage Storage) *Controller {
	return &Controller{
		storage: storage,
	}
}

// Payload는 일반적인 응답에 사용한다.
type Payload struct {
	Value string `json:"value"`
}
