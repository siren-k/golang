package client

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"strings"
)

func Process(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter some text: ")

		// ctrl-c를 막기 때문에 종료하려면 ctrl-c를 누르고 엔터키를 누르거나
		// 다른 위치에서 종료시킬 수 있다.
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read stdin", err)
		}

		// 읽어온 문자열에서 공백(space)를 제거한다.
		data = strings.TrimSpace(data)

		// 웹소켓에 메시지를 바이트로 작성한다.
		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message: ", err)
			return
		}

		// 에코 서버이므로 작성한 후에 읽어올 수 있다.
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("failed to read: ", err)
			return
		}

		log.Printf("received back from server: %#v\n", string(message))
	}
}
