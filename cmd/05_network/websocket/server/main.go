package main

import (
	"fmt"
	"golang/internal/05_network/websocket/server"
	"log"
	"net/http"
)

// 웹소켓 연결을 위해 8000번 포트에서 요청을 대기한다. 요청이 들어오면 githum.com/gorilla/websocket 패키지를 사용해
// 해당 요청을 웹소켓 연결로 업그레이드한다. 이전의 에코 서버 예제와 비슷하게 서버는 웹소켓 연결에서 메시지를 기다리고 클라이언트에
// 동일한 메시지로 응답한다. 핸들러이기 때문에 많은 웹소켓을 비동기로 처리할 수 있고 클라이언트가 종료될 때까지 연결을 유지한다.
func main() {
	fmt.Println("Listening on port :8000")
	// 모든 요청을 처리하기 위해 localhost:8000번 포트에
	// 하나의 핸드러를 마운트한다.
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(websocket.WsHanlder)))
}
