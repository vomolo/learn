package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Error: Length!!!")
		return
	}
	input := os.Args[1]
	output := os.Args[2]
	suf := ".txt"

	if !strings.HasSuffix(input, suf) {
		fmt.Println("Error: Only .txt files allowed!!!")
		return
	}
	if !strings.HasSuffix(output, suf) {
		fmt.Println("Error: Only .txt files allowed!!!")
		return
	}

	inputFile, err := os.Open(input)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer inputFile.Close()

	stat, err := inputFile.Stat()
	if err != nil {
		fmt.Println("ERROR: Failed to get input file stats!!")
		return
	}
	if stat.Size() == 0 {
		fmt.Println("ERROR: Input file is empty")
		return
	}

	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer outputFile.Close()
}
