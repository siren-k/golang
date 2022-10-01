package rest

import "net/http"

// APIClient는 예제에서 사용할 사용자 정의 클라이언트다
type APIClient struct {
	*http.Client
}

// NewAPIClient 생성자는 사용자 정의 Transport 값으로 클라이언트를 초기화한다.
func NewAPIClient(username, password string) *APIClient {
	t := http.Transport{}
	return &APIClient{
		Client: &http.Client{
			Transport: &APITransport{
				Transport: &t,
				username:  username,
				password:  password,
			},
		},
	}
}

// GetGoogle 함수는 API 호출이다.
// REST 관련 내용은 추상화한다.
func (c *APIClient) GetGoogle() (int, error) {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
