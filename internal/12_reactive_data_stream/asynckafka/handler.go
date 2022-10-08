package asynckafka

import (
	"github.com/Shopify/sarama"
	"net/http"
)

// KafkaController는 producer를 핸들러에 연결할 수 있도록 해준다.
type KafkaController struct {
	Producer sarama.AsyncProducer
}

// Handler 함수는 GET 파라미터로부터 메시지를 가져와 이 메시지를 kafka 큐에 비동기 방식으로 보낸다.
func (c *KafkaController) Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("msg must be set"))
		return
	}

	c.Producer.Input() <- &sarama.ProducerMessage{
		Topic: "example",
		Key:   nil,
		Value: sarama.StringEncoder(msg),
	}
	w.WriteHeader(http.StatusOK)
}
