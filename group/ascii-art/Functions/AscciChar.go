package functions

import "fmt"

// function that generates ascii characters using provided rune
func AsciiChar(C rune, fontData [95][8]string) ([]string, error) {
	if C < 32 || C > 126 { // Check if the character is ASCII
		return nil, fmt.Errorf("character is not in ASCII range")
	}
	char := C - 32 // characters in Ascii begin from index 32 so we subtract to get the right indexing
	var lines []string

	for i := 0; i <= 7; i++ { // loop through each line in the character and add it to the lines array
		lines = append(lines, fontData[char][i])
	}

	return lines, nil
}
