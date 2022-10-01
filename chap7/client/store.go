package client

import (
	"fmt"
	"net/http"
)

// Controller는 http.Client 객체를 가지며 이를 내부에서 사용한다.
type Controller struct {
	*http.Client
}

// DoOps 함수는 Controller 객체를 매개변수로 받아 작업을 수행한다.
func (c *Controller) DoOps() error {
	resp, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results of client.DoOps", resp.StatusCode)
	return nil
}
