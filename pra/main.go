package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str := " Hello World!"
	var currentWord []rune
	for _, run := range str {
		if run == ' ' {
			currentWord = currentWord[:0]
		} else {
			currentWord = append(currentWord, run)
		}
	}
	for _, char := range currentWord {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
