package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888"

func echoBackCapitalized(conn net.Conn) {
	// conn에 리더(reader)를 설정한다.(io.Reader)
	reader := bufio.NewReader(conn)

	// 읽어온 데이터의 첫 줄을 가져온다.
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}

	// 출력한 다음 데이터를 다시 보낸다.
	fmt.Printf("Received: %s", data)
	conn.Write([]byte(strings.ToUpper(data)))

	// 완료된 연결을 종료한다.
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Printf("listening on: %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("encountered an error accepting connection: %s\n", err.Error())
			// 오류가 있으면 다시 시도한다.
			continue
		}

		// 이 작업을 비동기로 처리하면 잠재적 워커 풀을 위해 좋은 사용 사례가 될 것이다.
		go echoBackCapitalized(conn)
	}
}
