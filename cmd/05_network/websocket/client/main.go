package main

import (
	"github.com/gorilla/websocket"
	"golang/internal/05_network/websocket/client"
	"log"
	"os"
	"os/signal"
)

// catchSig 함수는 ctrl-c로 프로그램을 종료하면 웹소켓 연결을 정리한다.
func catchSig(ch chan os.Signal, c *websocket.Conn) {
	// os.Signal을 대기하기 위해 블록한다.
	<-ch

	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
	}
	return
}

// 클라이언트에서는 Ctrl + C 이벤트를 처리하기 위해 catchSig 함수를 추가했다. 이를 통해 클라이언트가 종료될 때
// 서버와의 연결을 깔끔하게 종료할 수 있다 클라이언트는 표준 입력(STDIN)에서 사용자 입력을 받은 다음, 이를 서버로
// 보내고 응답을 로그로 기록한 후 이 작업을 반복한다.
func main() {
	// 채널에 os.Signal 값을 연결한다.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 웹소켓에 연결하기 위해 ws:// 구조를 사용한다.
	u := "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// catchSig에 전달한다.
	go catchSig(interrupt, c)

	client.Process(c)
}
