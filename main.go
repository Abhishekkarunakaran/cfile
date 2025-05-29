package main

import (
	"fmt"
	"os"

	clipboard "github.com/atotto/clipboard"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("File name required.")
		fmt.Println("Usage: cfile <filename>")
		return
	}
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("No such file exists.")
		fmt.Println("Enter a valid file name.")
		return
	}

	if err := clipboard.WriteAll(string(data)); err != nil {
		fmt.Println("Failed to copy to clipboard")
		fmt.Println("Try again.")
		return
	}
}
