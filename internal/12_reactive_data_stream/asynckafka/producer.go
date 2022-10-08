package asynckafka

import (
	"github.com/Shopify/sarama"
	"log"
)

// ProcessResponse 함수는 producer로부터 결과와 오류를 비동기로 받는다.
func ProcessResponse(producer sarama.AsyncProducer) {
	for {
		select {
		case result := <-producer.Successes():
			log.Printf("> message: \"%s\" sent to partition %d at offset %d\n", result.Value, result.Partition, result.Offset)
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
		}
	}
}
