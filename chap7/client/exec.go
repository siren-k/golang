package client

import (
	"fmt"
	"net/http"
)

// DoOps 함수는 클라이언트를 매개변수로 받은 다음,
// google.com에 연결한다.
func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results of DoOps:", resp.StatusCode)

	return nil
}

// DefaultGetGolang 함수는 golang.org의 주소로 Get 요청을 하기 위해
// 기본 클라이언트를 사용한다.
func DefaultGetGolang() error {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}
	fmt.Println("results of DefaultGetGolang", resp.StatusCode)

	return nil
}
