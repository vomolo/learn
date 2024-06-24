package main

import (
	"github.com/01-edu/z01"
)

func FirstRune(s string) {
	for i, char := range s {
		if i == 0 {
			z01.PrintRune(char)
		}
	}
}

func main() {
	// z01.PrintRune(FirstRune("Hello!"))
	// z01.PrintRune(FirstRune("Salut!"))
	FirstRune("💗Ola!")
	z01.PrintRune('\n')
}
