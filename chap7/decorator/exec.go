package decorator

import "fmt"

// Exec 함수는 클라이언트를 생성하고 google.com에 Get 요청을 한 다음 응답을 출력한다.
func Exec() error {
	c := Setup()

	resp, err := c.Get("https://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("Response code:", resp.StatusCode)

	return nil
}
