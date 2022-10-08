package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func sendMessage(producer sarama.SyncProducer, value string) {
	msg := &sarama.ProducerMessage{Topic: "example", Value: sarama.StringEncoder(value)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
		return
	}
	log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
}

// 이 예제는 카프카를 통해 단순한 메시지를 전달하는 방법을 보여준다. 더 복잡한 메시지는 json, gob, progobuf 등과 같은 직렬화 모팻을 사용해야 한다.
// 생산자(producer)는 sendMessage 함수를 통해 동기 방식으로 카프카에 메시지를 보낸 수 있다. 여기서는 카프카 클러스터가 다운될 수 있는 각각의
// 경우를 처리하지 않기 때문에 프로세스가 정지될 수 있다. 이는 타임아웃과 카프카에 대한 강한 의존성 문제가 발생할 수 있으므로 웹 핸들러와 같은 애플리케이션에서
// 고려해야 할 사항이다.
//
// 메시지 큐가 정상적으로 동작한다고 가정하면 소비자(consumer)는 카프카 스트림을 관찰하고 결과를 내기 위해 어떤 작업을 수행한다. 이 장의 이전 예제라면
// 이 스트림을 사용해 필요한 적업을 처리할 수 있을 것이다.
func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for i := 0; i < 10; i++ {
		sendMessage(producer, fmt.Sprintf("Message %d", i))
	}
}
