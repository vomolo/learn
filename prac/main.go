package main

import (
	"os"
	"strings"
)

func WordMatch(s1, s2 string) string {
	if strings.Contains(s1, s2) {
		return "Match found"
	} else {
		return "No match found"
	}
}

func main() {
	if len(os.Args) != 3 {
		println("Usage: go run . <string1> <string2>")
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	result := WordMatch(s1, s2)
	println(result)
}
