package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer 함수는 Buffer 함수에서 생성한 버퍼를 사용할 것이다.
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)
	// b.Bytes()를 사용해 버퍼를 바이트로 다시 전환하거나
	// b.String()을 사용해 버퍼를 문자열로 다시 빠르게 변환할 수 있다.
	fmt.Println(b.String())
	// io Reader이므로 다음과 같은
	// 일반적인 io reader 인터페이스 함수를 사용할 수 있다.
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)
	// 또한 바이트를 가져와 bytes reader를 생성할 수 있으며,
	// 이 reader들은 io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner,
	// io.RuneScanner 인터페이스를 구현한다.
	reader := bytes.NewReader([]byte(rawString))
	// 버퍼링 처리된 읽기 및 토큰화(tokenization)를 허용하는
	// 스캐너(Scanner)에 연결할 수도 있다.
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	// 루프를 통해 모든 스캔 이벤트를 처리한다.
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
	return nil
}
