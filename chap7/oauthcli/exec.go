package oauthcli

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetUsers 함수는 초기화된 oauth2 클라이언트를 사용해
// 사용자의 정보를 가져온다.
func GetUser(client *http.Client) error {
	url := fmt.Sprintf("https://api.github.com/user")

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status Code from", url, ":", resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)

	return nil
}
