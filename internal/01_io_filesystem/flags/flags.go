package flags

import (
	"flag"
	"fmt"
)

// Config 구조체는 플래그 값을 저장하는데 사용한다.
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup 함수는 전달되는 플래그 값으로 설정(config) 값을 초기화한다.
func (c *Config) Setup() {
	// 다음과 같이 flag를 직접 설정할 수 있다.
	// var someVar = flag.String("flag_name", "defualt_val", "description")
	// 하지만 구조체에 값을 넣는 것이 일반적으로 더 낫다.
	flag.StringVar(&c.subject, "subject", "", "subject is a string, id defaults to empty")
	// 축약어(단축어)
	flag.StringVar(&c.subject, "s", "", "subject is string, it defaults to empty (shorthand)")
	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome of what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")
	// 커스텀 변수 타입
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage 함수는 모든 내부 설정(config) 변수를 사용해 문장을 반환한다.
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}
	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
