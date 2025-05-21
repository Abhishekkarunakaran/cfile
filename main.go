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
		os.Exit(0)
	}
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("No such file exists.")
		fmt.Println("Enter a valid file name.")
		os.Exit(0)
	}

	if err := clipboard.WriteAll(string(data)); err != nil {
		fmt.Println("Failed to copy to clipboard")
		fmt.Println("Try again.")
		os.Exit(0)
	}
}
