package tez

import (
	"strconv"
	"strings"
	"unicode"
)

func CapInt(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], "(cap") {
			if strings.Contains(slice[i], "(cap,") {
				num, _ := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				for j := i - num; j < i; j++ {
					slice[j] = capitalizeWord(slice[j])
				}
				slice = append(slice[:i], slice[i+2:]...)
			} else {
				slice[i-1] = capitalizeWord(slice[i-1])
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
	return strings.Join(slice, " ")
}

func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}

	// Capitalize the first character if it's a letter
	firstChar := rune(word[0])
	if unicode.IsLetter(firstChar) {
		return strings.ToUpper(string(firstChar)) + word[1:]
	}

	// Capitalize the first letter after a non-letter character
	var result strings.Builder
	capitalizeNext := true

	for _, char := range word {
		if unicode.IsLetter(char) {
			if capitalizeNext {
				result.WriteString(strings.ToUpper(string(char)))
				capitalizeNext = false
			} else {
				result.WriteString(string(char))
			}
		} else {
			result.WriteString(string(char))
			capitalizeNext = true
		}
	}

	return result.String()
}
