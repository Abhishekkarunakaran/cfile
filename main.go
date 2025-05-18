package main

import (
	"log"
	"os"
	clipboard "github.com/atotto/clipboard"
)

func main() {

	fileName := os.Args[1]
	if fileName == "" {
		log.Fatalln("File required.")
		os.Exit(0)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(0)
	}

	if err := clipboard.WriteAll(string(data)); err != nil {
		log.Fatalln(err.Error())
		os.Exit(0)
	}
}