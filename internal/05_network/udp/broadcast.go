package udp

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Connections struct {
	Addrs map[string]*net.UDPAddr
	// 맵의 수정을 위해 락(lock)시킨다.
	Mu sync.Mutex
}

func Broadcast(conn *net.UDPConn, conns *Connections) {
	count := 0
	for {
		count++
		conns.Mu.Lock()
		// 알려진(확인한) 주소에 대해 루프로 반복 처리한다.
		for _, retAddr := range conns.Addrs {
			// 모두에게 메시지를 전송한다.
			msg := fmt.Sprintf("Sent %d", count)
			if _, err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encountered: %s\n", err.Error())
				continue
			}
		}
		conns.Mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
