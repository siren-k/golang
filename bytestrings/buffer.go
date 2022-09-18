package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer 함수는 bytes 버퍼를 초기화하는 일부 기법을 보여준다.
// 이 버퍼들은 io.Reader 인터페이스를 구현한다.
func Buffer(rawString string) *bytes.Buffer {
	// 원시 바이트(raw bytes)로 인코딩된 문자열에서 시작한다.
	rawBytes := []byte(rawString)
	//원시 바이트나 원래의 문자열에서 버퍼를 생성하는 방법에는 여러 가지가 있다.
	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	// 다른 방법
	b = bytes.NewBuffer(rawBytes)
	// 그리고 초기 바이트 배영을 완전히 피하는 다른 방법
	b = bytes.NewBufferString(rawString)
	return b
}

// ToString 함수는 io.Reader를 가져와 모두 사용한 다음,
// 문자열을 반환하는 예를 보여준다.
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
