package main

import (
	"fmt"
	"os"
	"strings"

	"asci/functions"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("USAGE: go run . Fontfile.txt Teststring")
		return
	}
	fontFile := os.Args[1]
	testString := os.Args[2]
	fontArray, err := functions.FontLoader(fontFile)
	if err != nil {
		fmt.Println("Error loading fontfile", err)
	}

	testString = strings.ReplaceAll(testString, "\\n", "\n")
	concatenatedLines := make([]string, 8)

	words := strings.Split(testString, "\n")

	for _, word := range words {
		count := 0
		if word == "" {
			count++
			if count < len(words) {
				fmt.Println()
				continue
			}

			for _, char := range word {

				AscciAr, err := functions.AscciArt(char, fontArray)
				if err != nil {
					fmt.Println("Unable to ")
				}

				for i, line := range AscciAr {
					concatenatedLines[i] += line
				}
				for _, lines := range concatenatedLines {
					fmt.Println(lines)
				}

			}

		}

	}
}
