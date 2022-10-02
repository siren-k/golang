package main

import (
	"golang/internal/01_io_filesystem/bytestrings"
)

// Go는 기본 유형을 사용할 때 인터페이스 간에 데이터를 변환하는 과정에서 다양한 유연성을 제공한다.
// 예를 들어, 문자열과 바이트 사이의 변환은 비교적 간단하다. 문자열을 처리하는 경우, strings 패키지는
// 검색, 문자열의 조작 등 다양한 편의 기능을 제공한다. 어떤 때는 정규 표현식이 적절할 수도 있지만,
// 대부분 strings와 strconv 패키지로도 충분하다. strings 패키지는 문자열을 제목처럼 나타내거나,
// 문자열을 배열로 분할하거나, 공백을 없애는 기능을 제공한다. 또한 bytes 패키지의 reader 타입 대신에
// 사용할 수 있는 자체 reader 인터페이스를 제공한다.
func main() {
	// bytes 라이브러리는 데이터를 처리할 때 유용한 여러 현의 기능을 제공한다. 예를 들어 버퍼(buffer)는
	// streamprocessing 라이브러리나 메소드(method)로 작업할 때 바이트 배열보다 휠씬 더 유연하다.
	// 일단 버퍼를 생성하면 io.Reader 인터페이스를 충족시키기 때문에 ioutil 함수를 사용해 데이터를
	// 조작할 수 있다. 스트리밍 애플리케이션의 경우 버퍼와 스캐너(scanner)를 사용하고 싶을 텐데,
	// bufio 패키지는 이런 경우에 유용하다. 때로 매우 작은 데이터 세트인 경우나 컴퓨터의 메모리 용량이
	// 클 때는 배열이나 슬라이스(slice)를 사용하는 것이 더 적합하다.
	if err := bytestrings.WorkWithBuffer(); err != nil {
		panic(err)
	}

	// 아래 내용을 모두 표준 출력(Stdout)에 출력한다.
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
