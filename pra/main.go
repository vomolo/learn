package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	// las := os.Args[len(args)-1]
	// firs := os.Args[1]

	for _, run := range args[0] {
		z01.PrintRune(run)
	}
	z01.PrintRune('\n')
}
