package decorator

import (
	"log"
	"net/http"
	"os"
)

// Setup 함수는 ClientInterface를 초기화한다.
func Setup() *http.Client {
	c := http.Client{}
	t := Decorate(
		&http.Transport{},
		Logger(log.New(os.Stdout, "", 0)),
		BasicAuth("username", "password"),
	)
	c.Transport = t

	return &c
}
