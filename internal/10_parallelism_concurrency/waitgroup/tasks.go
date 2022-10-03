package waitgroup

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// GetURL 함수는 url을 받고 경과 시간을 로그로 기록한다.
func GetURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("getting %s", url)
	resp, err := http.Get(url)
	log.Printf("completed getting %s in %s", url, time.Since(start))
	return resp, err
}

// CrawlError 오류를 집계하기 위한 사용자 정의 오류 타입이다.
type CrawlError struct {
	Errors []string
}

// Add 또 다른 오류를 추가한다.
func (c *CrawlError) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}

// Error error 인터페이스를 구현한다.
func (c *CrawlError) Error() string {
	return fmt.Sprintf("All Errors: %s", strings.Join(c.Errors, ", "))
}

// Present 오류를 반환할지 여부를 결정하는데 사용할 수 있다.
func (c *CrawlError) Present() bool {
	return len(c.Errors) != 0
}
