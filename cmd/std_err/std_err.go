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

//$ go run stdERR.go 2>/tmp/stdError
//This is Standard output
//$ cat /tmp/stdError
//Please give me one argument!

// ------

//$ go run stdERR.go 2>/dev/null
//This is Standard output

// ------

//$ go run stdERR.go >/tmp/output 2>&1
//$ cat /tmp/output
//This is Standard output
//Please give me one argument!

// ------

//$ go run stdERR.go >/dev/null 2>&1
