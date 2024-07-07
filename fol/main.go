package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Failed to open: %s\n", err)
		return
	}
	defer file.Close()
}
