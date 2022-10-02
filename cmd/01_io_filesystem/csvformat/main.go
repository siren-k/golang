package main

import (
	"fmt"
	"golang/internal/01_io_filesystem/csvformat"
)

// csv 패키지는 Go의 데이터 흐름(data flow)을 공용 인터페이스로 구현하는 것으로 간주하려는
// 이유를 보여주는 훌륭한 에제다. 단 한 줄로 원하는 데이터로 쉽계 변경할 수 있으며, 과도한
// 메모리나 시간을 소모하지 않고도 csv 데이터 포맷을 쉽게 조작할 수 있다. 예를 들어,
// 데이터 스트림에서 한 번에 하나의 레코드(record)를 읽고 한 번에 하나의 레코드를 수정된
// 포맷으로 별도의 스트림에 쓸 수 있다. 이렇게 하면 메모리나 프로세스 사용량이 크게 증가하지 않을 겻이다.
//
// 나중에 데이터 파이프라인과 워커 풀(worker pool)을 살펴볼 때, 이 개념들을 결합하는 방법과
// 이 스트림들을 병렬로 제어하는 방법을 확인할 수 있다.
func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}
	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}
	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Println("Buffer = ", buffer.String())
}
