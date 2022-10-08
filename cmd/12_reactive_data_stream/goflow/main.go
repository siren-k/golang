package main

import (
	"fmt"
	flow "github.com/trustmaster/goflow"
	"golang/internal/12_reactive_data_stream/goflow"
)

// github.com/trustmaster/goflow 패키지는 네트워크(또는 그래프)를 정의하고 여러 컴포넌트를 등록한 다음, 이 컴포넌트들을
// 연결하는 방식으로 동작한다. 문자열을 사용해 컴포넌트를 정의하기 때문에 오류가 발생할 수 있지만, 일번작으로 애플리케이션이 설정되고
// 제대로 동작할 때까지 런타임 초기에 장애가 발생하므로 대응이 가능하다.
//
// 이 예제에서는 두 개의 컴포넌트를 설정했다. 하나는 전달받은 문자열을 Base64로 인코딩하는 컴포넌트이고, 다른 하나는 전달된 모든
// 내용을 출력하는 컴포넌트다. 두 컴포넌트를 main.go에서 초기화한 in 채널에 연결했다. 이를 통해 이 채널(in 채널)로 전달되는
// 모든 내용이 예제의 파이프라인을 통해 흘러가게 된다.
//
// 이 접근 방법에서 많이 중점을 둔 부분은 내부에서 무슨 일이 일어나고 있는지를 무시하는 것이다. 모든 것을 연결된 블랙박스처럼 다루고
// 나머지를 goflow에 맡긴다. 에제를 통해 이 작업 파아프라인을 달성하기 위한 코드의 양이 얼마나 적은지를 볼 수 있다. 특히 다른
// 것보다 워커의 수를 제어하는 내용이 얼마나 작은지 확인할 수 있다.
func main() {
	net := goflow.NewEncodingApp()

	in := make(chan string)
	net.SetInPort("In", in)

	wait := flow.Run(net)

	for i := 0; i < 20; i++ {
		in <- fmt.Sprint("Message", i)
	}

	close(in)
	<-wait
}
