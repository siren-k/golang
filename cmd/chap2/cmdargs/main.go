package main

import (
	"fmt"
	"golang/chap2/cmdargs"
	"os"
	"strings"
)

func main() {
	c := cmdargs.MenuConf{}
	menu := c.SetupMenu()
	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	// 명령을 전환하는데 매개변수를 사용한다.
	// 플래그 또한 매개변수로 사용된다.
	if len(os.Args) > 1 {
		// 대소문자 구별은 하지 않는다.
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing params: %s, error: %v", os.Args[3:], err)
					return
				}
			}
		}
	}
}
