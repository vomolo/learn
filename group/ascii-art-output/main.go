package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-artoutput/functions"
)

func main() {
	usageMessage1 := `Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard`
	ar := os.Args[1]
	if strings.HasPrefix(ar, "-") && !strings.HasPrefix(ar, "--") {
		fmt.Println(usageMessage1)
		fmt.Println("EXPECING: --")
		os.Exit(1)
	}
	if strings.HasPrefix(ar, "--") && !strings.Contains(ar, "=") {
		fmt.Println(usageMessage1)
		fmt.Println("EXPECTING: =")
		os.Exit(1)
	}

	if strings.Contains(ar, "=") && !strings.Contains(ar, ".txt") {
		fmt.Println(usageMessage1)
		fmt.Println("EXPECING: .txt file extension")
		os.Exit(1)
	}
	// Define the output flag
	outputFlag := flag.String("output", "banner.txt", "Specify the output file name.")
	// Parse the flags
	flag.Parse()

	// Custom usage message to inform the user how to use the program
	usageMessage := "Usage: go run . --output=<fileName.txt> [STRING] [BANNER]"

	// Get the non-flag arguments
	args := flag.Args()

	// Validate the number of arguments. Expect 1 or 2 arguments: the input string and optionally the font file
	if len(args) == 0 || len(args) > 2 {
		fmt.Println("Check number of arguments")
		log.Fatal(usageMessage)

		return
	}

	// Assign the first argument to inputString
	inputString := args[0]

	var fontFile string
	// If there is a second argument, assign it to fontFile; otherwise, use "standard" as the default font
	if len(args) == 2 {
		fontFile = args[1]
	} else {
		fontFile = "standard"
	}

	// Load the font using the FontLoader function from the "functions" package
	fontArray, err := functions.FontLoader(fontFile)
	if err != nil {
		log.Fatal(err)
	}

	// Replace "\\n" in the input string with actual newline characters
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")
	// Split the input string into words using newline as the delimiter
	words := strings.Split(inputString, "\n")

	// Open the output file if specified
	var outputFile *os.File
	if *outputFlag != "" {
		outputFile, err = os.Create(*outputFlag)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		defer outputFile.Close() // Close the file when main() exits
	} else {
		// If no output file is specified, use os.Stdout (standard output)
		outputFile = os.Stdout
	}

	count := 0 // Initialize a count for empty lines
	for _, word := range words {
		// Handle empty lines by printing a newline in the output
		if word == "" {
			count++
			if count < len(words) {
				fmt.Fprintln(outputFile) // Print newline to the output file
				continue
			}
			continue
		}
		// Create a slice to hold the concatenated lines of ASCII art for this word
		concatenatedLines := make([]string, 8)

		// Process each character in the word
		for _, char := range word {
			asciiChar, err := functions.PrintChar(char, fontArray)
			if err != nil {
				log.Fatal("Error printing ascii", err)
			}
			// Concatenate the lines of the ASCII art for the character
			for i, line := range asciiChar {
				concatenatedLines[i] += line
			}
		}

		// Print the concatenated lines to the output file
		for _, lines := range concatenatedLines {
			fmt.Fprintln(outputFile, lines) // Print ASCII art line to the output file
		}
		fmt.Fprintln(outputFile)
	}

	// If an output file is specified, inform the user
	if *outputFlag != "" {
		fmt.Println("Output written to", *outputFlag)
	}
}
