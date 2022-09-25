package encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// Base64Example 함수는 base64 패키지를 사용하는 데모를 보여준다.
func Base64Example() error {
	// base64 패키지는 바이트/문자열로 동작하는
	// 바이너리 포맷을 지원할 수 없을 때 유용하다.
	// 도움 함수와 URL 인코딩을 사용한다.
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodeToString and URLEncoding: ", value)

	// 첫 번째 값을 디코딩한다.
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeToString and URLEncoding: ", string(decoded))

	return nil
}

// Base64ExampleEncoder 함수는 encoder/decoder를 사용해
// 비슷한 예제를 보여준다.
func Base64ExampleEncoder() error {
	// encoder/decoder 사용하기
	buffer := bytes.Buffer{}

	// 버퍼로 인코딩한다.
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)
	if _, err := encoder.Write([]byte("encoding some other data")); err != nil {
		return err
	}

	// 인코더를 닫았는지 확인한다.
	if err := encoder.Close(); err != nil {
		return err
	}
	fmt.Println("Using encoder and StdEncoding: ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}
	fmt.Println("Using decoder and StdEncoding: ", string(results))

	return nil
}
