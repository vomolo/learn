package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	store := make(map[rune]bool)

	for _, char2 := range s2 {
		store[char2] = true
	}

	seen := make(map[rune]bool)
	sli := []rune{}

	for _, char1 := range s1 {
		if store[char1] && !seen[char1] {
			sli = append(sli, char1)
			seen[char1] = true
		}
	}

	fmt.Println(string(sli))

}
