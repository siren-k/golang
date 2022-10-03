package waitgroup

import (
	"log"
	"sync"
	"time"
)

// Crawl 전달된 url 목록으로부터 응답을 수집한다.
// 값을 반환하기 전에 모든 요청에 대한 처리를 완료할 땍까지 대기한다.
func Crawl(sites []string) ([]int, error) {
	start := time.Now()
	log.Printf("starting crawling")
	wg := &sync.WaitGroup{}

	var resps []int
	cerr := &CrawlError{}
	for _, v := range sites {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			resp, err := GetURL(v)
			if err != nil {
				cerr.Add(err)
				return
			}
			resps = append(resps, resp.StatusCode)
		}(v)
	}
	wg.Wait()

	// crawl 오류를 발생시킨다.
	if cerr.Present() {
		return resps, cerr
	}
	log.Printf("completed crawling n %s", time.Since(start))

	return resps, nil
}
