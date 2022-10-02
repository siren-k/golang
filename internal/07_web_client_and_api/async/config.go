package async

import "net/http"

// NewClient 함수는 새 클라이언트를 생성하고 클라이언트에 적절한 채널을 설정한다.
func NewClient(client *http.Client, bufferSize int) *Client {
	respch := make(chan *http.Response, bufferSize)
	errch := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp:   respch,
		Err:    errch,
	}
}

// Client 구조체는 클라이언트를 저장하고 응답과 오류 정보를 가져오기 위한 두 채널을 가진다.
type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

// AsyncGet 함수는 Get 요청을 수행한 다음, 적합한 채널에 응답/오류를 반환한다.
func (c *Client) AsyncGet(url string) {
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}
