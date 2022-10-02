package async

// FetchAll 함수는 url의 목록을 받는다.
func FetchAll(urls []string, c *Client) {
	for _, url := range urls {
		go c.AsyncGet(url)
	}
}
