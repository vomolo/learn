package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arg := "hello"
	run := []rune(arg)

	z01.PrintRune(rune(run[len(run)-1]))

	// fmt.Println(string(arg[5]))
	z01.PrintRune('\n')
}
