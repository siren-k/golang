package main

import (
	"fmt"
	"golang/internal/02_command_line/ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red!",
	}
	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "Now I'm green!"
	fmt.Println(r.String())

	r.TextColor = ansicolor.ColorNone
	r.Text = "Back to normal..."
	fmt.Println(r.String())
}
