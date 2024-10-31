package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println()
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	store := []rune{}

	seen := make(map[rune]bool)
	for _, char1 := range s1 {
		if !seen[char1] {
			store = append(store, char1)
			seen[char1] = true
		}
	}
	for _, char2 := range s2 {

		if !seen[char2] {
			store = append(store, char2)
			seen[char2] = true
		}
	}
	fmt.Print(string(store))
	fmt.Println()
}
