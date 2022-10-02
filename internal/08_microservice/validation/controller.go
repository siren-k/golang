package validation

// Controller는 검증 함수를 갖는다.
type Controller struct {
	ValidatePayload func(p *Payload) error
}

// New 함수는 로컬 함수를 갖는 컨트롤러를 초기화하며 덮어 쓰기가 가능하다.
func New() *Controller {
	return &Controller{
		ValidatePayload: ValidatePayload,
	}
}
