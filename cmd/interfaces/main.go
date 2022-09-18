package main

import (
	"bytes"
	"fmt"
	"golang/interfaces"
)

// Go 인터페이스는 일반적인 작업을 처리하는 데이터를 래핑(wrapping)하는 깔끔한 추상화를 제공한다.
// 이런 특성은 I/O 작업을 수행할 때 분명하게 나타난다. 따라서 io 패키지는 인터페이스 구성을 학습하는데
// 훌룡한 재료라고 할 수 있다. pipe 패키지는 덜 사용되는 면이 있지만, 입력 스트림과 출력 스트림을 연결할 때
// 스레드 안정성(thread safe)과 함께 뛰어난 유연성을 제공한다.
func main() {
	// Copy() 함수는 인터페이스 사이에서 바이트를 복사하고, 복사한 데이터를 스트림으로 취급한다.
	// 데이터를 스트림으로 생각하면 매우 실용적인데, 특히 네트워크 트래픽이나 파일 시스템을 활용한
	// 작업에 유용하다.
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	// PipeReader와 PipeWriter 구조체는 io.Reader와 io.Writer 인터페이스를 구현한다.
	// 이 둘은 연결돼 있으며 인메모리 파이프를 생성한다. 이 파이프의 주목적은 동일한 스트림에서
	// 다른 소스로 쓰면서 동시에 스트림으로부터 읽는 것이다. 본직적으로 두 스트림을 파이프로 결합시킨다.
	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on  PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
