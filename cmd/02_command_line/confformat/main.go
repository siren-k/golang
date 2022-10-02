package main

import (
	"golang/internal/02_command_line/confformat"
)

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExample(); err != nil {
		panic(err)
	}
}
