package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"golang/internal/12_reactive_data_stream/asynckafka"
	"net/http"
)

// $ curl "http://localhost:3333/?msg=this"
// $ curl "http://localhost:3333/?msg=is"
// $ curl "http://localhost:3333/?msg=an"
// $ curl "http://localhost:3333/?msg=example"
//
// Listening on port :3333
// 2022/10/08 21:16:17 > message: "this" sent to partition 0 at offset 10
// 2022/10/08 21:16:23 > message: "is" sent to partition 0 at offset 11
// 2022/10/08 21:16:29 > message: "an" sent to partition 0 at offset 12
// 2022/10/08 21:16:34 > message: "example" sent to partition 0 at offset 13

//
func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.AsyncClose()

	go asynckafka.ProcessResponse(producer)

	c := asynckafka.KafkaController{Producer: producer}
	http.HandleFunc("/", c.Handler)

	fmt.Println("Listening on port :3333")
	panic(http.ListenAndServe(":3333", nil))
}
