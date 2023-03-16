package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	scanner := bufio.NewScanner(f)

	// --- defer os.Stdin.Close() ---
	//goland:noinspection GoUnhandledErrorResult
	//defer f.Close()

	// or, if you prefer:
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	// ------------------------------

	arr := make([]string, 0)
	for {
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}
	}

	fmt.Println("Your texts:", arr)

	// --- exit ---
	fmt.Println("Press Enter to exit")
	scanner.Scan()
}
