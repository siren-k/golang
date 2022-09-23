package main

import (
	"bytes"
	"fmt"
	"golang/envvar"
	"io/ioutil"
	"os"
)

// Config 구조체는 json 파일과 환경 변수로부터 읽어온 구성 값을 저장한다.
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is-safe" required:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	// 예제 json 파일을 저장하기 위해
	// 임시 파일을 생성한다.
	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	// secrets 값을 저장하기 위해
	// json 파일을 생성한다.
	secrets := `{
		"secret": "so os secret"
	}`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// 필요한 경우 환경 변수를 쉽게 설정할 수 있다.
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	// 또한 환경 변수를 쉽게 읽을 수 있다.
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE=", os.Getenv("EXAMPLE_ISSAFE"))

	// 최종 구성(config)은 json과 환경 변수가 혼합돼 있다.
	fmt.Printf("Final Config: %#v\n", c)
}
