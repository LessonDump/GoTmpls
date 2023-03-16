package main

import (
	"io"
	"os"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	str := ""
	args := os.Args

	if len(args) == 1 {
		str = "Please give me one argument!"
	} else {
		str = args[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, str)
	io.WriteString(os.Stderr, "\n")
}
