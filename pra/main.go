package main

import (
	"fmt"
	"os"
)

func main() {
	var lastWord string

	inputString := os.Args[1]
	las := len(inputString) - 1

	for i := las; i > 0; i-- {
		if inputString[i] == ' ' && lastWord == "" {
			continue
		}
		if inputString[i] == ' ' && lastWord != "" {
			break
		}
		lastWord = string(inputString[i]) + lastWord
	}
	fmt.Println(lastWord)
}
