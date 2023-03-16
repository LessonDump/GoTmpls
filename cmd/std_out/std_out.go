package main

import (
	"io"
	"os"
)

func main() {
	str := ""
	args := os.Args

	if len(args) == 1 {
		str = "Please give me one argument!"
	} else {
		str = args[1]
	}

	//goland:noinspection GoUnhandledErrorResult
	io.WriteString(os.Stdout, str)
	//goland:noinspection GoUnhandledErrorResult
	io.WriteString(os.Stdout, "\n")
}
