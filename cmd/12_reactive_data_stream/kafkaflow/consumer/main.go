package main

import (
	"github.com/Shopify/sarama"
	flow "github.com/trustmaster/goflow"
	"golang/internal/12_reactive_data_stream/kafkaflow"
)

// 이 에제는 이전 예제들의 아이디어를 합쳤다. 이전 예제에서는 카프카 소비자와 생산자를 설정했다. 이번 예제는 '카프카와 Sarama 예제 사용하기' 예제의
// 동기 생성자(synchronous producer)를 사용하지만, 비동기 생산자를 대신 사용할 수도 있다. '데이터 흐름 프로그래밍에 GoFlow 사용하기' 예제에서
// 했듯이, 이 메시지를 받으면 in 채널의 큐에 메시지를 추가한다. 들어오는 문자열을 Base64 인코딩하는 대신 대문자로 변환하도록 컴포넌트를 수정했다.
// 출력 컴포넌트를 재사용했으며 네트워크 구성은 비슷하게 설정했다.
//
// 결과적으로 카프카 소비자를 통해 받은 메시지는 모두 흐름 기반(flow-based) 작업 파이프라인으로 전송된다. 이를 통해 파이프라인 컴포넌트를 모듈화하고
// 재사용할 수 있으며, 같은 컴포넌트를 다른 환경에서 여러 번 사용할 수 있다. 마찬가지로 카프카에 쓰는 모든 생산자로부터 트래픽을 수신하기 때문에 여러
// 생산자를 하나의 데이터 스트림에 연결할 수 있다.
func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	net := kafkaflow.NewUpperApp()

	in := make(chan string)
	net.SetInPort("In", in)

	wait := flow.Run(net)
	defer func() {
		close(in)
		<-wait
	}()

	for {
		msg := <-partitionConsumer.Messages()
		in <- string(msg.Value)
	}
}
