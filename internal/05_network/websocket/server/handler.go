package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// upgrader는 http 연결을 가져와 웹소켓으로 변환한다.
// 여기서는 권장 버퍼 크기를 사용한다.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHanlder(w http.ResponseWriter, r *http.Request) {
	// 연결을 업그레이드한다.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection: ", err)
		return
	}

	for {
		// 루프에서 메시지를 읽고 동일한 메시지를 전달한다.
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message: ", err)
			return
		}

		log.Printf("received from client: %#v", string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("failed to write message: ", err)
			return
		}
	}
}
