package main

import (
	"flag"
	"fmt"
)

type Color string

const (
	ColorReset Color = "\u001b[0m"
	ColorBlue        = "\u001b[34m"
)

func colorize(color Color, message string) {
	fmt.Println(color, message, ColorReset)
}

func main() {
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	if *useColor {
		colorize(ColorBlue, "Hello")
		return
	}

	fmt.Println("Hello")
}
