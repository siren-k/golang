package main

import (
	mail2 "golang/chap5/mail"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

// 이메일 메시지의 예
const msg string = `Date: Thu, 24 Jul 2019 08:00:00 -0700
From: Aaron <fake_sender@example.com>
To: Reader <fake_receiver@example.com>
Subject: Gophercon 2019 is going to be awesome!

Feel free to share my book with others if you're attending.
This recipe can be used to process and parse email information.
`

// 초기 이메일을 파싱하고, 파싱한 결과를 헤더와 함께 PrintHeaderInfo 함수에 전달한다.
func main() {
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	mail2.PrintHeaderInfo(m.Header)

	// 헤더를 출력한 다음, 표준 출력(stdout)에 본문을 보낸다.
	if _, err := io.Copy(os.Stdout, m.Body); err != nil {
		log.Fatal(err)
	}
}
