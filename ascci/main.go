package main

import (
	"fmt"
	"os"
	"strings"

	"ascci/functions"
)

func main() {
	// Read command-line arguments for font file and input string
	if len(os.Args) < 3 {
		fmt.Println("Usage: run main.go <font_file><inptu_string>")
		return

	}

	fontFile := os.Args[1]

	inputString := os.Args[2]

	// if inputString == "\n" || inputString == "\\n" {
	// 	fmt.Println()
	// 	return
	// }

	// Load the font data
	fontArray, err := functions.Fontloader(fontFile)
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	//  Print the horizontally concatenated ASCII art representation of each word in the input string
	words := strings.Split(inputString, "\n")

	empty := emptyArray(words)
	if empty != "false"{
		fmt.Print(empty)
		return
	}
	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		// initialize an empty array to hold the lines of concatenated Ascii art for the word
		concatenatedLines := make([]string, 8)
		// Concatenate the Ascii art representation of each character in the word horizontally
		for _, char := range word {
			// retriev the Ascii art representation for the character using the AsciiChar function
			asciiArt := functions.AsciiChar(char, fontArray)
			// Check if the length of Ascii art matches the expected length
			if len(asciiArt) != len(concatenatedLines) {
				fmt.Println("Error: ASCII art representation does not match expected length")
				return
			}

			for i, line := range asciiArt {
				concatenatedLines[i] += line
			}

		}
		// Print the concatenated ASCII art lines for the word
		for _, line := range concatenatedLines {
			fmt.Println(line)
		}
		// // Print an empty line between words
		// fmt.Println()

	}
}

func emptyArray(words []string)string{
	result := ""

	for i, word := range words{
		if word != ""{
			result = "false"
			return result
		}
		if i != 0{
			result += "\n"
		}
	}
	return result
}