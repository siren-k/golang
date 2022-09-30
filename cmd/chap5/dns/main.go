package main

import (
	"fmt"
	"golang/chap5/dns"
	"log"
	"os"
)

// 제공된 주소에서 CNAME과 host 검색을 수행한다. 예제에서는 golang.org를 검색했다.
// String() 함수를 사용해 결과를 출력하는 lookup 구조체에 결과를 저장한다. String() 함수는
// 객체를 문자열로 출력할 때 자동으로 호출되며 직접 함수를 호출할 수도 있다. 프로그램을
// 실행할 때 주소가 제공될 수 있도록 main.go 파일에 기본적인 매개변수 확인 기능을 구현했다.
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panicf("failed to lookup: %s", err.Error())
	}

	fmt.Println(lookup)
}
