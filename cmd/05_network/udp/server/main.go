package main

import (
	"fmt"
	"golang/internal/05_network/udp"
	"net"
)

const addr = "localhost:8888"

func main() {
	conns := &udp.Connections{
		Addrs: make(map[string]*net.UDPAddr),
	}
	fmt.Printf("serving on %s\n", addr)

	// udp addr을 생성한다.
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	// 지정된 addr(주소)로 요청을 대기한다
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	// 접속을 종료한다.(지연 호출)
	defer conn.Close()

	// 알려진 모든 클라이언트에 비동기로 메시지를 전송한다.
	go udp.Broadcast(conn, conns)

	msg := make([]byte, 1024)
	for {
		// 메시지를 다시 보내기 위한 주소와 포트 정보의 수집을 위해 메시지를 수신한다.
		_, retAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}

		// 수신한 메시지를 맵에 저장한다.
		conns.Mu.Lock()
		conns.Addrs[retAddr.String()] = retAddr
		conns.Mu.Unlock()
		fmt.Printf("%s connected\n", retAddr)
	}
}
