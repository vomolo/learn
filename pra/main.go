package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	las := len(args) - 1

	for i := 0; i <= las; i++ {
		for _, run := range args[i] {
			z01.PrintRune(run)
		}
		z01.PrintRune('\n')
	}
}
