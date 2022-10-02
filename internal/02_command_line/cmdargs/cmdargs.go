package cmdargs

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"
const usage = `Usage:
%s [command]
Commands:
    Greet
    Version
`
const greetUsage = `Usage:
%s greet name [flag]
Positional Arguments
    name
        the name to greet
Flags:
`

// MenuConf 구조체는 중첩된 명령줄에 대한 모든 레벨을 저장한다.
type MenuConf struct {
	Goodbye bool
}

// SetupMenu 함수는 기본 플래그를 설정한다.
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}

// GetSubMenu 함수는 하위 메뉴에 대한 플래그 모음을 반환한다.
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.Goodbye, "goodbye", false, "Say goodbye instead of hello")
	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}
	return submenu
}

// Greet 함수는 greet 명령에 의해 호출된다.
func (m *MenuConf) Greet(name string) {
	if m.Goodbye {
		fmt.Println("Goodbye " + name + "!")
	} else {
		fmt.Println("Hello " + name + "!")
	}
}

// Version 함수는 상수(const)로 저장된 현재 버전을 출력한다.
func (m *MenuConf) Version() {
	fmt.Println("Version: " + version)
}
